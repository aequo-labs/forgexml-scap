// Package docs provides markdown and mermaid document rendering utilities
package docs

import (
	"bytes"
	"html/template"
	"regexp"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"go.abhg.dev/goldmark/mermaid"
)

// Renderer handles markdown and mermaid content rendering
type Renderer struct {
	md goldmark.Markdown
}

// NewRenderer creates a new markdown/mermaid renderer
func NewRenderer() *Renderer {
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			&mermaid.Extender{},
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
			html.WithUnsafe(),
		),
	)

	return &Renderer{md: md}
}

// RenderMarkdown converts markdown content to HTML
func (r *Renderer) RenderMarkdown(content []byte) (template.HTML, error) {
	var buf bytes.Buffer
	if err := r.md.Convert(content, &buf); err != nil {
		return "", err
	}
	return template.HTML(buf.String()), nil
}

// RenderMermaid wraps mermaid content in a markdown code block and renders it
func (r *Renderer) RenderMermaid(content []byte) (template.HTML, error) {
	// Wrap mermaid content in a code block for goldmark-mermaid to process
	wrapped := []byte("```mermaid\n" + string(content) + "\n```")
	return r.RenderMarkdown(wrapped)
}

// ExtractTitle extracts the title from markdown content (first h1 heading)
func ExtractTitle(content []byte, fallback string) string {
	contentStr := string(content)
	if strings.HasPrefix(contentStr, "# ") {
		lines := strings.Split(contentStr, "\n")
		if len(lines) > 0 {
			return strings.TrimSpace(lines[0][2:])
		}
	}
	return fallback
}

// TransformRelativeLinks transforms relative markdown links to /docs/ routes
// e.g., [link](./other-doc.md) becomes [link](/docs/other-doc)
func TransformRelativeLinks(content []byte, baseRoute string) []byte {
	contentStr := string(content)

	// Match markdown links with .md extension
	// Pattern: [text](./path.md) or [text](path.md) or [text](../path.md)
	linkPattern := regexp.MustCompile(`\]\((\./|\.\./)?([\w\-/]+)\.md\)`)

	contentStr = linkPattern.ReplaceAllStringFunc(contentStr, func(match string) string {
		// Extract the path without extension
		submatch := linkPattern.FindStringSubmatch(match)
		if len(submatch) >= 3 {
			path := submatch[2]
			// Get just the filename part if it has directories
			parts := strings.Split(path, "/")
			filename := parts[len(parts)-1]
			return "](" + baseRoute + "/" + filename + ")"
		}
		return match
	})

	// Also handle .mermaid links
	mermaidPattern := regexp.MustCompile(`\]\((\./|\.\./)?([\w\-/]+)\.mermaid\)`)
	contentStr = mermaidPattern.ReplaceAllStringFunc(contentStr, func(match string) string {
		submatch := mermaidPattern.FindStringSubmatch(match)
		if len(submatch) >= 3 {
			path := submatch[2]
			parts := strings.Split(path, "/")
			filename := parts[len(parts)-1]
			return "](" + baseRoute + "/" + filename + ")"
		}
		return match
	})

	return []byte(contentStr)
}

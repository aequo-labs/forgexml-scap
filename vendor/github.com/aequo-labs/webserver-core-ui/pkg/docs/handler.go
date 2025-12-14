package docs

import (
	"html/template"
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// DocInfo represents information about a single documentation file
type DocInfo struct {
	Title    string
	Path     string
	Category string
}

// DocHandlerConfig holds configuration for the documentation handlers
type DocHandlerConfig struct {
	// DocsFS is the embedded filesystem containing documentation files
	DocsFS fs.FS
	// DocsDir is the directory within the filesystem containing docs (e.g., "doc")
	DocsDir string
	// BaseRoute is the URL route prefix for docs (e.g., "/docs")
	BaseRoute string
	// IndexTemplate is the template name for the doc index page
	IndexTemplate string
	// ViewTemplate is the template name for viewing a single doc
	ViewTemplate string
	// CategoryRules maps filename suffixes to category names
	CategoryRules map[string]string
	// DefaultCategory is the category for files not matching any rule
	DefaultCategory string
}

// DefaultDocHandlerConfig returns a config with sensible defaults
func DefaultDocHandlerConfig() DocHandlerConfig {
	return DocHandlerConfig{
		DocsDir:         "doc",
		BaseRoute:       "/docs",
		IndexTemplate:   "doc-index-content",
		ViewTemplate:    "doc-view-content",
		DefaultCategory: "General",
		CategoryRules: map[string]string{
			"-guide": "Integration Guides",
		},
	}
}

// DocHandlers handles documentation-related HTTP requests
type DocHandlers struct {
	config   DocHandlerConfig
	renderer *Renderer
	render   func(w http.ResponseWriter, templateName string, data interface{}) error
}

// NewDocHandlers creates a new DocHandlers instance
func NewDocHandlers(config DocHandlerConfig, renderFunc func(w http.ResponseWriter, templateName string, data interface{}) error) *DocHandlers {
	return &DocHandlers{
		config:   config,
		renderer: NewRenderer(),
		render:   renderFunc,
	}
}

// ListDocs returns a list of all documentation files grouped by category
func (h *DocHandlers) ListDocs() (map[string][]DocInfo, error) {
	var docFiles []string

	err := fs.WalkDir(h.config.DocsFS, h.config.DocsDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && (strings.HasSuffix(path, ".md") || strings.HasSuffix(path, ".mermaid")) {
			docFiles = append(docFiles, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	titleCaser := cases.Title(language.English)
	docsByCategory := make(map[string][]DocInfo)

	for _, path := range docFiles {
		filename := filepath.Base(path)
		ext := filepath.Ext(filename)
		name := filename[:len(filename)-len(ext)]

		// Convert kebab-case to title case
		title := strings.ReplaceAll(name, "-", " ")
		title = titleCaser.String(title)

		// Determine category based on filename
		category := h.config.DefaultCategory
		for suffix, cat := range h.config.CategoryRules {
			if strings.HasSuffix(name, suffix) {
				category = cat
				break
			}
		}

		doc := DocInfo{
			Title:    title,
			Path:     h.config.BaseRoute + "/" + name,
			Category: category,
		}
		docsByCategory[category] = append(docsByCategory[category], doc)
	}

	return docsByCategory, nil
}

// HandleDocIndex renders the documentation index page
func (h *DocHandlers) HandleDocIndex(w http.ResponseWriter, r *http.Request) {
	docsByCategory, err := h.ListDocs()
	if err != nil {
		http.Error(w, "Error listing documentation", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"DocsByCategory": docsByCategory,
		"Title":          "Documentation",
		"NeedsMermaid":   true,
	}

	if err := h.render(w, h.config.IndexTemplate, data); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

// HandleDocView renders a specific documentation file
func (h *DocHandlers) HandleDocView(w http.ResponseWriter, r *http.Request) {
	// Extract the document name from the URL
	docName := strings.TrimPrefix(r.URL.Path, h.config.BaseRoute+"/")
	if docName == "" {
		h.HandleDocIndex(w, r)
		return
	}

	// Try to find the corresponding markdown file
	mdPath := h.config.DocsDir + "/" + docName + ".md"
	mdContent, err := fs.ReadFile(h.config.DocsFS, mdPath)
	if err != nil {
		// If markdown file not found, try mermaid file
		mermaidPath := h.config.DocsDir + "/" + docName + ".mermaid"
		mermaidContent, mermaidErr := fs.ReadFile(h.config.DocsFS, mermaidPath)
		if mermaidErr != nil {
			http.NotFound(w, r)
			return
		}

		// Render mermaid content
		htmlContent, renderErr := h.renderer.RenderMermaid(mermaidContent)
		if renderErr != nil {
			http.Error(w, "Error rendering mermaid", http.StatusInternalServerError)
			return
		}

		h.renderDocView(w, docName, htmlContent)
		return
	}

	// Transform relative links to /docs/ routes
	mdContent = TransformRelativeLinks(mdContent, h.config.BaseRoute)

	// Extract title from content
	title := ExtractTitle(mdContent, docName)

	// Render markdown content
	htmlContent, err := h.renderer.RenderMarkdown(mdContent)
	if err != nil {
		http.Error(w, "Error rendering markdown", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Title":        title,
		"Content":      htmlContent,
		"DocName":      title,
		"NeedsMermaid": true,
	}

	if err := h.render(w, h.config.ViewTemplate, data); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func (h *DocHandlers) renderDocView(w http.ResponseWriter, docName string, content template.HTML) {
	data := map[string]interface{}{
		"Title":        docName,
		"Content":      content,
		"DocName":      docName,
		"NeedsMermaid": true,
	}

	if err := h.render(w, h.config.ViewTemplate, data); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

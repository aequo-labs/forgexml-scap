package webserver

import (
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"strings"
	"time"

	"github.com/aequo-labs/webserver-core-ui/pkg/assets"
	"github.com/aequo-labs/webserver-core-ui/pkg/docs"
	"github.com/aequo-labs/webserver-core/pkg/logging"
	"github.com/aequo-labs/webserver-core/pkg/server"
	templatefuncs "github.com/aequo-labs/webserver-core/pkg/template"
)

// PageDataEnricher is a callback function that can add application-specific data to PageData
// This allows applications to add NavItems, DropdownMenus, and other data to shared pages like /about
type PageDataEnricher func(data *PageData)

// UIServer extends the base server with UI-specific functionality
type UIServer struct {
	*server.Server   // Embed base server
	templates        *template.Template
	funcMap          template.FuncMap
	logger           *logging.Logger
	aboutInfo        *AboutInfo       // Application-specific about information
	docsFS           fs.FS            // Embedded filesystem for documentation
	docsDir          string           // Directory within docsFS containing docs
	settingsURL      string           // URL for settings page (default: /config)
	pageDataEnricher PageDataEnricher // Optional callback to enrich page data
}

// PageData represents the standard data structure for templates
type PageData struct {
	Title               string
	Description         string // Page meta description
	ThemeMode           string
	ThemePreference     string
	AppName             string
	Version             string // Application version for footer
	NavItems            []NavItem
	HasDropdownMenus    bool
	DropdownMenus       []DropdownMenu
	FlashMessages       []FlashMessage
	ShowNavigation      bool
	ShowFooter          bool
	ShowSettings        bool
	ShowThemeToggle     bool
	ShowStatusIndicator bool
	ShowAbout           bool       // Show about icon in navbar
	ShowDocs            bool       // Show docs link in navbar/about page
	AboutInfo           *AboutInfo // Application-specific about information
	SettingsURL         string     // URL for settings page (default: /config)
	ShowLogin           bool
	IsAuthenticated     bool
	Username            string
	DomainCSS           []string
	Copyright           string
	FooterLinks         []NavItem
	NeedsCharts         bool
	NeedsExport         bool
	NeedsMermaid        bool
	NeedsPanZoom        bool
}

// NavItem represents a navigation menu item
type NavItem struct {
	Title    string
	URL      string
	IsActive bool
	Icon     string
}

// DropdownMenu represents a dropdown navigation menu
type DropdownMenu struct {
	Title string
	Items []NavItem
}

// FlashMessage represents a flash message
type FlashMessage struct {
	Type    string
	Message string
}

// AboutInfo contains application-specific information for the about page
type AboutInfo struct {
	AppName     string            // Application name
	Version     string            // Application version
	Description string            // Brief description of the application
	Features    []string          // List of key features
	License     interface{}       // License information (application-specific)
	Links       []NavItem         // Useful links (documentation, support, etc.)
	ExtraInfo   map[string]string // Additional key-value info to display
}

// NewUIServer creates a new UI server with base templates and assets already configured
func NewUIServer(logger *logging.Logger) (*UIServer, error) {
	if logger == nil {
		var err error
		cfg := logging.DefaultConfig()
		cfg.LogFileName = "uiserver.log"
		logger, err = logging.NewLogger(cfg)
		if err != nil {
			return nil, fmt.Errorf("failed to create logger: %w", err)
		}
	}
	s := &UIServer{
		Server:  server.New(),
		funcMap: templatefuncs.GetStandardFuncMap(),
		logger:  logger,
	}

	// Setup optional middleware (gzip, minification) - must be called before adding routes
	s.Server.SetupOptionalMiddleware()

	// Load base templates
	if err := s.loadBaseTemplates(); err != nil {
		return nil, err
	}

	// Setup base static assets
	s.setupBaseAssets()

	// Setup base API endpoints
	s.setupBaseAPI()

	return s, nil
}

// loadBaseTemplates loads all the shared templates from webserver-core-ui
func (s *UIServer) loadBaseTemplates() error {
	s.templates = template.New("").Funcs(s.funcMap)

	// Load shared templates from webserver-core-ui
	sharedTemplateFiles := []string{
		"templates/base/layout.gohtml",
		"templates/includes/head.gohtml",
		"templates/includes/css-includes.gohtml",
		"templates/includes/navigation.gohtml",
		"templates/includes/theme-toggle.gohtml",
		"templates/includes/footer.gohtml",
		"templates/includes/js-includes.gohtml",
		"templates/includes/flash-messages.gohtml",
		"templates/includes/about-content.gohtml",
		"templates/docs/doc-index-content.gohtml",
		"templates/docs/doc-view-content.gohtml",
		// Note: page-scripts.gohtml removed to avoid conflicts with page-specific scripts
	}

	for _, filename := range sharedTemplateFiles {
		data, err := assets.TemplateFiles.ReadFile(filename)
		if err != nil {
			s.logger.Warn("Could not load shared template", "filename", filename, "error", err)
			continue
		}
		// Extract template name based on file structure
		var templateName string
		if strings.HasPrefix(filename, "templates/base/") {
			templateName = filename[len("templates/base/"):]
		} else if strings.HasPrefix(filename, "templates/includes/") {
			templateName = filename[len("templates/includes/"):]
		} else if strings.HasPrefix(filename, "templates/docs/") {
			templateName = filename[len("templates/docs/"):]
		} else {
			templateName = filename[len("templates/"):]
		}
		templateName = templateName[:len(templateName)-7] // remove .gohtml
		_, err = s.templates.New(templateName).Parse(string(data))
		if err != nil {
			return err
		}
	}

	return nil
}

// setupBaseAssets configures the base static asset serving
func (s *UIServer) setupBaseAssets() {
	// Static assets from webserver-core-ui - create sub-filesystem starting at "static"
	staticFS, err := fs.Sub(assets.StaticFiles, "static")
	if err != nil {
		s.logger.Error("Error creating static sub-filesystem", "error", err)
		panic(err)
	}
	s.Router().Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))))
}

// setupBaseAPI configures the base API endpoints
func (s *UIServer) setupBaseAPI() {
	// Status endpoint for navbar status indicator
	s.Router().HandleFunc("GET /api/status", s.statusHandler)

	// Theme preference endpoint (basic implementation)
	s.Router().HandleFunc("GET /api/user/theme-preference", s.themePreferenceHandler)
	s.Router().HandleFunc("POST /api/user/theme-preference", s.themePreferenceHandler)

	// Theme switching endpoint for programmatic theme changes
	s.Router().HandleFunc("POST /api/theme/toggle", s.themeToggleHandler)
	s.Router().HandleFunc("POST /api/theme/set", s.themeSetHandler)
}

// statusHandler provides a default status endpoint
func (s *UIServer) statusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	status := map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now().Format(time.RFC3339),
	}
	json.NewEncoder(w).Encode(status)
}

// themePreferenceHandler provides a basic theme preference endpoint
func (s *UIServer) themePreferenceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		// Default to light theme
		response := map[string]interface{}{
			"theme": "light",
		}
		json.NewEncoder(w).Encode(response)
	} else if r.Method == "POST" {
		// Accept theme changes but don't persist them (basic implementation)
		response := map[string]interface{}{
			"success": true,
		}
		json.NewEncoder(w).Encode(response)
	}
}

// themeToggleHandler toggles between light and dark themes
func (s *UIServer) themeToggleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// For this basic implementation, we'll just return success
	// In a real implementation, you might check current theme and toggle it
	response := map[string]interface{}{
		"success":  true,
		"message":  "Theme toggled successfully",
		"newTheme": "dark", // This would be dynamic in a real implementation
	}
	json.NewEncoder(w).Encode(response)
}

// themeSetHandler sets a specific theme
func (s *UIServer) themeSetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request struct {
		Theme string `json:"theme"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate theme
	if request.Theme != "light" && request.Theme != "dark" {
		http.Error(w, "Invalid theme. Must be 'light' or 'dark'", http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"success": true,
		"message": "Theme set successfully",
		"theme":   request.Theme,
	}
	json.NewEncoder(w).Encode(response)
}

// AddTemplatesFromFS adds templates from an embedded filesystem
func (s *UIServer) AddTemplatesFromFS(templateFS embed.FS, templateDir string) error {
	s.logger.Debug("Loading templates from directory", "dir", templateDir)
	files, err := templateFS.ReadDir(templateDir)
	if err != nil {
		s.logger.Error("Error reading template directory", "dir", templateDir, "error", err)
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if !strings.HasSuffix(file.Name(), ".gohtml") {
			continue
		}

		s.logger.Debug("Loading template", "name", file.Name())
		data, err := templateFS.ReadFile(templateDir + "/" + file.Name())
		if err != nil {
			s.logger.Error("Error reading template file", "name", file.Name(), "error", err)
			return err
		}
		_, err = s.templates.Parse(string(data))
		if err != nil {
			s.logger.Error("Error parsing template", "name", file.Name(), "error", err)
			return err
		}
		s.logger.Debug("Successfully loaded template", "name", file.Name())

		// Debug: Log what template definitions were found
		if strings.Contains(string(data), "{{define") {
			lines := strings.Split(string(data), "\n")
			for i, line := range lines {
				if strings.Contains(line, "{{define") {
					s.logger.Debug("Found template definition", "line", i+1, "content", strings.TrimSpace(line))
				}
			}
		}
	}

	return nil
}

// AddAssetsFromFS adds static assets from an embedded filesystem
func (s *UIServer) AddAssetsFromFS(assetFS embed.FS, assetDir string, urlPrefix string) {
	subFS, err := fs.Sub(assetFS, assetDir)
	if err != nil {
		s.logger.Error("Error creating asset sub-filesystem", "dir", assetDir, "error", err)
		return
	}
	s.Router().Handle("GET "+urlPrefix, http.StripPrefix(urlPrefix, http.FileServer(http.FS(subFS))))
}

// AddCustomFunction adds a custom template function
func (s *UIServer) AddCustomFunction(name string, fn interface{}) {
	s.funcMap[name] = fn
	// Recreate templates with new function map
	s.loadBaseTemplates()
}

// Templates returns the template collection for rendering
func (s *UIServer) Templates() *template.Template {
	return s.templates
}

// RenderTemplate renders a template with the given data
func (s *UIServer) RenderTemplate(w http.ResponseWriter, templateName string, data interface{}) error {
	if err := s.templates.ExecuteTemplate(w, templateName, data); err != nil {
		s.logger.Error("Template error", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	return nil
}

// RenderPageWithContent renders the base layout with a specific content template
// This allows dynamic content injection into the master layout at runtime
//
// Usage example:
//
//	srv.RenderPageWithContent(w, "dashboard-content", data)
//	srv.RenderPageWithContent(w, "tools-content", data)
//
// The base layout expects {{template "content" .}} but we can have any named content template
func (s *UIServer) RenderPageWithContent(w http.ResponseWriter, contentTemplateName string, data interface{}) error {
	s.logger.Debug("RenderPageWithContent called", "template", contentTemplateName)
	// Clone the templates to avoid modifying the original
	tmpl, err := s.templates.Clone()
	if err != nil {
		s.logger.Error("Failed to clone templates", "error", err)
		http.Error(w, "Template error", http.StatusInternalServerError)
		return err
	}

	// Create a dynamic "content" template that calls our specific content template
	contentWrapper := fmt.Sprintf(`{{define "content"}}{{template "%s" .}}{{end}}`, contentTemplateName)

	// Parse the content wrapper template
	_, err = tmpl.Parse(contentWrapper)
	if err != nil {
		s.logger.Error("Failed to create content wrapper", "template", contentTemplateName, "error", err)
		http.Error(w, "Template error", http.StatusInternalServerError)
		return err
	}

	// Create a page-specific page-scripts wrapper
	// Create a page-specific page-scripts wrapper to avoid template name collisions
	// Different pages may have different JavaScript requirements
	pageScriptsWrapper := ""

	// Check if there's a page-specific scripts template defined
	// Use the naming convention: "{contentTemplate}-scripts"
	scriptsTemplateName := strings.Replace(contentTemplateName, "-content", "-page-scripts", 1)

	// Check if the scripts template exists
	if tmpl.Lookup(scriptsTemplateName) != nil {
		// Use the existing page-specific scripts template
		pageScriptsWrapper = fmt.Sprintf(`{{define "page-scripts"}}{{template "%s" .}}{{end}}`, scriptsTemplateName)
		s.logger.Debug("Using existing scripts template", "name", scriptsTemplateName)
	} else {
		// No page-specific scripts template, create empty page-scripts
		pageScriptsWrapper = `{{define "page-scripts"}}{{end}}`
		s.logger.Debug("No scripts template found, using empty page-scripts", "template", contentTemplateName)
	}

	// Parse the page-scripts wrapper - this will override the generic page-scripts template
	_, err = tmpl.Parse(pageScriptsWrapper)
	if err != nil {
		s.logger.Error("Failed to create page-scripts wrapper", "template", contentTemplateName, "error", err)
		// Continue anyway - page-scripts are optional
	} else {
		s.logger.Debug("Successfully parsed page-scripts wrapper", "template", contentTemplateName)
	}

	// Execute the layout template
	if err := tmpl.ExecuteTemplate(w, "layout", data); err != nil {
		s.logger.Error("Template execution error", "template", contentTemplateName, "error", err)
		http.Error(w, "Template error", http.StatusInternalServerError)
		return err
	}

	return nil
}

// GetBasePageData returns a PageData struct with sensible defaults
// ShowAbout and ShowDocs are automatically set based on whether SetAboutInfo/RegisterDocsFS were called
func (s *UIServer) GetBasePageData(activeNavItem string) PageData {
	return PageData{
		ThemeMode:           "light",
		ThemePreference:     "light",
		AppName:             "Web Application",
		Version:             "",
		Description:         "",
		NavItems:            []NavItem{},
		HasDropdownMenus:    false,
		DropdownMenus:       []DropdownMenu{},
		FlashMessages:       []FlashMessage{},
		ShowNavigation:      true,
		ShowFooter:          true,
		ShowSettings:        true,
		ShowThemeToggle:     true,
		ShowStatusIndicator: true,
		ShowAbout:           s.aboutInfo != nil,
		ShowDocs:            s.docsFS != nil,
		AboutInfo:           s.aboutInfo,
		SettingsURL:         s.GetSettingsURL(),
		ShowLogin:           false,
		IsAuthenticated:     false,
		Username:            "",
		DomainCSS:           []string{},
		Copyright:           "",
		FooterLinks:         []NavItem{},
		NeedsCharts:         false,
		NeedsExport:         false,
		NeedsMermaid:        false,
		NeedsPanZoom:        false,
	}
}

// SetAboutInfo configures the application-specific about information
// This enables the about icon in the navbar and the /about page
func (s *UIServer) SetAboutInfo(info AboutInfo) {
	s.aboutInfo = &info
	// Register the /about route
	s.Router().HandleFunc("GET /about", s.handleAbout)
	s.logger.Debug("About page configured", "appName", info.AppName)
}

// RegisterDocsFS registers an embedded filesystem containing documentation files
// docFS is the embedded filesystem, docDir is the directory within it containing docs
func (s *UIServer) RegisterDocsFS(docFS embed.FS, docDir string) {
	subFS, err := fs.Sub(docFS, docDir)
	if err != nil {
		s.logger.Error("Failed to create docs sub-filesystem", "dir", docDir, "error", err)
		return
	}
	s.docsFS = subFS
	s.docsDir = docDir

	// Register doc routes
	s.Router().HandleFunc("GET /docs", s.handleDocIndex)
	s.Router().HandleFunc("GET /docs/{path...}", s.handleDocView)
	s.logger.Debug("Documentation filesystem registered", "dir", docDir)
}

// HasAboutInfo returns true if about information has been configured
func (s *UIServer) HasAboutInfo() bool {
	return s.aboutInfo != nil
}

// HasDocsFS returns true if a documentation filesystem has been registered
func (s *UIServer) HasDocsFS() bool {
	return s.docsFS != nil
}

// SetSettingsURL configures the URL for the settings page
// Default is "/config", but applications can set it to "/settings" or any other path
func (s *UIServer) SetSettingsURL(url string) {
	s.settingsURL = url
}

// GetSettingsURL returns the configured settings URL (default: /config)
func (s *UIServer) GetSettingsURL() string {
	if s.settingsURL == "" {
		return "/config"
	}
	return s.settingsURL
}

// SetPageDataEnricher configures a callback that will be called to enrich PageData
// for shared pages like /about and /docs. This allows applications to add their
// NavItems, DropdownMenus, and other application-specific data.
func (s *UIServer) SetPageDataEnricher(enricher PageDataEnricher) {
	s.pageDataEnricher = enricher
}

// enrichPageData applies the page data enricher if one is configured
func (s *UIServer) enrichPageData(data *PageData) {
	if s.pageDataEnricher != nil {
		s.pageDataEnricher(data)
	}
}

// handleAbout renders the about page
func (s *UIServer) handleAbout(w http.ResponseWriter, r *http.Request) {
	data := s.GetBasePageData("about")
	data.Title = "About"
	data.ShowAbout = true
	data.ShowDocs = s.HasDocsFS()
	data.AboutInfo = s.aboutInfo

	// Allow application to enrich page data with NavItems, etc.
	s.enrichPageData(&data)

	// Set description from AboutInfo if available
	if s.aboutInfo != nil && s.aboutInfo.Description != "" {
		data.Description = s.aboutInfo.Description
	} else {
		data.Description = "About this application"
	}

	if err := s.RenderPageWithContent(w, "about-content", data); err != nil {
		s.logger.Error("Failed to render about page", "error", err)
	}
}

// GetAboutInfo returns the configured about information
func (s *UIServer) GetAboutInfo() *AboutInfo {
	return s.aboutInfo
}

// GetDocsFS returns the registered documentation filesystem
func (s *UIServer) GetDocsFS() fs.FS {
	return s.docsFS
}

// DocPageData extends PageData with documentation-specific fields
type DocPageData struct {
	PageData
	DocCategories map[string][]DocInfo // For doc index
	DocTitle      string               // For doc view
	DocContent    template.HTML        // For doc view (rendered HTML)
	DocPath       string               // Current doc path
}

// DocInfo contains information about a documentation file
type DocInfo struct {
	Title    string
	Path     string
	Category string
}

// handleDocIndex renders the documentation index page
func (s *UIServer) handleDocIndex(w http.ResponseWriter, r *http.Request) {
	if s.docsFS == nil {
		http.Error(w, "Documentation not available", http.StatusNotFound)
		return
	}

	data := DocPageData{
		PageData: s.GetBasePageData("docs"),
	}
	data.Title = "Documentation"
	data.ShowAbout = s.HasAboutInfo()
	data.ShowDocs = true
	data.AboutInfo = s.aboutInfo
	data.NeedsMermaid = true

	// Allow application to enrich page data with NavItems, etc.
	s.enrichPageData(&data.PageData)

	// List all docs from the filesystem
	categories := make(map[string][]DocInfo)
	err := fs.WalkDir(s.docsFS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		// Only include .md and .mermaid files
		if !strings.HasSuffix(path, ".md") && !strings.HasSuffix(path, ".mermaid") {
			return nil
		}

		// Determine category from directory structure
		category := "General"
		parts := strings.Split(path, "/")
		if len(parts) > 1 {
			category = strings.Title(parts[0])
		}

		// Extract title from filename
		title := strings.TrimSuffix(strings.TrimSuffix(d.Name(), ".md"), ".mermaid")
		title = strings.ReplaceAll(title, "-", " ")
		title = strings.ReplaceAll(title, "_", " ")
		title = strings.Title(title)

		categories[category] = append(categories[category], DocInfo{
			Title:    title,
			Path:     path,
			Category: category,
		})
		return nil
	})

	if err != nil {
		s.logger.Error("Failed to list documentation files", "error", err)
	}

	data.DocCategories = categories

	if err := s.RenderPageWithContent(w, "doc-index-content", data); err != nil {
		s.logger.Error("Failed to render doc index", "error", err)
	}
}

// handleDocView renders a specific documentation file
func (s *UIServer) handleDocView(w http.ResponseWriter, r *http.Request) {
	if s.docsFS == nil {
		http.Error(w, "Documentation not available", http.StatusNotFound)
		return
	}

	docPath := r.PathValue("path")
	if docPath == "" {
		http.Redirect(w, r, "/docs", http.StatusFound)
		return
	}

	// Read the file
	content, err := fs.ReadFile(s.docsFS, docPath)
	if err != nil {
		s.logger.Error("Failed to read doc file", "path", docPath, "error", err)
		http.Error(w, "Document not found", http.StatusNotFound)
		return
	}

	data := DocPageData{
		PageData: s.GetBasePageData("docs"),
	}
	data.ShowAbout = s.HasAboutInfo()
	data.ShowDocs = true
	data.AboutInfo = s.aboutInfo
	data.DocPath = docPath
	data.NeedsMermaid = true

	// Allow application to enrich page data with NavItems, etc.
	s.enrichPageData(&data.PageData)

	// Extract title from filename
	filename := docPath
	if idx := strings.LastIndex(docPath, "/"); idx >= 0 {
		filename = docPath[idx+1:]
	}
	title := strings.TrimSuffix(strings.TrimSuffix(filename, ".md"), ".mermaid")
	title = strings.ReplaceAll(title, "-", " ")
	title = strings.ReplaceAll(title, "_", " ")
	title = strings.Title(title)
	data.DocTitle = title
	data.Title = title + " - Documentation"

	// Render content based on file type
	if strings.HasSuffix(docPath, ".mermaid") {
		// Wrap mermaid content in a div
		data.DocContent = template.HTML(fmt.Sprintf(`<div class="mermaid">%s</div>`, string(content)))
	} else {
		// Use the docs renderer for markdown
		renderer := s.getDocsRenderer()
		if renderer != nil {
			rendered, err := renderer.RenderMarkdown(content)
			if err != nil {
				s.logger.Error("Failed to render markdown", "path", docPath, "error", err)
				data.DocContent = template.HTML("<p>Error rendering document</p>")
			} else {
				data.DocContent = rendered
			}
		} else {
			// Fallback: just show as preformatted text
			data.DocContent = template.HTML(fmt.Sprintf("<pre>%s</pre>", template.HTMLEscapeString(string(content))))
		}
	}

	if err := s.RenderPageWithContent(w, "doc-view-content", data); err != nil {
		s.logger.Error("Failed to render doc view", "error", err)
	}
}

// getDocsRenderer returns the docs renderer, creating it if needed
func (s *UIServer) getDocsRenderer() *docs.Renderer {
	return docs.NewRenderer()
}

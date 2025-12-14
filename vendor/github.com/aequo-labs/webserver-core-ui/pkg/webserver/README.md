# webserver-core-ui/pkg/webserver

This package provides a simplified way to create web servers that automatically include all the base templates and static assets from webserver-core-ui, while allowing projects to easily add their own templates and assets.

## Problem Solved

Previously, every project had to:
- Manually configure embedded filesystem serving with correct paths
- Load shared templates from webserver-core-ui with proper error handling
- Handle the complexity of template function registration
- Set up all the required PageData fields that shared templates expect
- Debug filesystem path issues when static assets don't serve

This package eliminates all that complexity.

## Basic Usage

```go
package main

import (
    "log"
    "net/http"
    
    "github.com/Code-Monger/webserver-core-ui/pkg/webserver"
)

func main() {
    // Create server with base templates and assets pre-configured
    srv, err := webserver.New()
    if err != nil {
        log.Fatal(err)
    }

    // Add your routes
    r := srv.Router()
    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data := webserver.PageData{
            Title:           "My App",
            AppName:         "My Application", 
            ShowNavigation:  true,
            ShowFooter:      true,
            ShowThemeToggle: true,
        }
        
        // Use RenderPageWithContent for dynamic template injection
        srv.RenderPageWithContent(w, "dashboard-content", data)
    })

    // Start server
    log.Println("Server starting on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
```

## Features

### Automatic Base Setup
- ✅ All shared templates from webserver-core-ui loaded automatically
- ✅ Static assets served correctly with proper filesystem paths (`/static/`)
- ✅ Template functions registered (formatNumber, formatCurrency, formatBytes, etc.)
- ✅ All required PageData fields with sensible defaults
- ✅ Built-in API endpoints for status and theme preferences

### Easy Extension
- ✅ Add your own templates: `srv.AddTemplatesFromFS(templates, "templates")`
- ✅ Add your own assets: `srv.AddAssetsFromFS(assets, "assets", "/assets/")`
- ✅ Access underlying router: `srv.Router()`

### Template Structure
The server expects your templates to follow the layout/content pattern:

**layout.html** (provided by webserver-core-ui):
```html
<!DOCTYPE html>
<html>
<head>{{template "head" .}}</head>
<body>
    {{template "navigation" .}}
    <main>{{template "content" .}}</main>
    {{template "footer" .}}
    {{template "js-includes" .}}
</body>
</html>
```

**Your content templates**:
```html
{{define "dashboard-content"}}
<div class="title-row">
    <div class="page-title">
        <h1 class="title is-4">{{.Title}}</h1>
        <p class="subtitle is-6 has-text-grey">{{.AppName}}</p>
    </div>
</div>

<div class="content-section">
    <div class="card">
        <div class="card-content">
            <p>Welcome to {{.AppName}}!</p>
            <p>Users: {{formatNumber .Stats.Users}}</p>
            <p>Revenue: {{formatCurrency .Stats.Revenue}}</p>
        </div>
    </div>
</div>
{{end}}
```

### PageData Structure

The `PageData` struct includes all fields that shared templates expect:

```go
type PageData struct {
    Title               string
    ThemeMode           string
    ThemePreference     string
    AppName             string
    NavItems            []NavItem
    HasDropdownMenus    bool
    DropdownMenus       []DropdownMenu
    FlashMessages       []FlashMessage
    ShowNavigation      bool
    ShowFooter          bool
    ShowSettings        bool
    ShowThemeToggle     bool
    ShowStatusIndicator bool
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
```

### Navigation Types

```go
type NavItem struct {
    Title    string
    URL      string
    IsActive bool
    Icon     string
}

type DropdownMenu struct {
    Title string
    Items []NavItem
}

type FlashMessage struct {
    Type    string
    Message string
}
```

## Advanced Usage

### Complete Example with Navigation

```go
func dashboardHandler(srv *webserver.Server) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        data := webserver.PageData{
            Title:           "Dashboard",
            AppName:         "My App",
            ShowNavigation:  true,
            ShowFooter:      true,
            ShowThemeToggle: true,
            NavItems: []webserver.NavItem{
                {Title: "Home", URL: "/", Icon: "fas fa-home", IsActive: true},
                {Title: "Tools", URL: "/tools", Icon: "fas fa-tools"},
                {Title: "Resources", URL: "/resources", Icon: "fas fa-database"},
            },
        }
        
        srv.RenderPageWithContent(w, "dashboard-content", data)
    }
}
```

### Extending PageData

For app-specific data, embed the base PageData:

```go
type AppPageData struct {
    webserver.PageData
    Stats     ServerStats
    Tools     []Tool
    Resources []Resource
}

func getAppData(activeNav string) AppPageData {
    return AppPageData{
        PageData: webserver.PageData{
            Title:           "Dashboard",
            AppName:         "My Application",
            ShowNavigation:  true,
            ShowFooter:      true,
            ShowThemeToggle: true,
            NavItems: []webserver.NavItem{
                {Title: "Home", URL: "/", Icon: "fas fa-home", IsActive: activeNav == "home"},
                {Title: "Tools", URL: "/tools", Icon: "fas fa-tools", IsActive: activeNav == "tools"},
            },
        },
        Stats:     getServerStats(),
        Tools:     getAvailableTools(),
        Resources: getResources(),
    }
}
```

### Template Functions

Built-in template functions available in all templates:

```html
<!-- Number and currency formatting -->
<p>Total: {{formatNumber 12345}}</p>          <!-- 12,345 -->
<p>Price: {{formatCurrency 99.95}}</p>        <!-- $99.95 -->
<p>Size: {{formatBytes 1048576}}</p>          <!-- 1.0 MB -->
<p>Rate: {{formatPercent 0.856}}</p>          <!-- 85.6% -->

<!-- Date formatting -->
<p>Created: {{formatDate .CreatedAt}}</p>     <!-- 2025-01-15 -->
<p>Updated: {{formatDateTime .UpdatedAt}}</p> <!-- 2025-01-15 14:30:25 -->

<!-- Math operations -->
<p>Total: {{add .Price .Tax}}</p>
<p>Amount: {{multiply .Quantity .Price}}</p>

<!-- JSON encoding -->
<script>
const data = {{json .Data}};
</script>
```

### Built-in API Endpoints

The server automatically provides these endpoints:

- **`GET /api/status`** - Health check endpoint for navbar status indicator
- **`GET/POST /api/user/theme-preference`** - Theme preference management

## What You Get Out of the Box

- **Bulma CSS Framework** - Modern, responsive CSS framework
- **FontAwesome Icons** - Complete icon library
- **Theme System** - Light/dark mode with automatic switching
- **Navigation Template** - Responsive navbar with dropdowns
- **Loading Overlay** - Automatic loading indicators
- **Sticky Tables** - Hardware-accelerated sticky headers
- **Modal Components** - Theme-aware modal dialogs
- **Template Functions** - Number, currency, date formatting and more

## Migration from Manual Setup

**Old way** (lots of boilerplate):
```go
// 50+ lines of template loading, filesystem setup, error handling...
templates := template.New("base")
templates = templates.Funcs(template.FuncMap{
    "json": func(v interface{}) string { /* ... */ },
    // ... more functions
})
staticFS, err := fs.Sub(assets.StaticFiles, "static")
// ... more setup code
```

**New way**:
```go
srv, err := webserver.New()
if err != nil {
    log.Fatal(err)
}

srv.AddTemplatesFromFS(myTemplates, "templates")
log.Fatal(http.ListenAndServe(":8080", srv.Router()))
```

This reduces setup code by ~80% and eliminates common filesystem path bugs.

## Development Notes

### Key Function: RenderPageWithContent()
Use this instead of direct template execution to enable dynamic content injection:

```go
// ✅ Correct - enables dynamic template resolution
srv.RenderPageWithContent(w, "dashboard-content", data)

// ❌ Avoid - requires hardcoded template names
srv.RenderTemplate(w, "layout", data)
```

### Embedded Assets
Static files are served from `/static/` path and include:
- CSS: Bulma, FontAwesome, custom components
- JS: Core functionality, theme switching, loading overlays
- Fonts: FontAwesome webfonts

### Template Organization
- **Base templates**: Provided by webserver-core-ui
- **Content templates**: Your application-specific content
- **Includes**: Reusable template fragments

## Examples

See `cmd/theme-test/` for a complete working implementation demonstrating all features.
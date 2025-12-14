package assets

import "embed"

// ⚠️ IMPORTANT: Files are embedded at COMPILE TIME!
// After changing ANY file in static/ or templates/ directories,
// you MUST rebuild the binary:
//   pkill -f theme-test && go build -a -o theme-test ./cmd/theme-test
// Otherwise, the server will continue serving OLD versions of the files!

//go:embed static
var StaticFiles embed.FS

//go:embed templates
var TemplateFiles embed.FS

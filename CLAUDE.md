# CLAUDE.md - ForgeXML SCAP Development Guide

## ⚠️ CRITICAL: NEVER USE GO REPLACE DIRECTIVES ⚠️
**NEVER add `replace` directives to go.mod files. This is absolutely prohibited.**

When updating dependencies across repositories:
1. Use `crossrepo` tool with `operation='release'` to tag a new version of the source repo
2. Use `crossrepo` tool with `operation='update'` to propagate changes to dependent repos
3. The crossrepo tool handles: git commit, git push, tagging, go.mod updates, vendor sync

Replace directives break the dependency chain and cause issues for other developers.

## Project Overview

ForgeXML-SCAP is a generated Go application for working with SCAP (Security Content Automation Protocol) XML documents. It includes:
- Generated Go structs from XSD schemas (ARF, OVAL, XCCDF, CPE)
- A web UI for editing XML documents
- DAL (Data Access Layer) for database persistence

## Key Directories

- `cmd/ui/` - Generated web UI application
- `internal/generated/` - Generated Go structs from XSD schemas
- `internal/dal/` - Generated Data Access Layer
- `schemas/` - Source XSD schema files
- `instances/` - Sample XML files for testing

## Regenerating Code

To regenerate all code from the XSD schemas:
```bash
cd /home/mmcnew/repos/forgexml
go run ./cmd/forgexml -config /home/mmcnew/repos/forgexml-scap/forgexml.json
```

## Building the UI

```bash
cd /home/mmcnew/repos/forgexml-scap
go build -o ./bin/ui ./cmd/ui
./bin/ui -port 8080
```

## Dependencies

This project depends on:
- `github.com/aequo-labs/webserver-core-ui` - Web UI framework
- `github.com/aequo-labs/webserver-core` - Core web server functionality

When these dependencies are updated, use `crossrepo` to propagate changes - NEVER use replace directives.

## Testing

```bash
# Run all tests
go test ./...

# Run with coverage
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
```

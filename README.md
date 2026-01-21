# Scaffold

Scaffold is a Go-based command-line tool that helps you quickly generate a clean, ready-to-use project structure so you can start building immediately instead of setting things up manually.

Link to project: https://github.com/tskaushal/scaffold

alt tag:
Screenshot of the scaffold CLI generating a new project from the terminal

## How It's Made:

**Tech used:** Go, Cobra CLI

Scaffold is built as a developer tooling project using Go. It uses Cobra to handle command-line commands and flags, and Go’s filesystem APIs to create directories and generate starter files dynamically.

The core idea is simple: instead of manually creating folders, boilerplate files, and configuration every time, Scaffold automates the entire process. The CLI takes arguments like project name and type, then writes the appropriate structure and code to disk.

The project also helped me understand how real scaffolding tools work internally — generating code as templates, handling paths correctly across operating systems, and packaging a CLI so it can be installed and used globally via `go install`.

## Optimizations (optional):

- Improved error handling to give clearer CLI feedback
- Structured the project using `cmd/` and `internal/` for clean separation
- Designed the CLI so new templates (web, API, etc.) can be added easily in the future



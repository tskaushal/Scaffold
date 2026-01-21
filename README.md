# Scaffold

Scaffold is a Go-based command-line tool that helps you quickly generate a clean, ready-to-use project structure so you can start building immediately without manual setup.

Link to project: https://github.com/tskaushal/scaffold

alt tag:
Screenshot showing the scaffold CLI generating a new project from the terminal

How It's Made:

Tech used: Go, Cobra CLI

Scaffold is built using Go as a developer tooling project. It uses the Cobra CLI library to handle commands, subcommands, and flags, and Go’s filesystem APIs to create directories and generate starter files.

The tool works by taking user input such as the project name and type, then programmatically creating the required folder structure and boilerplate code. The generated project can then be run or extended by the user. This project was built to understand how real scaffolding tools like create-react-app or cobra init work internally.

How To Install:

Make sure Go (1.21 or later) is installed on your system.

Install the CLI using Go:
go install github.com/tskaushal/scaffold/cmd/scaffold@latest

Ensure that `$GOPATH/bin` is added to your PATH.

Verify installation:

scaffold --help


How To Use:

Create a new project:


scaffold create myapp


Create a CLI-type project:


scaffold create water --type cli


Navigate into the generated project and run it:


cd myapp
go run ./cmd


Optimizations

(optional)

The project is structured using `cmd/` and `internal/` directories for clean separation of concerns. The CLI is designed so that new templates (web, API, etc.) can be added easily in the future without modifying the core logic.

Lessons Learned:

Building developer tools is very different from building applications. The scaffolding code runs once, while the generated code runs later. I also learned how important correct module naming, Windows PATH handling, and proper project structure are when publishing a real Go CLI that others can install and use.

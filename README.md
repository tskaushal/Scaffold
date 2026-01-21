# Scaffold
[![Ask DeepWiki](https://devin.ai/assets/askdeepwiki.png)](https://deepwiki.com/tskaushal/Scaffold)

Scaffold is a Go-based command-line tool that helps you quickly generate a clean, ready-to-use project structure so you can start building immediately without manual setup.

This project was built to understand how real scaffolding tools like `create-react-app` or `cobra init` work internally.

## Features

Scaffold provides ready-to-use templates for various types of Go applications:

*   **CLI (`cli`):** Generates a command-line application using the [Cobra](https://github.com/spf13/cobra) framework. Includes a basic command structure, Makefile, and `go.mod` file.
*   **Web (`web`):** Creates a simple web server project with a standard layout for handlers, static assets (CSS, JS), and HTML templates.
*   **API (`api`):** Scaffolds a REST API structure including handlers, models, and middleware for common tasks like logging and CORS.

## Installation

Make sure you have Go (version 1.24 or later) installed on your system.

1.  **Install the CLI:**
    ```sh
    go install github.com/tskaushal/scaffold/cmd/scaffold@latest
    ```

2.  **Verify Setup:**
    Ensure that `$GOPATH/bin` is in your system's `PATH`. You can verify the installation by running:
    ```sh
    scaffold --help
    ```

## Usage

Use the `create` command to generate a new project. You can specify the project type using the `--type` (or `-t`) flag. If no type is specified, it defaults to `cli`.

```sh
scaffold create <project-name> --type <project-type>
```

### Examples

**Create a CLI Application:**
```sh
scaffold create my-cli-app --type cli
```
Or simply (as `cli` is the default):
```sh
scaffold create my-cli-app
```

**Create a Web Application:**
```sh
scaffold create my-webapp --type web
```

**Create a REST API:**
```sh
scaffold create my-api --type api
```

### Running a Generated Project

After creating a project, navigate into its directory and run it.

For **Web** and **API** projects:
```sh
cd <project-name>
go run ./cmd/main.go
```

For **CLI** projects, you can use the generated Makefile:
```sh
cd <project-name>
make run

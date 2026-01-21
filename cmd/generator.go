package cmd

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateProject(name, projectType string) error {

	if _, err := os.Stat(name); err == nil {
		return fmt.Errorf("folder '%s' already exists", name)
	}

	fmt.Printf("Creating project folder: %s\n", name)
	err := os.MkdirAll(name, 0755)
	if err != nil {
		return fmt.Errorf("failed to create project folder: %w", err)
	}

	switch projectType {
	case "cli":
		err = createCLIProject(name)
	case "web":
		err = createWebProject(name)
	case "api":
		err = createAPIProject(name)
	default:
		err = createCLIProject(name)
	}

	if err != nil {
		os.RemoveAll(name)
		return err
	}

	return nil
}

// creates a CLI project structure
func createCLIProject(name string) error {
	folders := []string{
		filepath.Join(name, "cmd"),
		filepath.Join(name, "internal"),
	}

	for _, folder := range folders {
		os.MkdirAll(folder, 0755)
	}

	mainGo := `package main

import "` + name + `/cmd"

func main() {
	cmd.Execute()
}
`
	os.WriteFile(filepath.Join(name, "main.go"), []byte(mainGo), 0644)

	rootGo := `package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "` + name + `",
	Short: "` + name + ` - A CLI application",
	Long:  "A command line application built with Go and Cobra",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
`
	os.WriteFile(filepath.Join(name, "cmd", "root.go"), []byte(rootGo), 0644)

	versionCmd := `package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("` + name + ` version 1.0.0")
	},
}
`
	os.WriteFile(filepath.Join(name, "cmd", "version.go"), []byte(versionCmd), 0644)

	makefile := `build:
	go build -o ` + name + ` .

run:
	go run main.go

test:
	go test ./...

clean:
	rm -f ` + name + `

install:
	go install

.PHONY: build run test clean install
`
	os.WriteFile(filepath.Join(name, "Makefile"), []byte(makefile), 0644)

	gitignore := `*.exe
*.exe~
*.dll
*.so
*.dylib
*.test
*.out
vendor/
go.work
` + name + `
`
	os.WriteFile(filepath.Join(name, ".gitignore"), []byte(gitignore), 0644)

	readme := `# ` + name + `

A CLI application built with Go and Cobra.

## Installation

` + "```bash" + `
go install
` + "```" + `

## Usage

` + "```bash" + `
` + name + ` --help
` + name + ` version
` + "```" + `

## Build

` + "```bash" + `
make build
` + "```" + `

## Development

` + "```bash" + `
make run
` + "```" + `
`
	os.WriteFile(filepath.Join(name, "README.md"), []byte(readme), 0644)

	gomod := `module ` + name + `

go 1.21

require github.com/spf13/cobra v1.8.0

require (
	github.com/inconshreveable/mousetrap v1.1.0
	github.com/spf13/pflag v1.0.5
)
`
	os.WriteFile(filepath.Join(name, "go.mod"), []byte(gomod), 0644)

	return nil
}

// createWebProject creates a web server project structure
func createWebProject(name string) error {
	folders := []string{
		filepath.Join(name, "cmd"),
		filepath.Join(name, "handlers"),
		filepath.Join(name, "static", "css"),
		filepath.Join(name, "static", "js"),
		filepath.Join(name, "templates"),
	}

	for _, folder := range folders {
		os.MkdirAll(folder, 0755)
	}

	mainGo := `package main

import (
	"log"
	"net/http"
	"` + name + `/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	
	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
`
	os.WriteFile(filepath.Join(name, "cmd", "main.go"), []byte(mainGo), 0644)

	homeHandler := `package handlers

import (
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	data := map[string]string{
		"Title": "` + name + `",
		"Message": "Welcome to your web application!",
	}
	tmpl.Execute(w, data)
}
`
	os.WriteFile(filepath.Join(name, "handlers", "home.go"), []byte(homeHandler), 0644)

	indexHTML := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>
    <link rel="stylesheet" href="/static/css/style.css">
</head>
<body>
    <div class="container">
        <h1>{{.Title}}</h1>
        <p>{{.Message}}</p>
    </div>
    <script src="/static/js/main.js"></script>
</body>
</html>
`
	os.WriteFile(filepath.Join(name, "templates", "index.html"), []byte(indexHTML), 0644)

	styleCSS := `* {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
}

body {
    font-family: Arial, sans-serif;
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    min-height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
}

.container {
    background: white;
    padding: 40px;
    border-radius: 10px;
    box-shadow: 0 10px 30px rgba(0,0,0,0.3);
    text-align: center;
}

h1 {
    color: #667eea;
    margin-bottom: 20px;
}

p {
    color: #666;
    font-size: 18px;
}
`
	os.WriteFile(filepath.Join(name, "static", "css", "style.css"), []byte(styleCSS), 0644)

	mainJS := `console.log('` + name + ` loaded');

document.addEventListener('DOMContentLoaded', function() {
    console.log('Page ready');
});
`
	os.WriteFile(filepath.Join(name, "static", "js", "main.js"), []byte(mainJS), 0644)

	gitignore := `*.exe
*.dll
*.so
*.dylib
*.test
*.out
vendor/
go.work
` + name + `
`
	os.WriteFile(filepath.Join(name, ".gitignore"), []byte(gitignore), 0644)

	readme := `# ` + name + `

A web application built with Go.

## Run

` + "```bash" + `
cd cmd
go run main.go
` + "```" + `

Visit: http://localhost:8080

## Project Structure

` + "```" + `
` + name + `/
├── cmd/
│   └── main.go
├── handlers/
│   └── home.go
├── static/
│   ├── css/
│   └── js/
└── templates/
    └── index.html
` + "```" + `
`
	os.WriteFile(filepath.Join(name, "README.md"), []byte(readme), 0644)

	gomod := `module ` + name + `

go 1.21
`
	os.WriteFile(filepath.Join(name, "go.mod"), []byte(gomod), 0644)

	return nil
}

// createAPIProject creates a REST API project structure
func createAPIProject(name string) error {
	folders := []string{
		filepath.Join(name, "cmd"),
		filepath.Join(name, "handlers"),
		filepath.Join(name, "models"),
		filepath.Join(name, "middleware"),
	}

	for _, folder := range folders {
		os.MkdirAll(folder, 0755)
	}

	mainGo := `package main

import (
	"log"
	"net/http"
	"` + name + `/handlers"
	"` + name + `/middleware"
)

func main() {
	mux := http.NewServeMux()
	
	mux.HandleFunc("/api/health", handlers.HealthHandler)
	mux.HandleFunc("/api/users", handlers.UsersHandler)
	
	handler := middleware.Logger(mux)
	handler = middleware.CORS(handler)
	
	log.Println("API Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
`
	os.WriteFile(filepath.Join(name, "cmd", "main.go"), []byte(mainGo), 0644)

	healthHandler := `package handlers

import (
	"encoding/json"
	"net/http"
)

type HealthResponse struct {
	Status  string ` + "`json:\"status\"`" + `
	Message string ` + "`json:\"message\"`" + `
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	response := HealthResponse{
		Status:  "ok",
		Message: "` + name + ` API is running",
	}
	
	json.NewEncoder(w).Encode(response)
}
`
	os.WriteFile(filepath.Join(name, "handlers", "health.go"), []byte(healthHandler), 0644)

	usersHandler := `package handlers

import (
	"encoding/json"
	"net/http"
	"` + name + `/models"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	users := []models.User{
		{ID: 1, Name: "John Doe", Email: "john@example.com"},
		{ID: 2, Name: "Jane Smith", Email: "jane@example.com"},
	}
	
	json.NewEncoder(w).Encode(users)
}
`
	os.WriteFile(filepath.Join(name, "handlers", "users.go"), []byte(usersHandler), 0644)

	userModel := `package models

type User struct {
	ID    int    ` + "`json:\"id\"`" + `
	Name  string ` + "`json:\"name\"`" + `
	Email string ` + "`json:\"email\"`" + `
}
`
	os.WriteFile(filepath.Join(name, "models", "user.go"), []byte(userModel), 0644)

	loggerMiddleware := `package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Completed in %v", time.Since(start))
	})
}
`
	os.WriteFile(filepath.Join(name, "middleware", "logger.go"), []byte(loggerMiddleware), 0644)

	corsMiddleware := `package middleware

import "net/http"

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next.ServeHTTP(w, r)
	})
}
`
	os.WriteFile(filepath.Join(name, "middleware", "cors.go"), []byte(corsMiddleware), 0644)

	gitignore := `*.exe
*.dll
*.so
*.dylib
*.test
*.out
vendor/
go.work
` + name + `
`
	os.WriteFile(filepath.Join(name, ".gitignore"), []byte(gitignore), 0644)

	readme := `# ` + name + `

A REST API built with Go.

## Run

` + "```bash" + `
cd cmd
go run main.go
` + "```" + `

## Endpoints

` + "```bash" + `
GET  /api/health  - Health check
GET  /api/users   - Get all users
` + "```" + `

## Test

` + "```bash" + `
curl http://localhost:8080/api/health
curl http://localhost:8080/api/users
` + "```" + `

## Project Structure

` + "```" + `
` + name + `/
├── cmd/
│   └── main.go
├── handlers/
│   ├── health.go
│   └── users.go
├── models/
│   └── user.go
└── middleware/
    ├── logger.go
    └── cors.go
` + "```" + `
`
	os.WriteFile(filepath.Join(name, "README.md"), []byte(readme), 0644)

	gomod := `module ` + name + `

go 1.21
`
	os.WriteFile(filepath.Join(name, "go.mod"), []byte(gomod), 0644)

	return nil
}

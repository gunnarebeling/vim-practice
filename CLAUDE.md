# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Purpose

This is a **LazyVim/Neovim learning playground** built around a simple Go + HTMX task manager application. The primary goal is practicing Neovim workflows, not building production software.

## Build & Run Commands

```bash
cd task-manager
go run main.go          # Start server at http://localhost:8080
go build -o taskmanager # Build binary
```

## Project Structure

```
nvim-playground/
├── task-manager/       # Go + HTMX task manager app
│   ├── main.go         # Server, routes, TaskStore
│   ├── templates/      # HTML templates with HTMX
│   └── static/         # CSS
└── learning/           # Neovim practice curriculum
    ├── progress.md     # Progress tracking log
    └── chapters/       # Lesson chapters with exercises
```

## Architecture

- **Backend**: Single `main.go` with in-memory TaskStore (mutex-protected)
- **Frontend**: HTMX for dynamic updates without JavaScript
- **Routes**: `/` (index), `/tasks` (POST), `/tasks/toggle/{id}` (PUT), `/tasks/delete/{id}` (DELETE)

## Learning System

The `/learning` directory contains structured Neovim practice:
- Each chapter focuses on specific Neovim/LazyVim skills
- Chapters end with timed challenge tasks
- Progress is logged in `progress.md`

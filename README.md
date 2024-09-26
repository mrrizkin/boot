# Boot

## Overview

`boot` is a web application project that uses a combination of Go for backend development and Vite for managing frontend assets. The project also leverages several popular libraries and tools such as TailwindCSS for styling, Alpine.js for frontend reactivity, and HTMX for handling modern interactions without the need for a full frontend framework.

## Table of Contents

- [Getting Started](#getting-started)
- [Scripts](#scripts)
- [Dependencies](#dependencies)
  - [Development Dependencies](#development-dependencies)
  - [Runtime Dependencies](#runtime-dependencies)
- [Build and Development](#build-and-development)

## Getting Started

To start developing or building the project, ensure that you have the following tools installed:

- Node.js
- PNPM (Package Manager)
- Go

### Install Dependencies

Run the following command to install both development and runtime dependencies:

```bash
pnpm install
```

## Scripts

The project contains several scripts for development and production builds:

### Development Scripts

- `pnpm dev:assets`: Runs Vite in development mode to serve and watch the frontend assets.
- `pnpm dev:app`: Starts the backend application using `air` for live-reloading in Go.
- `pnpm dev`: Runs both the frontend and backend in parallel using `concurrently`. It waits 5 seconds before starting the backend to ensure that the frontend assets are ready.

### Build Scripts

- `pnpm build:assets`: Builds the frontend assets using Vite for production.
- `pnpm build:app`: Compiles the Go backend application located in `./cmd/main/main.go`.
- `pnpm build`: Executes both `build:assets` and `build:app` scripts in sequence to create a production build.

## Dependencies

### Development Dependencies

The following packages are used for development purposes:

- `autoprefixer`: Automatically adds vendor prefixes to CSS.
- `concurrently`: Runs multiple commands concurrently (e.g., both frontend and backend servers).
- `postcss`: A tool for transforming CSS with JavaScript plugins.
- `tailwindcss`: A utility-first CSS framework.
- `vite`: A fast frontend build tool and dev server.
- `vite-plugin-backend`: A Vite plugin to work with backend integrations.
- `vite-plugin-full-reload`: A Vite plugin that enables full reload on backend changes.

### Runtime Dependencies

These are the main runtime dependencies for the project:

- `alpinejs`: A lightweight JavaScript framework for declarative rendering.
- `axios`: A promise-based HTTP client for the browser and Node.js.
- `htmx.org`: A modern library for simplifying dynamic web page interactions without using JavaScript frameworks.

## Build and Development

To start developing, run the following command to serve both the frontend and backend in parallel:

```bash
pnpm dev
```

To create a production build:

```bash
pnpm build
```

This will generate optimized frontend assets and compile the Go backend into a binary.

# Boot

## Overview

Boot is a sophisticated, full-stack web application boilerplate that leverages the power of Go for backend development and modern frontend technologies. It's designed to provide a robust starting point for building scalable, high-performance web applications. Currently, Boot uses Gofiber as its primary web framework, with plans to potentially support Chi and Echo in the future.

## Key Features

- **Go-powered Backend**: Utilizes the Gofiber framework for high-performance HTTP serving
- **Modern Frontend Tooling**: Incorporates Vite for efficient asset bundling and hot module replacement
- **Responsive Styling**: Integrates TailwindCSS for utility-first styling
- **Dynamic UI**: Employs Alpine.js for lightweight reactivity and HTMX for seamless server interactions
- **Flexible Database Support**: Configurable with MySQL, PostgreSQL, or SQLite via GORM
- **Advanced Logging**: Implements structured logging with zerolog
- **Environment Management**: Supports easy configuration via .env files
- **Server-side Rendering**: Utilizes the Gonja templating engine for efficient HTML generation
- **Form Validation**: Incorporates request validation using go-playground/validator
- **Session Management**: Built-in support for various session storage options
- **Asset Bundling**: Custom Vite integration for optimized asset management
- **Extensible Architecture**: Modular design allowing for easy customization and extension

## Project Structure

```
.
├── app/
│   ├── domains/              # Business logic implementation
│   ├── handlers/             # HTTP request handling and payload validation
│   ├── models/               # Database schema definitions
│   ├── types/                # Application-level type definitions
│   └── utils/                # Application-specific utilities
├── cmd/
│   └── main/                 # Application entry point
├── public/                   # Static file serving
├── resources/
│   ├── assets/               # Frontend assets (CSS, JS, web components)
│   └── views/                # Gonja templates for server-side rendering
├── routes/                   # Routing and middleware definitions
├── storage/                  # Application storage (e.g., logs)
├── system/                   # Core system components
│   ├── config/               # Configuration management
│   ├── database/             # Database connection handling
│   ├── error/                # System-level error definitions
│   ├── server/               # Server bootstrapping
│   ├── session/              # Session management
│   ├── validator/            # Request validation
│   └── view/                 # View rendering and custom template functions
└── third-party/              # Third-party integrations
    ├── asset-bundler/        # Custom Vite integration
    ├── hashing/              # Password hashing utilities
    ├── logger/               # Logging configuration
    └── scheduler/            # Task scheduling
```

## Getting Started

### Prerequisites

- Go (version 1.22.2 or later)
- Node.js
- PNPM (Package Manager)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/mrrizkin/boot.git
   cd boot
   ```

2. Install dependencies:
   ```bash
   pnpm run setup
   ```

3. Configure your environment:
   - Copy `.env.example` to `.env`
   - Adjust settings in `.env` as needed

4. Start the development server:
   ```bash
   pnpm run dev
   ```

## Development Workflow

- `pnpm run dev`: Concurrently starts the Vite dev server and Go backend with hot-reloading
- `pnpm run build`: Builds frontend assets and compiles the Go backend for production

## Customization

Boot is designed with extensibility in mind:

1. **Backend Logic**: Implement your business logic in `app/domains/`
2. **API Endpoints**: Define new routes in `routes/` and implement handlers in `app/handlers/`
3. **Database Schema**: Add or modify models in `app/models/`
4. **Frontend Assets**: Manage your CSS and JavaScript in `resources/assets/`
5. **Views**: Create or edit Gonja templates in `resources/views/`
6. **Middleware**: Implement custom middleware in `routes/middleware/`
7. **Configuration**: Adjust application settings in `system/config/`

## Best Practices

- Follow Go's official [style guide](https://golang.org/doc/effective_go) and [code review comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Utilize the domain-driven design principles in the `app/domains/` directory
- Leverage the built-in error handling and logging mechanisms for consistent error management
- Use the provided validator for all incoming request payloads
- Implement unit tests for your business logic and integration tests for your API endpoints

## Contributing

We welcome contributions to Boot. Please read our [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
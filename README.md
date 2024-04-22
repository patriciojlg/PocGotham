# POC Go Development Stack with Chi, Tailwind, htmx, and Templ

This is a Proof of Concept (POC) project showcasing the Go development stack using Chi router, Tailwind CSS, htmx, and Templ templating engine.

## Description

This project aims to demonstrate the integration and usage of the following technologies:
- Chi: A lightweight, idiomatic and composable router for building Go HTTP services.
- Tailwind CSS: A utility-first CSS framework for rapidly building custom designs.
- htmx: A modern approach to AJAX, providing seamless and easy-to-use interactions.
- Templ: A templating engine for Go, offering the flexibility and power needed for templating tasks.

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/your-repo.git
```
2. Install dependencies:


```bash

go mod tidy
```

3. Run the project:

```bash

go run main.go
```
## Pending Task

    Create a settings.go file to define constants and configuration settings for the project.
### TODO
    [ ] Set endpoints names with const on settings.go
    [ ] Add an error handler that allows adding validators that return error/nil. Implement them by appending them to a list of functions through a method.
    [ ] Add minimal structure of the service package for consuming AWS Cognito.

## Contributing

Feel free to contribute by forking the repository and submitting pull requests.
## License

This project is licensed under the MIT License - see the LICENSE.md file for details.

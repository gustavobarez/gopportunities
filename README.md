# Gopportunities

![GOLANG_BADGE](https://img.shields.io/badge/Golang-%2300FFFF?style=for-the-badge&logo=go)
![GIN_BADGE](https://img.shields.io/badge/Gin-%2300FFFF.svg?style=for-the-badge&logo=gin&logoColor=black)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-%23396c94.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![DOCKER_COMPOSE](https://img.shields.io/badge/Docker%20Compose-%231d63ed.svg?style=for-the-badge&logo=docker&logoColor=white)

💼 Golang job opportunities API with Go-Gin router, PostgresSQL database, GoORM integration, Swagger documentation, and organized package structure for recruitment platform management.

---

## Features

- Modern job opportunities API built with Golang
- RESTful architecture using Go-Gin framework for high-performance routing
- SQLite database for lightweight and efficient data storage
- GoORM integration for seamless database operations
- Complete CRUD operations for job opportunities management
- Integrated Swagger documentation for API testing and exploration
- Clean architecture with proper separation of concerns
- Automated testing suite to ensure code quality and reliability
- Modern package structure for maintainable codebase
- Environment-based configuration management
- Error handling and validation for production reliability

## Installation

To use this project, you need to follow these steps:

1. Clone the repository: `git clone https://github.com/gustavobarez/gopportunities.git`
2. Install the dependencies: `go mod download`
3. Build the application: `go build`
4. Run the application: `./main`

## Makefile Commands

The project includes a Makefile to help you manage common tasks more easily. Here's a list of the available commands and a brief description of what they do:

- `make run`: Run the application without generating API documentation.
- `make run-with-docs`: Generate the API documentation using Swag, then run the application.
- `make build`: Build the application and create an executable file named `gopportunities`.
- `make test`: Run tests for all packages in the project.
- `make docs`: Generate the API documentation using Swag.
- `make clean`: Remove the `gopportunities` executable and delete the `./docs` directory.

To use these commands, simply type `make` followed by the desired command in your terminal. For example:

```sh
make run
```

## Used Tools

This project uses the following tools:

- [Golang](https://golang.org/) for backend development
- [Go-Gin](https://github.com/gin-gonic/gin) for route management
- [GoORM](https://gorm.io/) for database communication
- [PostgreSQL](https://www.postgresql.org/) for database storage
- [Swagger](https://swagger.io/) for API documentation and testing
- [Docker](https://www.docker.com/) for containerization

## Usage

After the API is running, you can use the Swagger UI to interact with the endpoints for searching, creating, editing, and deleting job opportunities. The API can be accessed at `http://localhost:$PORT/swagger/index.html`.

Default $PORT if not provided=8080.

## Contributing

To contribute to this project, please follow these guidelines:

1. Fork the repository
2. Create a new branch: `git checkout -b feature/your-feature-name`
3. Make your changes and commit them using Conventional Commits
4. Push to the branch: `git push origin feature/your-feature-name`
5. Submit a pull request

---

## License

This project is licensed under the MIT License - see the LICENSE.md file for details.

## Credits

This project was created by [Gustavo Barez](https://github.com/gustavobarez).

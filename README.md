# Go: Developing a REST API - Gin and GORM

## Acknowledgements

- This is a project of creating a a REST API using Gin and GORM, connect it to a PostgreSQL database using a Docker image.

## Authors

- [Alura](https://cursos.alura.com.br/formacao-go)

## Requirements

- Golang
- PostgreSQL
- Gin
- GORM
- Docker
- Postman

## Environment Variables

To run this project, you will need to do the following:

- docker-compose up
- go run main.go

## Features

- Access from postman
- Create / update / edit / delete an item from a db using the frontend

## Demo

- Postgre database created via Docker
  ![image](https://github.com/tiagoc0sta/class33/assets/63982700/4d4ec76a-e39e-491a-9e12-fb75afdfc0ed)

- Perform CRUD tests via Postman
  ![image](https://github.com/tiagoc0sta/class33/assets/63982700/3e265af4-ba2e-46bd-a4fa-265a1e62999d)

## Tests

- Implement API testing
- run on terminal: go test
- Testify = https://github.com/stretchr/testify

# Golang: Website Monitoring

Lessons Learned:

- Understood the importance of structuring the API endpoints and handling HTTP requests and responses efficiently.

Integrating with a Database

- Learned how to integrate a Go API with a database using Docker.
- Gained knowledge of using Docker for containerization and database management.
  Using GORM

- Learned how to use GORM, Go's popular ORM (Object-Relational Mapping) library, for database operations.
- Explored features such as model definitions, querying, and relationships management provided by GORM.

Creating Middleware

- Learned how to create middleware in Go to handle cross-cutting concerns such as logging, authentication, and request validation.
- Understood the benefits of using middleware to avoid code duplication and maintain clean API logic.

Integrating with a React Frontend

- Learned how to integrate a Go API with a React frontend for a full-stack application.
- Explored methods for communication between the frontend and backend, such as RESTful APIs and WebSocket.

## Installation

1. Open Docker Desktop

2. Run Docker

```bash
  docker-compose up
```

3. Run API

```bash
  run main.go
```

4. Open Postman and perform CRUD operations on API. See the reponse on the frontend

### Pages:

- http://localhost:8080/vehicles - see vehicles

- http://localhost:54321 = login to pgAdmin - PostgreSQL

## Notes:

### Gin framework- to create API - site: https://github.com/gin-gonic/gin

### Gorm ORM

- (Object–relational mapping
- avoid the need for queries) utilizado para conectar com o DB
- site: https://gorm.io/docs/index.html

### PostgreSQL

- Store data

### Docker - create two images:

- 1 postgres
- 2 pgadmin-compose

---

### Validator - V2 - https://pkg.go.dev/gopkg.in/validator.v2

### gitignore - Explain importance

### Render page on the FrontEnd

### Swagger - Framework for API documentation / testing

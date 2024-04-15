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

### Postgre database created via Docker
- http://localhost:54321 = login to pgAdmin - PostgreSQL
  ![image](https://github.com/tiagoc0sta/class33/assets/63982700/4d4ec76a-e39e-491a-9e12-fb75afdfc0ed)

### Postman - Perform CRUD tests via
  ![image](https://github.com/tiagoc0sta/class33/assets/63982700/3e265af4-ba2e-46bd-a4fa-265a1e62999d)

### View all API data on JSON format
- http://localhost:8080/vehicles - see vehicles (json)

![image](https://github.com/tiagoc0sta/api-go-gin/assets/63982700/849edbb7-77df-4ac8-b69e-c746518141be)

### Swagger - Framework for API documentation / testing
- http://localhost:8080/swagger/index.html 

![image](https://github.com/tiagoc0sta/api-go-gin/assets/63982700/dc73cebb-134c-41ce-ae20-2b9befc2eb43)

- Command to edit swagger: $ swag init --parseDependency --parseInternal --parseDepth 1
- Swagger official documentation: https://github.com/swaggo/gin-swagger

### Frontend
- http://localhost:8080/index 
![image](https://github.com/tiagoc0sta/api-go-gin/assets/63982700/31320394-7dbb-46a5-aad4-6dbdb5cb33ee)

## Tests

- Implement automated testing
- run on terminal: go test
![image](https://github.com/tiagoc0sta/api-go-gin/assets/63982700/87424b7e-1369-42de-9457-a9213d86384d)

- Testify official documentation: https://github.com/stretchr/testify

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




## Notes:

### Gin framework- to create API - Official site: https://github.com/gin-gonic/gin

### Gorm ORM

- (Object–relational mapping - avoid the need for queries) used to connect to DB
- site: https://gorm.io/docs/index.html

### PostgreSQL

- Store data

### Docker - create two images:

- 1 postgres
- 2 pgadmin-compose

---

### Validator - V2 - https://pkg.go.dev/gopkg.in/validator.v2

### gitignore file - Include all .yml filmes in gitignore in order to avoid sharing them to the cloud.

![image](https://github.com/tiagoc0sta/api-go-gin/assets/63982700/1fb13b62-61c2-4c9c-ab04-dded88dc19a5)


### Render page on the FrontEnd





### Github - provide project link
- https://github.com/tiagoc0sta/api-go-gin

### explain MVC model - Model View Controller

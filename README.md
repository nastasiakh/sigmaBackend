# Getting Started with Sigma Golang app 

This project is Fullstack Developer Pre-Interview Assignment

## Project structure 
```
├── Dockerfile
├── internal
│   ├── entities
│   │   └── user.go
│   └── handlers
│       ├── controllers
│       │   └── user.controller.go
│       ├── di
│       │   └── container.go
│       ├── middlewares
│       │   └── corsMiddleware.go
│       ├── repositories
│       │   └── user.repository.go
│       ├── routes
│       │   └── routes.go
│       └── services
│           └── user.services.go
├── main.go
```

## How to run

### `docker build -t sigma:latest .`
### `docker run -p 8080:8080 sigma/user-crud:latest`

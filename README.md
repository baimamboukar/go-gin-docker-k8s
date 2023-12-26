# Go-Gin-Docker-K8s
[![CI/CD Workflow](https://github.com/baimamboukar/go-gin-docker-k8s/actions/workflows/main.yaml/badge.svg)](https://github.com/baimamboukar/go-gin-docker-k8s/actions/workflows/main.yaml) [![Built with Go](https://img.shields.io/badge/Built%20with-Go-1f425f.svg)](https://golang.org/)[![Go Language](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)](https://golang.org/)

A backend example written in Go and the Gin web framework for startups and tech companies.

## Overview

This project serves as a simple backend example for managing startups and tech companies. It is built using the Go programming language with the Gin web framework. The application utilizes a Vercel-hosted PostgreSQL database, and Gorm is employed for managing database services. Dockerization and deployment are handled through AWS EKS, and Kubernetes configurations are provided in the `kubernetes` folder. The project includes authentication and logging middlewares.

## Table of Contents

1. [Installation](#installation)
2. [Configuration](#configuration)
3. [Usage](#usage)
4. [Middlewares](#middlewares)
   - [Auth Middleware](#auth-middleware)
   - [Logger Middleware](#logger-middleware)
5. [Kubernetes](#kubernetes)
6. [Makefile](#makefile)
7. [VS Code Extensions](#vs-code-extensions)
8. [CI/CD Workflow](#cicd-workflow)
9. [API Documentation](#api-documentation)
10. [Contributing](#contributing)
11. [License](#license)
12. [Acknowledgments](#acknowledgments)


### Tools and Technologies
- Go
- Gin
- Gorm
- PostgreSQL
- Docker
- Kubernetes
- AWS EKS
- Vercel
- Postman
- VS Code

## Installation

Follow these steps to set up the project locally.


- Clone the repository
```bash
git clone https://github.com/baimamboukar/go-gin-docker-k8s.git
```
- Navigate to the project directory
```bash
cd go-gin-docker-k8s
```

### Install dependencies
```bash
make install
```

### Configuration

The project uses environment variables for configuration. Copy the provided .env.example file to .env and update the values accordingly.

```bash
cp .env.example .env.local
# Update .env file with appropriate values
```

### Usage
Run the following command to start the application locally.

```bash
make run
```

### Middlewares
- Auth Middleware
- Logger Middleware


### Kubernetes Configs
kubectl apply -f kubernetes/



### Makefile
The Makefile includes the following commands...
make lint
make test


### VS Code Extensions
To enhance the development experience, consider installing the following VS Code extensions...

### CI/CD Workflow
The project has a continuous integration workflow...

### API Documentation
The API documentation is maintained using Postman. Click the button below to view and run the API documentation in Postman.


### Contributing
...

### License
This project is licensed under the MIT License.



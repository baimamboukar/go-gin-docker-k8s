# Go-Gin-Docker-K8s
[![CI/CD Workflow](https://github.com/baimamboukar/go-gin-docker-k8s/actions/workflows/main.yaml/badge.svg)](https://github.com/baimamboukar/go-gin-docker-k8s/actions/workflows/main.yaml)

[![Go Language](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)](https://golang.org/) • [![Postman](https://img.shields.io/badge/Postman-FF6C37?style=flat&logo=postman&logoColor=white)](YOUR_POSTMAN_API_COLLECTION_URL) • [![Vercel](https://img.shields.io/badge/Vercel-000000?style=flat&logo=vercel&logoColor=white)](YOUR_VERCEL_DEPLOYMENT_URL) • [![PostgreSQL](https://img.shields.io/badge/PostgreSQL-336791?style=flat&logo=postgresql&logoColor=white)](https://www.postgresql.org/) •[![AWS](https://img.shields.io/badge/AWS-232F3E?style=flat&logo=amazon-aws&logoColor=white)](https://aws.amazon.com/) • [![Kubernetes](https://img.shields.io/badge/Kubernetes-326CE5?style=flat&logo=kubernetes&logoColor=white)](https://kubernetes.io/) • [![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white)](https://www.docker.com/) • [![Prometheus](https://img.shields.io/badge/Prometheus-DA4226?style=flat&logo=prometheus&logoColor=white)](PROMETHEUS_LINK) • [![Grafana](https://img.shields.io/badge/Grafana-F46800?style=flat&logo=grafana&logoColor=white)](GRAFANA_LINK)




A backend example written in Go and the Gin web framework. The project includes a Dockerfile and Kubernetes configurations for deployment.

## Overview

This project serves as a simple backend example. It is built using the **Go programming language** with the **Gin web framework**, **Dockerized** and orchestrated using **K8s**, and delpoyed to **AWS ECS**. This backend utilizes a **Vercel-hosted** **PostgreSQL** database, and **Gorm** is employed for managing database services. The project also includes a **Makefile** for running common commands, **VS Code extensions** for enhancing the development experience, and a **CI/CD workflow** for automating the development process. The **Postman API documentation** is also included.


> [!NOTE]
> As example application for the REST APIs, the APIs are used to managed startup and tech companies. It is somehow more original than classic TODO APIs

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

<!-- 
### Tools and Technologies
- **Go**
Go language `v1.19`. Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.
- **Gin**
Gin is a web framework written in Go (Golang). It features a martini-like API with much better performance, up to 40 times faster thanks to httprouter. If you need performance and good productivity, you will love Gin.
- **Gorm**
The fantastic ORM library for Golang, aims to be developer friendly.
- **PostgreSQL**
PostgreSQL is a powerful, open source object-relational database system with over 30 years of active development that has earned it a strong reputation for reliability, feature robustness, and performance.
- **Docker**
Docker is a set of platform as a service (PaaS) products that use OS-level virtualization to deliver software in packages called containers.
- **Kubernetes**
Kubernetes is an open-source container-orchestration system for automating computer application deployment, scaling, and management.
- **AWS ECS**
Amazon Elastic Container Service (Amazon ECS) is a fully managed container orchestration service. Customers such as Duolingo, Samsung, GE, and Cookpad use ECS to run their most sensitive and mission critical applications because of its security, reliability, and scalability.
- **Vercel**
Vercel is a cloud platform for static sites and Serverless Functions that fits perfectly with your workflow. It enables developers to host Jamstack websites and web services that deploy instantly, scale automatically, and requires no supervision, all with no configuration.
- **Postman**
Postman is a collaboration platform for API development. Postman's features simplify each step of building an API and streamline collaboration so you can create better APIs—faster.
- **VS Code**
Visual Studio Code is a free source-code editor made by Microsoft for Windows, Linux and macOS. Features include support for debugging, syntax highlighting, intelligent code completion, snippets, code refactoring, and embedded Git. -->

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



### API Documentation
The API documentation is done using Postman.

### Contributing
Contributions are welcome. Feel free to open a pull request or branch from this project.

### License
This project is licensed under the MIT LICENSE



# Lawn Mower

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/1e409c5ff2a340259be2b10e5da88a66)](https://app.codacy.com/gh/jrobic/lawn-mower?utm_source=github.com&utm_medium=referral&utm_content=jrobic/lawn-mower&utm_campaign=Badge_Grade_Settings)

Monorepo Go, React, Typescript. Learn Golang Microservices and other stuff
🌀 This repository is a work in progress and will be completed over time 🚀

# Table of Contents

- [Lawn Mower](#lawn-mower)
- [Table of Contents](#table-of-contents)
  - [The Goals of This Project](#the-goals-of-this-project)
  - [Plan](#plan)
  - [Folder Structure](#folder-structure)
  - [Technologies - Libraries](#technologies---libraries)
  - [The Domain And Bounded Context - Service Boundary](#the-domain-and-bounded-context---service-boundary)

## The Goals of This Project

- Learning Golang and new tools like Kubernetes, ArgoCD, Prometheus and others.
- Microservices based on `Domain Driven Design (DDD)` implementation.
- Communication internally between our microservices with `gRPC` synchronously.
- Implementing various types of testing like `Unit Testing`, `Integration Testing` and `End To End Testing`.
- Using `Health Check` for reporting the health of app infrastructure components.
- Using Docker-Compose and Kubernetes for our deployment mechanism.

## Plan

> This project is a work in progress, new features will be added over time.

High-level plan is represented in the Table

| Feature                   | Status      |
| ------------------------- | ----------- |
| Api Gateway               | To Do       |
| Catalog Service           | In Progress |
| Booking Service           | To Do       |
| User Service              | To Do       |
| Catalog Management Client | To Do       |
| Booking Management Client | To Do       |
| User Management Client    | To Do       |
| Lawn Mower Client         | To Do       |
| Lawn Mower App            | To Do       |

## Folder Structure

```
├── apps
│   └── catalog-service
│   └── booking-service
│   └── user-service
│   └── api-gateway
│   └── catalog-management-client
│   └── booking-management-client
│   └── user-management-client
│   └── lawn-mower-client
│   └── lawn-mower-app
├── docs
└── packages
    └── lawn-mower-ui
```

## Technologies - Libraries

> This project is a work in progress, new technologies or libraries will be added over time.
>
- **[`Go`](https://go.dev/)** - Build fast, reliable, and efficient software at scale
- **[`gRPC`](https://grpc.io/)** - A high performance, open source universal RPC framework
<!-- - **[``]()** -  -->

## The Domain And Bounded Context - Service Boundary

- `Catalog Service`: The Catalog Service is a service to handle catalog|product related operations

- `Booking Service`: The Booking Service manage all operation related to booking product.

- `User Service`: The User Service manage user information, tracking activities.

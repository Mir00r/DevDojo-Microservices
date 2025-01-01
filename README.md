# DevDojo

**DevDojo** is a comprehensive learning and experimentation hub featuring a wide range of backend microservices built using various programming languages, frameworks, and technologies. This repository aims to provide developers—from beginners to seasoned professionals—a platform to explore, understand, and implement modern backend architectures and best practices.

## What is DevDojo?

DevDojo is a polyglot environment where you can find multiple microservices, each implemented in a different language and framework. Whether you’re interested in Java, Kotlin, Go, C#, Python, or simply curious about microservices architecture, DevDojo serves as an open playground. Here, you can study the code, test features, and experiment with configurations, all while learning the nuances of different tech stacks.

## Key Features

- **Polyglot Services:**  
  Each directory hosts a microservice built in a different language and technology stack. Examples include:
    - **Java (Spring Boot)**: RESTful API demonstrating typical MVC patterns and DI.
    - **Kotlin (Spring Boot)**: A service showcasing Kotlin’s concise syntax and Spring Boot’s rapid development features.
    - **Go (Gin)**: A high-performance service built with Gin, illustrating Go’s simplicity, concurrency, and speed.
    - **C# (.NET)**: A microservice running on .NET Core, showing C#’s powerful features and compatibility across platforms.
    - **Python (Django)**: A service leveraging Django’s batteries-included philosophy for rapid API development.

- **Microservices Architecture Best Practices:**  
  Learn how microservices communicate, discover service boundaries, and implement APIs that are independently deployable, scalable, and maintainable.

- **Infrastructure & DevOps Integration:**  
  See examples of containerization (Docker), orchestration (Kubernetes), and CI/CD pipelines. Understand how different services fit into the DevOps lifecycle.

- **Security & Observability:**  
  Explore common security patterns (JWT, OAuth2) and tools for monitoring and logging (Prometheus, ELK stack, Grafana), ensuring each service is not only functional but also secure and observable in production environments.

## Root-Level Directory Structure
```
project-root/
│
├── services/                 # Each microservice lives in its own folder
│   ├── auth-service/         # Authentication microservice
│   ├── user-service/         # User management microservice
│   ├── role-service/         # Role and permission management microservice
│   ├── audit-service/        # Audit and logging microservice
│   └── ...                   # Future microservices for other domains
│
├── shared-libraries/         # Shared code and utilities
│   ├── common/               # Common utilities (e.g., logging, error handling)
│   ├── models/               # Shared models and DTOs
│   ├── config/               # Shared configuration utilities
│   └── security/             # Security-related shared code (e.g., JWT utilities)
│
├── infra/                    # Infrastructure configuration
│   ├── k8s/                  # Kubernetes manifests
│   ├── docker/               # Dockerfiles and Docker Compose
│   ├── ci-cd/                # CI/CD pipeline configurations
│   ├── monitoring/           # Monitoring and alerting (e.g., Prometheus, Grafana)
│   └── secrets/              # Encrypted secrets and environment variables
│
├── api-gateway/              # Central API Gateway configuration and logic
│   ├── kong/                 # Kong API Gateway setup
│   ├── nginx/                # NGINX configuration (optional)
│   └── ...                   # Other gateway tools
│
├── documentation/            # API and system documentation
│   ├── openapi/              # OpenAPI/Swagger specifications
│   └── architecture-diagrams/ # Diagrams for architecture, data flow, etc.
│
├── scripts/                  # Deployment, testing, and utility scripts
│   ├── db-migrations/        # Database migration scripts (e.g., Liquibase, Flyway)
│   ├── load-testing/         # Load testing scripts (e.g., k6, JMeter)
│   └── local-setup/          # Scripts to run services locally
│
├── logs/                     # Log files (git-ignored)
│
├── .gitignore                # Git ignore file
├── README.md                 # Project documentation
└── LICENSE                   # License information
```

## Detailed Folder Structure for Each Microservice
```
<service name>-service/
│
├── cmd/                      # Main entry points for the service
│   └── main.go               # Main file to start the service
├── config/               # Service-specific configuration
│       ├── app_config.go
│       └── env/              # Environment variables handling
│── constants/               
│       ├── app_constant.go
|── containers/               
│       ├── container.go
|── errors/               
│       ├── errors.go
|── middlewares/               
│       ├── error_middleware.go
|── routes/               
│       ├── routes.go
├── internal/                 # Internal application code (not accessible externally)
│   ├── api/                  # API layer
│   │   ├── controllers/      # REST API controllers (e.g., login, MFA)
│   │
│   ├── services/             # Business logic (e.g., token generation, authentication)
│   │   ├── auth.go
│   │   └── user.go
│   │
│   ├── models/               # Data models (e.g., User, Token)
│   |    ├── dtos/
│   │       └── sample_dto.go
│   |    ├── entities/
│   │       └── sample_entities.go
│   │
│   ├── repositories/         # Database access logic (e.g., PostgreSQL queries)
│   │   └── user_repository.go
│   │
│   ├── utils/                # Utility functions (e.g., hashing, JWT handling)
│   │   ├── bcrypt.go
│   │   └── jwt.go
│   │
│
├── test/                     # Unit and integration tests
│   ├── auth_test.go
│   └── mfa_test.go
├── scripts/                     
│   ├── migration_gen.go
│
├── db/                       # Database migrations and seeds
│   ├── migrations/
│   │   ├── 001_create_users_table.sql
│   │   └── 002_create_tokens_table.sql
│   └── seeds/
│       └── sample_data.sql
|   └── migration.go
|   └── init.go
│
├── build/                    # Build configurations
│   ├── Dockerfile            # Dockerfile for the microservice
│   ├── Makefile              # Makefile for building and testing
│   └── helm/                 # Helm charts for Kubernetes deployment
│
├── docs/                     # Service-specific documentation
│   └── openapi.yaml          # OpenAPI/Swagger spec for the service
│
└── logs/                     # Service-specific log files (git-ignored)
```

## Shared Libraries (shared-libraries/)
```
shared-libraries/
│
├── common/                   # Common utilities
│   ├── logger.go             # Centralized logging
│   ├── error.go              # Custom error handling
│   └── pagination.go         # Pagination utilities
│
├── models/                   # Shared data models
│   └── user.go               # User model shared across services
│
├── config/                   # Shared configuration code
│   └── app_config.go
│
├── security/                 # Shared security utilities
│   ├── jwt_utils.go          # JWT generation and validation
│   ├── hashing.go            # Password hashing utilities
│   └── cors.go               # Cross-Origin Resource Sharing setup
└── ...
```

## Infrastructure (infra/)
```
infra/
│
├── k8s/                      # Kubernetes manifests
│   ├── auth-service.yaml     # Deployment for the auth service
│   ├── user-service.yaml     # Deployment for the user service
│   ├── ingress.yaml          # Ingress configurations
│   └── secrets.yaml          # Encrypted secrets
│
├── docker/                   # Docker setup for local testing
│   ├── docker-compose.yaml   # Docker Compose for local development
│   └── Dockerfile            # Base Dockerfile for services
│
├── ci-cd/                    # CI/CD pipeline configurations
│   ├── github-actions/       # GitHub Actions YAML files
│   └── jenkins/              # Jenkins pipelines
│
├── monitoring/               # Monitoring setup
│   ├── prometheus/           # Prometheus configuration
│   └── grafana/              # Grafana dashboards
│
└── secrets/                  # Encrypted secrets for deployment
```

## Repository Structure

```
DevDojo/
├─ java-springboot-service/
│  ├─ src/
│  ├─ pom.xml
│  └─ README.md
├─ kotlin-springboot-service/
│  ├─ src/
│  ├─ build.gradle.kts
│  └─ README.md
├─ go-gin-service/
│  ├─ cmd/
│  ├─ internal/
│  └─ README.md
├─ csharp-dotnet-service/
│  ├─ src/
│  ├─ test/
│  └─ README.md
├─ python-django-service/
│  ├─ project/
│  └─ README.md
└─ docs/
   └─ architecture-diagrams/
```

- **Language-Specific Directories:** Each service directory contains all the source code, configuration files, and a dedicated README to explain setup, running, and testing instructions for that particular service.
- **Docs & Architecture:** The `docs/` folder may include architectural diagrams, notes, and guides on how to deploy or integrate these services with other systems.

## Getting Started

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/YourUsername/DevDojo.git
   cd DevDojo
   ```

2. **Prerequisites:**
    - **General:** Docker, Docker Compose (if you’re planning to run services in containers)
    - **Language-Specific:**
        - Java & Kotlin: JDK 11+
        - Go: Go 1.18+
        - C#: .NET 6+
        - Python: Python 3.9+ and `pipenv` or `virtualenv`

3. **Running a Service:**
   Each service directory includes its own README with step-by-step instructions. For example:
    - **Java (Spring Boot)**:
      ```bash
      cd java-springboot-service
      ./mvnw spring-boot:run
      ```
    - **Go (Gin)**:
      ```bash
      cd go-gin-service
      go run cmd/main.go
      ```

   Follow similar instructions for other services.

4. **Testing & Validation:**
   Most services come with their own test suites. Run them according to the instructions in the respective service’s README.

## Learning Objectives

- **Compare & Contrast Technologies:**  
  Understand the strengths and trade-offs of different programming languages and frameworks.

- **Microservices Principles:**  
  Gain hands-on experience with concepts like domain-driven design, distributed tracing, circuit breakers, and service discovery.

- **Practical DevOps:**  
  Learn how to containerize your services, configure CI/CD pipelines, and deploy to different environments.

- **Security & Stability:**  
  Explore best practices in authentication, authorization, and monitoring to keep services reliable and secure.

## Contributing

We warmly welcome contributions! If you have a microservice, example code, or documentation to add, feel free to open a Pull Request. Please review the [CONTRIBUTING.md](CONTRIBUTING.md) file for guidelines on coding standards, branching strategies, and testing procedures.

## License

This project is licensed under the [MIT License](LICENSE). Feel free to use the code and modify it for your personal or professional learning needs.

## Contact & Support

For questions, suggestions, or feedback, please open an issue in the repository’s [Issue Tracker](../../issues). We aim to foster a supportive community where everyone can learn and grow together.

---

**Happy Coding & Learning!******

# Online Stock Brokerage System

This is a microservice-based online stock brokerage system built using Java 21, Spring Boot, PostgreSQL, Hibernate, Flyway, and Docker.

## Services

1. **API Gateway**: Routes requests to the appropriate microservices.
2. **Config Server**: Manages configuration for all microservices.
3. **Eureka Server**: Handles service discovery.
4. **User Service**: Manages user-related operations.
5. **Stock Service**: Manages stock-related operations.
6. **Trade Service**: Manages trading operations.
7. **Notification Service**: Manages notifications.
8. **Payment Service**: Manages deposits and withdrawals.

## Running the System

1. Clone the repository.
2. Run `docker-compose up` to start all services.
3. Access the API Gateway at `http://localhost:8080`.

## Database Migrations

Database migrations are managed by Flyway. Each service has its own schema and migrations.

## Design Patterns

- **Microservice Architecture**: Each service is independently deployable and scalable.
- **API Gateway**: Centralized request routing.
- **Service Discovery**: Eureka Server for service registration and discovery.
- **Circuit Breaker**: Implemented using Hystrix for fault tolerance.
- **Repository Pattern**: Used for database access in each service.

## Clean Coding Principles

- **SOLID Principles**: Applied throughout the codebase.
- **DRY (Don't Repeat Yourself)**: Reusable components and utilities.
- **KISS (Keep It Simple, Stupid)**: Simple and straightforward code.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)

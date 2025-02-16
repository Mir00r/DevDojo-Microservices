# API Gateway Service

## Service Overview
The API Gateway service acts as the central entry point for all client requests in our microservices architecture. It handles routing, authentication, rate limiting, and load balancing for the entire system.

### Key Features
- **Centralized Routing**: Routes requests to appropriate microservices
- **Authentication & Authorization**: JWT validation and role-based access control
- **Rate Limiting**: Prevents API abuse through sophisticated rate limiting
- **Load Balancing**: Distributes traffic across service instances
- **Monitoring**: Comprehensive monitoring and logging
- **Protocol Support**: Handles HTTP/HTTPS, WebSocket, and gRPC
- **Service Discovery**: Dynamic service discovery and registration

## Technology Stack

### Core Technologies
- **NGINX**: High-performance reverse proxy and load balancer
    - Used for: Request routing, load balancing, SSL termination
    - Version: 1.24.0 (Latest stable)

- **Go**: Primary implementation language for custom logic
    - Used for: Custom middleware, service discovery, dynamic routing
    - Version: 1.21.x
    - Benefits: High performance, low latency, excellent concurrency

### Infrastructure
- **Docker**: Containerization
    - Used for: Service containerization and deployment
    - Version: 24.x

- **Kubernetes**: Container orchestration
    - Used for: Service orchestration, scaling, and management
    - Version: 1.28.x

### Monitoring & Logging
- **Prometheus**: Metrics collection
    - Used for: Performance metrics, custom metrics
    - Version: 2.45.x

- **Grafana**: Metrics visualization
    - Used for: Dashboards, alerts
    - Version: 10.x

- **ELK Stack**: Logging solution
    - Elasticsearch: 8.x
    - Logstash: 8.x
    - Kibana: 8.x

### Security
- **OAuth2/JWT**: Authentication and authorization
- **mTLS**: Service-to-service communication
- **Rate Limiting**: Request throttling

## Project Structure
```
api-gateway/
├── config/
│   ├── nginx/                 # NGINX configurations
│   │   ├── nginx.conf        # Main NGINX config
│   │   └── conf.d/           # Configuration fragments
│   │       ├── routes.conf   # Service routing rules
│   │       ├── proxy.conf    # Proxy settings
│   │       ├── security.conf # Security settings
│   │       └── rate-limiting.conf
│   ├── ssl/                  # SSL certificates
│   └── environment/          # Environment configurations
│       ├── development.env
│       ├── staging.env
│       └── production.env
├── src/                      # Go source code
│   ├── middlewares/         # Custom middleware implementations
│   │   ├── auth/           # Authentication middleware
│   │   ├── ratelimit/      # Rate limiting logic
│   │   └── logging/        # Logging middleware
│   ├── services/           # Core services
│   │   ├── discovery/      # Service discovery
│   │   └── health/         # Health checking
│   └── utils/              # Utility functions
├── monitoring/              # Monitoring configurations
│   ├── prometheus/         # Prometheus configs
│   └── grafana/            # Grafana dashboards
├── docker/                 # Docker configurations
│   ├── Dockerfile         # Main service Dockerfile
│   └── docker-compose.yml # Local development setup
├── kubernetes/            # Kubernetes manifests
│   ├── deployment.yaml
│   ├── service.yaml
│   └── ingress.yaml
├── scripts/               # Utility scripts
│   ├── setup.sh
│   └── deploy.sh
└── tests/                # Test suites
    ├── integration/
    └── unit/

## Setup Instructions

### Prerequisites
- Docker 24.x
- Docker Compose 2.x
- Go 1.21.x
- NGINX 1.24.x
- kubectl (for Kubernetes deployment)

### Local Development Setup

1. Clone the repository:
```bash
git clone https://github.com/your-org/api-gateway.git
cd api-gateway
```

2. Set up environment variables:
```bash
cp config/environment/development.env .env
```

3. Install dependencies:
```bash
go mod download
```

4. Start the development environment:
```bash
docker-compose up -d
```

5. Verify the setup:
```bash
curl http://localhost:8080/health
```

### Production Deployment

1. Build the Docker image:
```bash
docker build -t api-gateway:latest .
```

2. Deploy to Kubernetes:
```bash
kubectl apply -f kubernetes/
```

## Configuration

### NGINX Configuration
The NGINX configuration is split into multiple files for better maintainability:

- `nginx.conf`: Main configuration file
- `routes.conf`: Service routing rules
- `security.conf`: Security settings
- `rate-limiting.conf`: Rate limiting rules

Example of updating rate limits:
```nginx
# In config/nginx/conf.d/rate-limiting.conf
limit_req_zone $binary_remote_addr zone=api_limit:10m rate=10r/s;
```

### Environment Variables
Key environment variables that need to be configured:

```plaintext
NODE_ENV=production
PORT=8080
USER_SERVICE_URL=http://user-service:3000
AUTH_SERVICE_URL=http://auth-service:4000
RATE_LIMIT_WINDOW_MS=300000
RATE_LIMIT_MAX_REQUESTS=30
```

## Monitoring & Logging

### Metrics
- Prometheus metrics available at `/metrics`
- Default Grafana dashboards in `monitoring/grafana/dashboards/`

### Logging
- Structured JSON logging enabled by default
- Logs shipped to ELK stack in production
- Log levels configurable via environment variables

## API Documentation

### Service Routes
| Route | Service | Description |
|-------|----------|-------------|
| `/api/users/*` | User Service | User management endpoints |
| `/api/auth/*` | Auth Service | Authentication endpoints |
| `/api/reviews/*` | Review Service | Review management endpoints |

### Health Check
- Endpoint: `/health`
- Method: GET
- Response: `{"status": "UP"}`

## Contributing
1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

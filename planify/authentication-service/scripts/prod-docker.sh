# Production
docker build --target production -t auth-service:prod .
docker run -p 3000:3000 auth-service:prod

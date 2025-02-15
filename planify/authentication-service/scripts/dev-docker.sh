# Development
docker build --target development -t auth-service:dev .
docker run -p 3000:3000 -v $(pwd):/usr/src/app auth-service:dev

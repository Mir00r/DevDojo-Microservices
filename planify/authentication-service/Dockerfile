# Stage 1: Development
FROM node:18-alpine AS development

# Set working directory
WORKDIR /usr/src/app

# Copy package files
COPY package*.json ./

# Install dependencies including 'devDependencies'
RUN npm install

# Copy source code
COPY . .

# Expose port
EXPOSE 3000

# Command to run in development
CMD ["npm", "run", "dev"]

# Stage 2: Production Build
FROM node:18-alpine AS build

WORKDIR /usr/src/app

# Copy package files
COPY package*.json ./

# Install dependencies excluding 'devDependencies'
RUN npm ci --only=production

# Copy source code
COPY . .

# Stage 3: Production
FROM node:18-alpine AS production

# Add node user for security
RUN addgroup -g 1001 -S nodejs
RUN adduser -S nodeapp -u 1001

# Set working directory
WORKDIR /usr/src/app

# Copy package files
COPY package*.json ./

# Copy built node modules and source code
COPY --from=build /usr/src/app/node_modules ./node_modules
COPY --from=build /usr/src/app/src ./src
COPY --from=build /usr/src/app/config ./config

# Copy migration files if needed
COPY --from=build /usr/src/app/migrations ./migrations
COPY --from=build /usr/src/app/models ./models

# Set environment variables
ENV NODE_ENV=production
ENV PORT=3000

# Change ownership to non-root user
RUN chown -R nodeapp:nodejs /usr/src/app

# Switch to non-root user
USER nodeapp

# Expose port
EXPOSE 3000

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=30s \
  CMD node -e "try { require('http').get('http://localhost:3000/health', (r) => r.statusCode === 200 ? process.exit(0) : process.exit(1)); } catch (e) { process.exit(1); }"

# Command to run in production
CMD ["node", "src/server.js"]

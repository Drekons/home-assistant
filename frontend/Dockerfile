# Build stage
FROM node:20-alpine as build-stage

WORKDIR /app

# Copy package.json and yarn.lock (if exists)
COPY package.json yarn.lock* ./

# Install dependencies
RUN yarn install --frozen-lockfile

# Copy project files
COPY . .

# Build the project
RUN yarn build

# Production stage
FROM nginx:stable-alpine as production-stage

# Copy built assets from build-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html

# Copy nginx configuration if you have a custom one
# COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
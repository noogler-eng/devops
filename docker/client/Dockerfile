# syntax=docker/dockerfile:1.5

# Uses Node.js 19.4 (Debian Bullseye variant) as the 
# base image for building the React app.
# Assigns this stage the alias build, 
# so we can reference it later
FROM node:19.4-bullseye AS build

# Specify working directory other than /
# Sets the working directory inside the container to /usr/src/app
WORKDIR /usr/src/app

# Copy only files required to install
# dependencies (better layer caching)
COPY package*.json ./

# Use cache mount to speed up install of existing dependencies
RUN --mount=type=cache,target=/usr/src/app/.npm \
  npm set cache /usr/src/app/.npm && \
  npm install

COPY . .

# Runs the React production build (creates an optimized dist folder).
RUN npm run build

# Use separate stage for deployable image
# Uses Nginx (unprivileged) as the base image for the final deployable app.
# Lightweight Alpine-based image with an added Perl module
FROM nginxinc/nginx-unprivileged:1.23-alpine-perl

# Use COPY --link to avoid breaking cache if we change the second stage base image
# Copies a custom Nginx configuration file (likely sets up routing for React's SPA)
COPY --link nginx.conf /etc/nginx/conf.d/default.conf

# Copies the built React app (dist/ folder) from the build 
# stage into Nginx’s default serving directory.
COPY --link --from=build usr/src/app/dist/ /usr/share/nginx/html

# Opens port 8080, which is where Nginx serves the app.
EXPOSE 8080
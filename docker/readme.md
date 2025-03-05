README.md

# Docker Setup Guide

This guide provides instructions for building and running a Dockerized application using `docker` and `docker-compose`.

## 🚀 Building and Running Containers

### 1️⃣ Build Docker Image
To build a Docker image, use the following command:
```sh
docker build -t image_name .
```
Replace image_name with your preferred name for the image.

2️⃣ Run Containers with Docker Compose
To start the services defined in docker-compose.yml:

```sh
docker compose up
```

🔹 Run in Detached Mode

To keep the containers running in the background:

```sh
docker compose up -d
```

3️⃣ Stop and Remove Containers
To stop running containers:

```sh
docker compose down
```
To remove containers, networks, and volumes:

```sh
docker compose down -v
```

4️⃣ View Running Containers
To check all active containers:

```sh
docker ps
```
To view all containers (including stopped ones):
```sh
docker ps -a
```
5️⃣ View Container Logs
To see logs of a specific container:
```sh
docker logs container_name
```
For real-time logs:
```sh
docker logs -f container_name
```

6️⃣ Enter a Running Container
To get an interactive shell inside a running container:
```sh
docker exec -it container_name sh
```
For containers running with Bash:
```sh
docker exec -it container_name bash
```
7️⃣ Remove Docker Images & Containers
To remove a container:
```sh
docker rm container_id
```
To remove an image:
```sh
docker rm image_id
```

To remove all stopped containers:
```sh
docker container prune
```

To remove all unused images:
```sh
docker image prune -a
```

8️⃣ List Docker Images
To list all available Docker images:
```sh
docker images
```
9️⃣ Check Running Docker Services
To list services running inside a Docker Compose project:
```sh
docker compose ps
```
🔥 Bonus: Connecting to PostgreSQL Database Inside Docker
If you have a PostgreSQL container running, you can connect to it with:
```sh
docker exec -it container_name psql -U postgres -d postgres
```
🛑 Troubleshooting

🔹 "Port Already in Use" Error
Find the process using the port:
```sh
lsof -i :PORT_NUMBER
```
Kill the process:
```sh
kill -9 PROCESS_ID
```
🔹 "No Space Left on Device"
Clean up unused Docker objects:
```sh
docker system prune -a
```
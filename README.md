# Getir Case Challenge

## Introduction
The application is a simple RESTful API with two endpoints. \
\
To test the api, you can use https://getir-case-challenge.herokuapp.com

### Allowed Endpoints and Methods

| Endpoint | Method |
| -------- | ------ |
| /records | POST |
| /in-memory | GET |
| /in-memory | POST |

## Configuration
The configuration is performed via environment variables. If you want to run the app properly, you should set these environment variables correctly.

| Variable | Description |
| -------- | ---------- |
| `DB_CONNECTION_STRING` | MongoDB connection string |
| `DB_NAME` | Default database name |
| `REDIS_URL` | In-memory database connection string |
| `PORT` | REST API port to serve |
| `APP_MODE` | TEST or PROD, if you use docker |

## Deployment

In order to deploy the app, you can use docker. You can use two different methods to deploy the app via docker.

### Docker

Example:
```bash
docker build -t skarakasoglu/g-case-challenge:1.0.0
docker run -p 8080:8080 --name GetirCaseChallenge \
-e DB_NAME=db -e DB_CONNECTION_STRING=connectionString \
-e REDIS_URL=redisConnectionString -e PORT=8080 \
-e APP_MODE=PROD
skarakasoglu/g-case-challenge
```

I generally like creating containers by using docker-compose because the configuration of the container becomes documented in a docker-compose.yml file.

### Docker-compose

Examples:
```bash
docker-compose up --build
```

or

```bash
docker-compose up
```

If you want to build the image using Dockerfile before the creating container, then you should use `-build` flag. If you already have the image you specify in the docker-compose.yml file and you only want to create new container with the environment variables, you can ignore `--build` flag.
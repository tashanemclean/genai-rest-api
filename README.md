# genai-rest-api

## Section 1: Local Development Environment Setup

This microservice is designed as to be the REST API gateway for the service
[genai-backend](https://github.com/tashanemclean/genai-backend)

### Setup Local Environment

To use this miroservice, you need to have the following tools installed on your
developer machine:

1. go [Golang](https://go.dev/doc/install)

### Configure Environment variables

To use environment variables, create configs directory in project root, place
dev.json inside it.

```
configs/dev.json
```

```
"PORT": "8080",
"ENVIRONMENT": "dev",
"API_BASE_URL": "http://localhost:9000"
```

### Running the app

To run , run:

```
$ go run main.go
```
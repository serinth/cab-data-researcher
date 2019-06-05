# Requirements
- Golang 1.12+ with Go Modules enabled, properly configured (Should also work with Go 1.11)
- Docker and Docker Compose
- Bash shell. If you use zsh or something else the `Makefile` may not work

## Tooling
The tooling can be installed with:
```bash
make tools
```

Or you can get the binaries and install them yourself.

- protoc-gen-grpc-gateway
- protoc-gen-go
- protoc-gen-swagger (optional)


# Quick Start - Go Application

Clone this repo. Ensure ports defined in the `docker-compose.yml` file are not taken up by other services including the docker daemon.

```bash
make compile-protobuf # only if needed to re-compile proto files, see Defining Services below
go run main.go

# health probe
curl localhost:8080/_ah/health
```

# Tests

```bash
go test -v ./...
```

# Scaffolding

Generated using my code generator: https://github.com/serinth/generator-go-grpcgw-api

Using [GRPC Gateway](https://github.com/grpc-ecosystem/grpc-gateway).

# Defining Services

Services are defined in `/proto`. When compiled, the generated interfaces need to be implemented. In this example, they're implemented in `/protoServices` but how it's structured is completely up to you.

In the protobuf 3 language, it's very simple to define standard HTTP methods. We define the request model and view model to be returned in the proto files.

The normal approach to the services is:
 1. Define the endpoints and models
 2. Run the protobuf compiler (Provided in the Makefile). This will generate boilerplate GRPC code with interfaces that your services must implement. There is also the added benefit of Swagger docs automatically being generated with a protoc plugin.
 3. Wire up the new services in `main.go`

There is a complete health endpoint example.

# Configurations

Configurations are loaded from the toml files in `/configs`. The behaviour is as follows:
 
 1. .toml file gets loaded first
 2. environment variable will override what's in the toml file
 3. required environment variables do not need to be added to the toml files

# Mandatory Environment Variables
| Variable | Example | Description |
| --- | --- | --- |
| ENVIRONMENT | local | name the config toml files to be the same as the environment variable name.

# Optional Environment Variable Overrides
| Variable | Example | Description |
| --- | --- | --- |
| ENABLE_DEBUGGING | true | Always enable log.debug
| API_PORT | ":8080" | The RESTful endpoint port
| GRPC_PORT | ":8081" | The GRPC endpoint port
| GRPC_HOST | "localhost" | The hostname of the GRPC server

# Swagger Docs
Swagger documentation can be generated from the proto files using
```bash
make generate-swagger
```

The `/proto` folder already has examples for this project which can be put into https://editor.swagger.io/ or run locally.
# CURL Commands
```bash
curl -X POST localhost:8080/cab/trips -d '{"Medallions":[{"Id": "D7D598CD99978BD012A87A76A7C891B7"},{"Id": "5455D5FF2BD94D10B304A15D4B7F2735"}], "SkipCache": true, "Date": "2013-12-01"}' -v
```
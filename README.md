# Flights_Backend

## Overview

The Flights API is built using Go and offers RESTful endpoints for flight information. It integrates with a remote flight data provider to retrieve flight details and serves these details through well-defined API endpoints.

The frontend is developed using React and provides a user-friendly interface for users to interact with the Flights API. Users can search for specific flights by ID and view the list of latest flights in a tabular format.

## 1- How to run

### Prerequisites

To run the API and the frontend, you need to have Docker and Docker Compose installed on the machine.
- Docker (Min version: 20.10.12): [Get Docker](https://docs.docker.com/get-docker/)
- Docker Compose (Min version: 2.27.1): [Get Docker Compose](https://docs.docker.com/compose/install/)

### Init submodules

In the main init the backend and frontend sumbodule:
```
git submodule init
git submodule update
```

### Build and run

In the main folder of the project, where the `docker-compose.yml` file is located, execute:

```
docker-compose up --build
```

This command builds the Dockerfiles and pulls any necessary images from DockerHub. Once docker-compose has finished building and running the images, the services will be up and running.

### Access the services
- **Backend API**: The backend API can be accessed at http://localhost:8000.
- **Frontend**: The frontend can be accessed at http://localhost:3000.
- **Swagger Documentation**: The Swagger documentation can be accessed at http://localhost:8000/swagger/index.html.

### Swagger Documentation

Swagger allows you to visualize and interact with the API in a user-friendly interface, providing detailed information about the various endpoints, including the requests and responses. You can explore the different endpoints, view the request and response formats, and interact with the API directly from the documentation interface.

To access the API documentation, open your browser and navigate to the following URL:

```
http://localhost:8000/swagger/index.html
```

![swagger1](https://github.com/carlos2380/webCarlos2380/blob/master/flights/FlightsSwagger.jpeg)

## 2- Documentation

## Flags

Flags are used to configure the application at startup.

```go
port := flag.String("port", "8000", "Port to server will be listening.")
flag.Parse()
```

## TESTS

### Table Tests

I use table-driven tests to validate our code with various inputs. This method keeps tests organized and maintainable.

- https://github.com/carlos2380/flights_backend/blob/main/internal/fetcher/radarbox/decoderRadarbos_test.go

### Mock Tests

Mocks are used to simulate dependencies, allowing isolated testing of individual components. This results in more focused and faster tests.

 - https://github.com/carlos2380/flights_backend/blob/main/internal/fetcher/fetcher_mock.go

## Linter

I have integrated a linter into our project to ensure code quality and consistency. The linter checks for potential issues and enforces coding standards.
- https://github.com/carlos2380/flights_backend/blob/main/.golangci.yml

#### Intall

````
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
`````

#### Run

```
golangci-lint run
```

## Swagger Integration

We have integrated Swagger into the API to automatically generate documentation. This allows for a user-friendly interface to visualize and interact with the API endpoints.

### Installation

To install Swagger, I use the following command:

```
go install github.com/swaggo/swag/cmd/swag@latest
````

### Documentation Generation

To generate the Swagger documentation, run the following command:

```
swag init -g cmd/server/main.go
```
This command scans the Go project files, extracts annotations from the code, and generates the swagger.json and swagger.yaml files in the docs folder.

## Continuous Integration (CI)

I set up a Continuous Integration (CI) pipeline to automatically test and validate the codebase with each push to the repository. This ensures that our code remains reliable and that new changes do not introduce any issues.

- https://github.com/carlos2380/flights_backend/blob/main/.github/workflows/ci.yml

## Error Handling

I implemented a  centralized error handling mechanism with a dedicated errors.go file. 
- https://github.com/carlos2380/flights_backend/blob/main/internal/errors/errors.go

This approach ensures consistent and clear error messages, making the application easier to maintain and debug, while adhering to best programming practices.

## Decoupling with Interfaces

In this project, we have decoupled the code by using interfaces. This design promotes modularity, making the codebase more flexible, maintainable, and easier to test and extend. By defining interfaces for key components, we can easily mock these components during testing, ensuring that each part of our application can be developed and tested in isolation.

For each endpoint in the handler, there is an associated interface. This approach makes it straightforward to change how flight data is fetched. For instance, while we currently use Radarbox to fetch flight data, it would be simple to switch to using a database or another method without altering the handler logic. For example, we can change the implementation of FlightsFetcher without affecting FlightInfoFetcher. This flexibility allows us to adopt different data sources for different endpoints, if needed.

## Server Shutdown

To prevent the server from shutting down while there are pending tasks, a graceful shutdown mechanism is implemented. The server captures shutdown signals and waits for pending tasks to complete before shutting down.
```
    sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
    <-sig
	log.Logger.Info("Shutting down server...")
	cancel()

	time.Sleep(2 * time.Second)
	log.Logger.Info("Server stopped")
```

## Areas for Improvement

### Enhanced Test Coverage

More Test Cases: Expand the current test suite to cover a wider range of scenarios, including edge cases and error conditions. This ensures that the application handles all possible inputs and states correctly.

### Automatic Documentation Generation

Integrate automatic generation of documentation into the CI pipeline. This ensures that the documentation is always up-to-date with the latest code changes.

### Monitoring and Metrics

Prometheus Integration: Integrate Prometheus for monitoring application performance and gathering metrics. This helps in identifying performance bottlenecks and monitoring the health of the application.

Grafana Dashboards: Set up Grafana dashboards to visualize the metrics collected by Prometheus, providing a clear and accessible overview of application performance and health.

Health Check: Implement an endpoint dedicated to verifying the status of key system components, such as database connection, service availability, etc.

### Security Enhancements

Implement a robust authentication and authorization mechanisms to secure the application.

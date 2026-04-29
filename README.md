# dnscheck

dnscheck is a lightweight API service for querying the IP, CName, NS, MX, and TXT records of a given URL. 
It leverages the Fiber framework for handling HTTP requests, standard Go goroutines and channels for concurrent DNS lookups (for experimentation purposes), and Redis for an efficient caching mechanism.

The structure is kept simple and straightforward following the principle of **KISS** (Keep It Stupid Simple).

## Directory Structure

```text
.
├── cache/          # Redis client initialization and caching methods
├── controller/     # HTTP handlers to process incoming requests
├── model/          # Data structures for DNS records
├── router/         # API route configurations
├── service/        # Core business logic for performing concurrent DNS lookups
├── Dockerfile      # Docker image configuration for the Go application
├── docker-compose.yml # Orchestration for the app and Redis containers
└── main.go         # Application entry point
```

## Build Process

The project is containerized to make the build and execution process simple. It uses a multi-stage Dockerfile to build the Go application and a `docker-compose.yml` to orchestrate the application along with its Redis dependency.

To build and run the project using Docker:

```sh
docker compose up --build
```

If you prefer to build and run it locally without Docker:

1. Ensure a Redis instance is running and accessible.
2. Export the Redis URL environment variable (default in compose is `redis:6379`):
   ```sh
   export REDIS_URL=localhost:6379
   ```
3. Build and run the Go application:
   ```sh
   go run main.go
   ```

## Usage

The application exposes a single REST endpoint running on port 3000.

**Endpoint:** `GET /dns/{url-to-query}`

### Examples

Use the cURL command below to interact with the endpoint:

```sh
curl "http://localhost:3000/dns/{url-to-query}"
```

eg: `curl "http://localhost:3000/dns/google.com"`

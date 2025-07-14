# Targeting Engine (Adserving Service)

A scalable, modular ad-serving service in Go (using go-kit patterns) with a MySQL backend and in-memory campaign targeting. Designed for high performance, extensibility, and production-readiness.

---

## Features
- **GET /v1/delivery**: Returns all active campaigns matching targeting rules (app, OS, country)
- **In-memory campaign cache**: Fast, efficient, and refreshed in the background
- **Case-insensitive matching**: All targeting is normalized
- **Swagger/OpenAPI**: API docs at `/swagger/`
- **Health endpoint**: `/health`
- **Configurable**: YAML-based config for DB, server, refresh interval
- **Production-ready**: Modular, idiomatic Go, Dockerfile, Makefile

---

## Architecture
- **cmd/server/**: Entrypoint
- **internal/service/**: Business logic, campaign refresh, matching
- **internal/endpoint/**: go-kit endpoint definitions
- **internal/transport/**: HTTP handlers (ad, health, swagger)
- **internal/repository/**: MySQL access
- **internal/model/**: Data models, errors
- **pkg/config/**: Config loader
- **res/mysql/**: DB schema
- **res/data/**: Dummy/test data, Postman/REST requests

---

## Setup

### 1. Prerequisites
- Go 1.20+
- MySQL 8+
- (Optional) Docker

### 2. Database
- Create schema using files in `res/mysql/`
- Load dummy data:
  ```sh
  mysql -u <user> -p <db> < res/data/dummy_data.sql
  ```

### 3. Configuration
- Copy `configs/config.yaml.example` to `configs/config.yaml` and edit as needed (DB DSN, port, etc)

### 4. Build & Run
- With Makefile:
  ```sh
  make build
  ./bin/server
  ```
- Or with Docker:
  ```sh
  docker build -t targeting-engine .
  docker run -p 8080:8080 --env-file .env targeting-engine
  ```

---

## API Usage

### Delivery Endpoint
- **GET /v1/delivery?app=...&country=...&os=...**
- Returns 200 with campaigns, 204 if no match, 400 for missing/invalid params

### Health
- **GET /health**

### Swagger Docs
- **GET /swagger/**

See `res/data/postman_requests.http` for ready-to-use test cases and expected responses.

---

## Testing
- Unit and integration tests: `go test ./...`
- API tests: Use Postman or REST Client with `res/data/postman_requests.http`

---

## Contribution
- Fork, branch, and PR as usual
- Keep code modular and idiomatic
- Add/maintain tests for new features

---

## License
MIT

## Observability: Prometheus, Grafana, Loki

This service supports full observability out of the box:

- **Prometheus**: Metrics exposed at `/metrics` (HTTP server)
- **Grafana**: Dashboards for metrics and logs
- **Loki**: Aggregated, structured JSON logs (via logrus)

### Running the Observability Stack

A `docker-compose.yml` is provided to run Prometheus, Grafana, Loki, and Promtail alongside the service:

```sh
docker-compose up
```

- Prometheus: [http://localhost:9090](http://localhost:9090)
- Grafana: [http://localhost:3000](http://localhost:3000) (default: admin/admin)
- Loki: [http://localhost:3100](http://localhost:3100)

### Metrics
- **/metrics** endpoint exposes Prometheus metrics:
  - `ad_requests_total{status="200|204|400|500"}`
  - `ad_request_duration_seconds`
  - `campaign_cache_refresh_total`, `campaign_cache_refresh_errors_total`, `campaign_cache_refresh_duration_seconds`

### Logs
- All logs are structured JSON (logrus), shipped to Loki via Promtail.
- View logs in Grafana (Explore tab, select Loki as data source).

### Dashboards
- Import the sample dashboard from `docs/grafana_dashboard.json` into Grafana for instant metrics and log visualization.

---

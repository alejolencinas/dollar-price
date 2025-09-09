# Dollar price

An API for getting the dollar price in Argentina. Internally, it scraps the BNA webpage to get these values. It keeps the values cached for 10 minutes.

## Getting started

### Requirements
- Go 1.23+
- Docker (optional, for containerized runs)

### Clone and setup
```bash
git clone https://github.com/alejolencinas/dollar-price.git
cd dollar-price
```

### Configuration
Environment variables (all optional with sensible defaults):
- `API_PORT`: Port the server listens on. Default: `8080`
- `APP_ENV`: Environment name. Default: `development`
- `BNA_URL`: Source page to scrape. Default: `https://www.bna.com.ar/Personas`

Example (macOS/Linux):
```bash
export API_PORT=8080
export APP_ENV=development
export BNA_URL=https://www.bna.com.ar/Personas
```

### Run locally
```bash
go run ./cmd/server
```

Server starts on `http://localhost:${API_PORT}` (default `8080`).

### Run with Docker
Build the image:
```bash
docker build -t dollar-price .
```

Run the container:
```bash
docker run --rm -p 8080:8080 \
  -e API_PORT=8080 \
  -e APP_ENV=development \
  -e BNA_URL=https://www.bna.com.ar/Personas \
  dollar-price
```

### API
- `GET /api/v1/ping` → health check
- `GET /api/v1/dollar` → returns latest buy/sell values and `fetched_at`

Example:
```bash
curl http://localhost:8080/api/v1/ping
curl http://localhost:8080/api/v1/dollar
```

### Project layout
- `cmd/server`: application entrypoint
- `internal/server`: HTTP server and routes
- `internal/api`: handlers
- `internal/scraper`: scraping and caching logic
- `internal/config`: environment-based configuration

### Development
- Use `go run ./cmd/server` for iterative development
- Environment variables are read at startup; restart after changes
# Task Queue System (WIP)

A distributed job processing system built in Go that allows clients to submit tasks for asynchronous execution by a pool of worker goroutines.

## Features

- **HTTP API** for job submission and status tracking
- **Worker Pool** with configurable concurrency
- **Database Persistence** with PostgreSQL/SQLite support
- **Job Status Tracking** (pending, processing, completed, failed)
- **Retry Logic** with exponential backoff
- **Authentication** via JWT tokens
- **Graceful Shutdown** and error recovery
- **Docker Support** for easy deployment

## Quick Start

### Prerequisites
- Go 1.21+
- PostgreSQL or SQLite
- Docker (optional)

### Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/task-queue-system
cd task-queue-system

# Install dependencies
go mod download

# Set up environment variables
cp .env.example .env
# Edit .env with your database settings

# Run database migrations
go run scripts/migrate.go

# Start the server
go run cmd/server/main.go

# In another terminal, start workers
go run cmd/worker/main.go
```

### Basic Usage

**Submit a job:**
```bash
curl -X POST http://localhost:8080/api/v1/jobs \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your-token>" \
  -d '{
    "type": "image_resize",
    "payload": {
      "image_url": "https://example.com/image.jpg",
      "width": 800,
      "height": 600
    }
  }'
```

**Check job status:**
```bash
curl http://localhost:8080/api/v1/jobs/{job_id} \
  -H "Authorization: Bearer <your-token>"
```

## Architecture

### Components
- **API Server**: Handles job submission and status queries
- **Worker Pool**: Processes jobs asynchronously  
- **Database**: Persists jobs and maintains state
- **Queue System**: Coordinates job distribution to workers

### Supported Job Types
- `image_resize` - Image processing and resizing
- `email_send` - Email delivery
- `data_export` - Database export operations
- `file_process` - General file processing

## Configuration

Environment variables:
```bash
DATABASE_URL=postgres://user:pass@localhost/taskqueue
SERVER_PORT=8080
WORKER_COUNT=5
JWT_SECRET=your-secret-key
LOG_LEVEL=info
```

## Development

### Running Tests
```bash
# Unit tests
go test ./internal/...

# Integration tests  
go test ./tests/integration/...

# All tests with coverage
make test-coverage
```

### Database Operations
```bash
# Run migrations
go run scripts/migrate.go

# Seed test data
go run scripts/seed.go

# Reset database
make db-reset
```

### Docker Development
```bash
# Start all services
docker-compose up

# Run just the database
docker-compose up postgres

# Build production images
make docker-build
```

## API Documentation

### Authentication
All API endpoints require a valid JWT token in the Authorization header:
```
Authorization: Bearer <jwt-token>
```

### Endpoints

#### Submit Job
```
POST /api/v1/jobs
Content-Type: application/json

{
  "type": "image_resize",
  "payload": {...},
  "priority": "normal"
}
```

#### Get Job Status
```
GET /api/v1/jobs/{job_id}

Response:
{
  "job_id": "123",
  "status": "completed",
  "result": {...},
  "created_at": "2025-09-01T10:30:00Z",
  "completed_at": "2025-09-01T10:31:30Z"
}
```

#### List Jobs
```
GET /api/v1/jobs?status=pending&limit=50
```

#### Health Check
```
GET /health

Response:
{
  "status": "ok",
  "workers": 5,
  "queue_depth": 23
}
```

## Job Status Flow

```
pending → processing → completed
   ↓           ↓
failed ← ─ ─ ─ ┘
   ↓
retrying → processing
```

## Deployment

### Production Setup
1. Set up PostgreSQL database
2. Configure environment variables
3. Run database migrations
4. Deploy server and worker containers
5. Set up monitoring and logging

### Scaling
- **Horizontal**: Add more worker instances
- **Vertical**: Increase worker count per instance
- **Database**: Use connection pooling and read replicas

## Monitoring

The system provides metrics at:
- `/metrics` - Prometheus metrics
- `/health` - Basic health status
- Database job statistics

## Contributing

1. Fork the repository
2. Create a feature branch
3. Add tests for new functionality
4. Ensure all tests pass
5. Submit a pull request

## License

MIT License - see LICENSE file for details

## Roadmap

- [ ] Job scheduling (cron-like functionality)
- [ ] Job dependencies and workflows  
- [ ] Web UI for job monitoring
- [ ] Job result webhooks
- [ ] Advanced retry policies
- [ ] Multi-tenant support

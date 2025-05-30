# Flight Order Service

A Go service that performs topological sorting of flight journeys to determine the optimal flight path.

## Features

- **Topological Sort Algorithm**: Implements Kahn's algorithm to:
  - Determine flight order from given journey pairs
  - Detect cycles in flight paths
  - Handle invalid journey inputs

- **REST API**:
  - POST `/journey/order` - Accepts flight journey pairs and returns ordered path
  - Error handling for invalid requests

- **Flight Service**:
  - Processes flight journey requests
  - Integrates with topological sort implementation
  - Handles service dependencies

## Technologies Used

- **Go** (1.21+)
- **Echo** (Web framework)
- **Cobra** (CLI framework)
- **Slog** (Structured logging)

## Installation & Running

1. Clone the repository:
```bash
git clone https://github.com/kp/flight-order.git
cd flight-order
```

2. Build and run:
```bash
go build
./flight-order apis
```

3. The service will run on `http://localhost:8000`

## API Usage

### Request
```bash
POST /flight/api/v0/journey/order
Content-Type: application/json

{
    "journies": [
        ["LAX", "DXB"],
        ["JFK", "LAX"], 
        ["SFO", "SJC"],
        ["DXB", "SFO"]
    ]
}
```

### Successful Response
```json
{
    "journey_order": ["JFK", "LAX", "DXB", "SFO", "SJC"]
}
```

### Error Responses
- `400 Bad Request` - Invalid input format
- `500 Internal Server Error` - Cycle detected in flight paths

## Project Structure
- `cmd/` - CLI commands
- `flight/` - Core flight service and topological sort
  - `sort.go` - Topological sort implementation
  - `types.go` - Data structures and interfaces
  - `controller.go` - API handlers
- `server/` - HTTP server setup

## Development

Run tests:
```bash
go test ./...
```


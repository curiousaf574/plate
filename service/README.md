# Plate Service

Backend API service for the Plate Kubernetes deployment platform.

## Setup

1. Install dependencies:
```bash
go mod download
```

2. Setup PostgreSQL database:
```bash
createdb plate
```

3. Configure the service:
```bash
cp config.yaml.example config.yaml
# Edit config.yaml with your settings
```

4. Run the service:
```bash
go run main.go serve
```

## Configuration

The service can be configured via:
- `config.yaml` file
- Environment variables
- Command line flags

### Required Components

- **PostgreSQL**: Database for storing projects, deployments, and logs
- **Kubernetes cluster**: Target deployment environment
- **ArgoCD**: GitOps deployment orchestration
- **Gitea**: Git repository hosting
- **Helm**: Package manager for Kubernetes

## API Endpoints

### Projects
- `GET /api/v1/projects` - List all projects
- `POST /api/v1/projects` - Create a new project
- `GET /api/v1/projects/:id` - Get project details
- `PUT /api/v1/projects/:id` - Update project
- `DELETE /api/v1/projects/:id` - Delete project

### Deployments
- `GET /api/v1/deployments` - List all deployments
- `POST /api/v1/deployments` - Create deployment
- `GET /api/v1/deployments/:id` - Get deployment details
- `DELETE /api/v1/deployments/:id` - Delete deployment
- `GET /api/v1/deployments/:id/logs` - Get deployment logs

### Environments
- `GET /api/v1/environments` - List environments
- `POST /api/v1/environments` - Create environment
- `GET /api/v1/environments/:id` - Get environment
- `PUT /api/v1/environments/:id` - Update environment

### Operations
- `POST /api/v1/deploy` - Trigger deployment
- `GET /api/v1/status` - Get deployment status
- `GET /health` - Health check

## Architecture

The service integrates with:

1. **Kubernetes API** - For cluster management and resource deployment
2. **ArgoCD API** - For GitOps application management
3. **Helm** - For package templating and deployment
4. **Gitea API** - For repository management and webhooks
5. **PostgreSQL** - For persistent data storage

## Development

```bash
# Run tests
go test ./...

# Build binary
go build -o plate-service .

# Run with custom config
./plate-service serve --config /path/to/config.yaml
```
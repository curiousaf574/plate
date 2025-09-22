# Plate API Documentation

## Overview

The Plate Service API provides RESTful endpoints for managing projects, deployments, and environments. All endpoints return JSON responses and use standard HTTP status codes.

## Base URL

```
http://localhost:8080/api/v1
```

## Authentication

Currently, the API runs in development mode without authentication. In production, it will use JWT tokens.

```bash
# Future production usage
curl -H "Authorization: Bearer <token>" <endpoint>
```

## Error Responses

All error responses follow this format:

```json
{
  "error": "Error message description"
}
```

## Endpoints

### Health Check

#### GET /health

Check if the service is running.

**Response:**
```json
{
  "status": "healthy"
}
```

---

## Projects

### List Projects

#### GET /api/v1/projects

Retrieve all projects.

**Response:**
```json
[
  {
    "id": 1,
    "name": "web-app",
    "description": "Frontend React application for the main website",
    "repository": "https://git.example.com/plate/web-app.git",
    "runtime": "nodejs",
    "build_cmd": "npm install && npm run build",
    "start_cmd": "npm start",
    "port": 3000,
    "env_vars": "{\"NODE_ENV\":\"production\",\"API_URL\":\"https://api.example.com\"}",
    "created_at": "2025-09-19T10:00:00Z",
    "updated_at": "2025-09-19T10:00:00Z"
  }
]
```

### Create Project

#### POST /api/v1/projects

Create a new project.

**Request Body:**
```json
{
  "name": "my-app",
  "description": "My application description",
  "repository": "https://git.example.com/org/my-app.git",
  "runtime": "nodejs",
  "build_cmd": "npm install",
  "start_cmd": "npm start",
  "port": 3000,
  "env_vars": "{\"NODE_ENV\":\"production\"}"
}
```

**Response:**
```json
{
  "id": 5,
  "name": "my-app",
  "description": "My application description",
  "repository": "https://git.example.com/org/my-app.git",
  "runtime": "nodejs",
  "build_cmd": "npm install",
  "start_cmd": "npm start",
  "port": 3000,
  "env_vars": "{\"NODE_ENV\":\"production\"}",
  "created_at": "2025-09-19T12:00:00Z",
  "updated_at": "2025-09-19T12:00:00Z"
}
```

### Get Project

#### GET /api/v1/projects/{id}

Retrieve a specific project by ID.

**Parameters:**
- `id` (path): Project ID

**Response:**
```json
{
  "id": 1,
  "name": "web-app",
  "description": "Frontend React application",
  "repository": "https://git.example.com/plate/web-app.git",
  "runtime": "nodejs",
  "build_cmd": "npm install && npm run build",
  "start_cmd": "npm start",
  "port": 3000,
  "env_vars": "{\"NODE_ENV\":\"production\"}",
  "created_at": "2025-09-19T10:00:00Z",
  "updated_at": "2025-09-19T10:00:00Z",
  "deployments": [
    {
      "id": 1,
      "environment_id": 3,
      "version": "v1.2.3",
      "status": "running",
      "url": "https://web-app.plate.local"
    }
  ]
}
```

### Update Project

#### PUT /api/v1/projects/{id}

Update a project.

**Parameters:**
- `id` (path): Project ID

**Request Body:**
```json
{
  "name": "web-app-updated",
  "description": "Updated description",
  "port": 3001
}
```

### Delete Project

#### DELETE /api/v1/projects/{id}

Delete a project and all its deployments.

**Parameters:**
- `id` (path): Project ID

**Response:**
```json
{
  "message": "Project deleted successfully"
}
```

---

## Deployments

### List Deployments

#### GET /api/v1/deployments

Retrieve all deployments.

**Response:**
```json
[
  {
    "id": 1,
    "project_id": 1,
    "environment_id": 3,
    "version": "v1.2.3",
    "status": "running",
    "url": "https://web-app.plate.local",
    "argo_app_name": "web-app-production",
    "helm_release": "web-app-production",
    "created_at": "2025-09-19T10:00:00Z",
    "updated_at": "2025-09-19T10:30:00Z",
    "project": {
      "id": 1,
      "name": "web-app"
    },
    "environment": {
      "id": 3,
      "name": "production"
    }
  }
]
```

### Create Deployment

#### POST /api/v1/deployments

Create a new deployment record.

**Request Body:**
```json
{
  "project_id": 1,
  "environment_id": 2,
  "version": "v1.3.0"
}
```

### Get Deployment

#### GET /api/v1/deployments/{id}

Retrieve a specific deployment.

**Parameters:**
- `id` (path): Deployment ID

### Delete Deployment

#### DELETE /api/v1/deployments/{id}

Delete a deployment and clean up resources.

**Parameters:**
- `id` (path): Deployment ID

### Get Deployment Logs

#### GET /api/v1/deployments/{id}/logs

Retrieve logs for a deployment.

**Parameters:**
- `id` (path): Deployment ID

**Response:**
```json
[
  {
    "id": 1,
    "deployment_id": 1,
    "level": "info",
    "message": "Deployment started",
    "timestamp": "2025-09-19T10:00:00Z"
  },
  {
    "id": 2,
    "deployment_id": 1,
    "level": "info",
    "message": "ArgoCD application created",
    "timestamp": "2025-09-19T10:01:00Z"
  }
]
```

### Deploy Application

#### POST /api/v1/deploy

Trigger a new deployment.

**Request Body:**
```json
{
  "project_id": 1,
  "environment_id": 3,
  "version": "v1.3.0"
}
```

**Response:**
```json
{
  "id": 7,
  "project_id": 1,
  "environment_id": 3,
  "version": "v1.3.0",
  "status": "pending",
  "argo_app_name": "web-app-production",
  "helm_release": "web-app-production",
  "created_at": "2025-09-19T12:00:00Z",
  "updated_at": "2025-09-19T12:00:00Z"
}
```

---

## Environments

### List Environments

#### GET /api/v1/environments

Retrieve all environments.

**Response:**
```json
[
  {
    "id": 1,
    "name": "development",
    "namespace": "plate-dev",
    "domain": "dev.plate.local",
    "created_at": "2025-09-19T08:00:00Z",
    "updated_at": "2025-09-19T08:00:00Z"
  },
  {
    "id": 2,
    "name": "staging",
    "namespace": "plate-staging", 
    "domain": "staging.plate.local",
    "created_at": "2025-09-19T08:00:00Z",
    "updated_at": "2025-09-19T08:00:00Z"
  },
  {
    "id": 3,
    "name": "production",
    "namespace": "plate-prod",
    "domain": "plate.local",
    "created_at": "2025-09-19T08:00:00Z",
    "updated_at": "2025-09-19T08:00:00Z"
  }
]
```

### Create Environment

#### POST /api/v1/environments

Create a new environment.

**Request Body:**
```json
{
  "name": "testing",
  "namespace": "plate-test",
  "domain": "test.plate.local"
}
```

### Get Environment

#### GET /api/v1/environments/{id}

Retrieve a specific environment.

**Parameters:**
- `id` (path): Environment ID

### Update Environment

#### PUT /api/v1/environments/{id}

Update an environment.

**Parameters:**
- `id` (path): Environment ID

**Request Body:**
```json
{
  "domain": "test.newdomain.com"
}
```

---

## Status

### Get Status

#### GET /api/v1/status

Get deployment status across environments.

**Query Parameters:**
- `env` (optional): Filter by environment name
- `detailed` (optional): Return detailed status (`true`/`false`)

**Simple Status Response:**
```json
{
  "web-app-production": "running",
  "web-app-staging": "running",
  "api-service-production": "running",
  "api-service-development": "failed"
}
```

**Detailed Status Response:**
```json
[
  {
    "environment": "production",
    "status": "running",
    "url": "https://web-app.plate.local",
    "pods": [
      {
        "name": "web-app-7d4b8f9c-abc12",
        "status": "Running",
        "ready": "1/1"
      }
    ],
    "last_update": "2025-09-19T10:30:00Z"
  }
]
```

## Status Codes

- `200` - Success
- `201` - Created
- `400` - Bad Request
- `404` - Not Found
- `500` - Internal Server Error

## Rate Limiting

Currently no rate limiting is implemented. In production, rate limiting will be applied per IP address and authentication token.

## Examples

### Using curl

```bash
# List all projects
curl http://localhost:8080/api/v1/projects

# Create a new project
curl -X POST http://localhost:8080/api/v1/projects \
  -H "Content-Type: application/json" \
  -d '{
    "name": "my-app",
    "description": "My application",
    "runtime": "nodejs",
    "port": 3000
  }'

# Deploy an application
curl -X POST http://localhost:8080/api/v1/deploy \
  -H "Content-Type: application/json" \
  -d '{
    "project_id": 1,
    "environment_id": 3,
    "version": "v1.0.0"
  }'

# Get deployment status
curl "http://localhost:8080/api/v1/status?detailed=true&env=production"
```

### Using JavaScript (Axios)

```javascript
// List projects
const projects = await axios.get('http://localhost:8080/api/v1/projects');

// Create project
const newProject = await axios.post('http://localhost:8080/api/v1/projects', {
  name: 'my-app',
  description: 'My application',
  runtime: 'nodejs',
  port: 3000
});

// Deploy application
const deployment = await axios.post('http://localhost:8080/api/v1/deploy', {
  project_id: 1,
  environment_id: 3,
  version: 'v1.0.0'
});
```
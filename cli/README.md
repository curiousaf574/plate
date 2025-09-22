# Plate CLI

Plate CLI is a command-line tool for deploying applications to Kubernetes using GitOps workflows.

## Installation

```bash
go build -o plate .
sudo mv plate /usr/local/bin/
```

## Usage

### Import a project
```bash
plate import .
plate import --name my-app --env production
```

### Deploy to Kubernetes
```bash
plate deploy
plate deploy --env production --watch
```

### Check deployment status
```bash
plate status
plate status --env production --detailed
```

## Configuration

Create a `.plate.yaml` file in your home directory:

```yaml
api-url: "https://api.plate.example.com"
token: "your-api-token"
```

## Project Structure

After importing, Plate creates a `.plate/` directory with:
- `config.yaml` - Project configuration
- Generated `Dockerfile` (if not present)

## Supported Runtimes

- Node.js
- Python
- Go
- Rust
- Java
- PHP
- Ruby
- Generic (custom)
# Getting Started with Plate

Welcome to Plate - the developer platform that makes deployment simple! This guide will get you up and running in just a few minutes.

## What is Plate?

Plate is a developer-first platform that eliminates deployment complexity. Just focus on your code - Plate handles the rest.

**Key Benefits:**
- 🚀 **Deploy in seconds** - From code to live URL instantly
- 🔍 **Zero configuration** - Automatically detects your project type
- 🌍 **Multiple environments** - Dev, staging, production made easy
- 📊 **Beautiful dashboard** - Monitor everything visually
- 🛠️ **Any language** - Node.js, Python, Go, Java, and more

## Quick Start (2 minutes)

### Prerequisites

- Your application code (any supported language)
- Node.js 18+ and npm (for the dashboard)
- Go 1.21+ (for the CLI)

### 1. Get Plate Running

```bash
# Clone Plate
git clone <repository-url>
cd plate

# Start the platform (in 3 terminals)

# Terminal 1: Start API
cd service && go run main.go serve

# Terminal 2: Start Dashboard  
cd ui && npm install && npm run dev

# Terminal 3: Build CLI
cd cli && go build -o plate .
```

### 2. Import Your First App

```bash
# Go to your project
cd /path/to/your-awesome-app

# Import it into Plate
../plate/cli/plate import

# Deploy to development
../plate/cli/plate deploy
```

### 3. See It Live!

- **Your App**: Check the URL provided by the deploy command
- **Dashboard**: Visit http://localhost:3000
- **API**: http://localhost:8080

That's it! Your app is now deployed and monitored.

## Supported Projects

Plate automatically works with:

### Web Applications
- **React**: Create React App, Next.js, Vite
- **Vue.js**: Vue CLI, Nuxt.js, Vite  
- **Angular**: Angular CLI projects
- **Static Sites**: HTML, CSS, JS files

### Backend APIs
- **Node.js**: Express, Fastify, Koa
- **Python**: Flask, Django, FastAPI
- **Go**: Gin, Echo, standard library
- **Java**: Spring Boot, Quarkus

### Full-Stack
- **MEAN/MERN**: Complete JavaScript stacks
- **Django**: Python web framework
- **Rails**: Ruby on Rails applications

### Examples by Project Type

**Node.js Project**
```bash
# Your package.json is automatically detected
plate import
plate deploy --env development
# ✅ Live at: https://my-app-dev.plate.local
```

**Python Project**
```bash
# Requirements.txt detected
plate import
plate deploy --env staging  
# ✅ Live at: https://my-app-staging.plate.local
```

**Go Project**
```bash
# Go.mod detected
plate import
plate deploy --env production
# ✅ Live at: https://my-app.plate.local
```

## How It Works

1. **Detection**: Plate scans your project and detects the runtime
2. **Configuration**: Automatically sets up build and start commands
3. **Deployment**: Packages and deploys your app
4. **URL Generation**: Provides instant live URLs
5. **Monitoring**: Tracks status and provides logs

## Environments Explained

### Development (`development`)
- **Purpose**: Active development and testing
- **URL Pattern**: `your-app-dev.yourcompany.com`
- **Auto-deploy**: On every push (optional)

### Staging (`staging`)  
- **Purpose**: Pre-production testing and demos
- **URL Pattern**: `your-app-staging.yourcompany.com`
- **Features**: Production-like environment

### Production (`production`)
- **Purpose**: Live application for users
- **URL Pattern**: `your-app.yourcompany.com`
- **Features**: High availability, monitoring, backups

## CLI Commands

### Basic Commands

```bash
# Import current directory
plate import

# Import specific directory  
plate import /path/to/project

# Import with custom name
plate import --name my-awesome-app

# Deploy to development (default)
plate deploy

# Deploy to specific environment
plate deploy --env production

# Check status
plate status

# Get detailed status
plate status --detailed
```

### Advanced Usage

```bash
# Deploy specific version
plate deploy --env production --version v1.2.3

# Watch deployment progress
plate deploy --env staging --watch

# Check specific environment
plate status --env production
```

## Dashboard Features

Visit http://localhost:3000 to access the web dashboard:

### Applications Page
- View all imported applications
- See deployment status across environments
- Quick deploy buttons
- Application health metrics

### Deployments Page  
- Monitor all deployments
- Filter by environment or status
- View deployment logs
- Quick actions (redeploy, view logs)

### Settings Page
- Configure default environments
- Set up notifications
- Manage access tokens
- Customize build settings

## Project Configuration

After importing, Plate creates a `.plate/` directory in your project:

```
your-project/
├── .plate/
│   └── config.yaml      # Plate configuration
├── your-app-files...
└── ...
```

### Configuration File

The `.plate/config.yaml` contains:

```yaml
name: my-app
runtime: nodejs
build_command: npm install && npm run build
start_command: npm start
port: 3000
environments:
  - development
  - staging  
  - production
env_vars:
  NODE_ENV: production
  API_URL: https://api.mycompany.com
```

You can customize this file to override defaults.

## Environment Variables

Set environment-specific variables:

```yaml
# .plate/config.yaml
env_vars:
  # Common variables
  APP_NAME: my-awesome-app
  
  # Environment-specific
  development:
    NODE_ENV: development
    API_URL: http://localhost:3000
    
  staging:
    NODE_ENV: staging  
    API_URL: https://api-staging.mycompany.com
    
  production:
    NODE_ENV: production
    API_URL: https://api.mycompany.com
```

## Troubleshooting

### Common Issues

**App won't start:**
```bash
# Check deployment logs
plate status --detailed

# View application logs  
plate logs --env development
```

**Build failing:**
```bash
# Check build configuration
cat .plate/config.yaml

# Test build locally
npm install && npm run build
```

**Environment issues:**
```bash
# List available environments
plate status

# Check specific environment
plate status --env production --detailed
```

### Getting Help

**Check logs:**
```bash
# Application logs
plate logs

# Build logs
plate logs --build

# Deployment logs  
plate logs --deploy
```

**Reset configuration:**
```bash
# Re-import project
plate import --force

# Reset to defaults
rm -rf .plate && plate import
```

## Next Steps

### For Development Teams
1. **Set up CI/CD**: Auto-deploy on git push
2. **Configure notifications**: Get alerts on deployment status
3. **Environment promotion**: Deploy dev → staging → production
4. **Team access**: Share projects with team members

### For Production Use
1. **Custom domains**: Set up your own domain names
2. **SSL certificates**: Automatic HTTPS for all environments
3. **Scaling**: Configure auto-scaling rules
4. **Monitoring**: Set up detailed application monitoring
5. **Backups**: Automatic backup and recovery

### Advanced Features
1. **Multi-service**: Deploy microservice architectures
2. **Database integration**: Connect to databases
3. **External services**: Integrate with third-party APIs
4. **Custom builds**: Use custom Docker images

## Examples and Templates

Check out example projects:

- [Node.js Express API](examples/nodejs-api/)
- [React Frontend App](examples/react-app/)  
- [Python Flask API](examples/python-flask/)
- [Go Web Service](examples/go-service/)
- [Full-Stack MERN](examples/mern-stack/)

## Support

Need help? We're here for you:

- **Documentation**: Complete guides in `/docs`
- **Examples**: Sample projects in `/examples`
- **Issues**: Report bugs on GitHub
- **Community**: Join our Discord server

**Ready to deploy your first app?** Run `plate import` in your project directory! 🚀

## Detailed Setup

### Local Development Environment

#### 1. Install Dependencies

**macOS:**
```bash
# Install Homebrew if not already installed
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

# Install required tools
brew install go node k3d helm kubectl
```

**Linux (Ubuntu/Debian):**
```bash
# Install Go
sudo apt update
sudo apt install golang-go

# Install Node.js
curl -fsSL https://deb.nodesource.com/setup_18.x | sudo -E bash -
sudo apt-get install -y nodejs

# Install kubectl
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl

# Install Helm
curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash

# Install k3d
curl -s https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | bash
```

#### 2. Setup Local Kubernetes

```bash
# Create local Kubernetes cluster
k3d cluster create plate-dev --port "8080:80@loadbalancer"

# Verify cluster is running
kubectl cluster-info
```

#### 3. Install Platform Dependencies

**ArgoCD:**
```bash
# Create namespace and install ArgoCD
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml

# Wait for ArgoCD to be ready
kubectl wait --for=condition=available --timeout=300s deployment/argocd-server -n argocd

# Get admin password
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d
echo  # for newline

# Port forward ArgoCD UI (optional)
kubectl port-forward svc/argocd-server -n argocd 8081:443 &
```

**Gitea:**
```bash
# Add Helm repository
helm repo add gitea-charts https://dl.gitea.io/charts/
helm repo update

# Install Gitea
helm install gitea gitea-charts/gitea \
  --namespace gitea \
  --create-namespace \
  --set service.http.type=LoadBalancer \
  --set gitea.admin.username=admin \
  --set gitea.admin.password=admin123 \
  --wait

# Port forward Gitea UI (optional)
kubectl port-forward svc/gitea-http -n gitea 3001:3000 &
```

### CLI Usage Examples

#### Import a Project

```bash
# Basic import (auto-detects runtime)
./plate import .

# Import with specific options
./plate import --name my-app --env production --runtime nodejs

# Import from specific directory
./plate import /path/to/project --name external-app
```

#### Deploy Applications

```bash
# Deploy to default environment (development)
./plate deploy

# Deploy to specific environment
./plate deploy --env production

# Deploy and watch progress
./plate deploy --env staging --watch
```

#### Check Status

```bash
# Get status of all deployments
./plate status

# Get detailed status
./plate status --detailed

# Get status for specific environment
./plate status --env production --detailed
```

#### Configuration

Create a configuration file at `~/.plate.yaml`:

```yaml
api-url: "http://localhost:8080"
token: ""  # Will be used when authentication is implemented
```

### Web UI Features

#### Dashboard
- Overview of all projects and deployments
- System health status
- Recent deployment activity
- Quick action buttons

#### Projects Page
- List all imported projects
- View project details and configuration
- Trigger deployments
- Manage project settings

#### Deployments Page
- Monitor all deployments across environments
- View deployment logs
- Rollback capabilities
- Filter by environment or status

#### Settings Page
- Configure platform settings
- Manage environments
- Set up integrations

### API Usage

#### Using curl

```bash
# Get all projects
curl http://localhost:8080/api/v1/projects

# Create a new project
curl -X POST http://localhost:8080/api/v1/projects \
  -H "Content-Type: application/json" \
  -d '{
    "name": "my-app",
    "description": "My awesome application",
    "runtime": "nodejs",
    "port": 3000
  }'

# Trigger deployment
curl -X POST http://localhost:8080/api/v1/deploy \
  -H "Content-Type: application/json" \
  -d '{
    "project_id": 1,
    "environment_id": 1,
    "version": "v1.0.0"
  }'

# Get deployment status
curl "http://localhost:8080/api/v1/status?detailed=true"
```

#### Using JavaScript

```javascript
// In your web application
const api = axios.create({
  baseURL: 'http://localhost:8080/api/v1'
});

// Get projects
const projects = await api.get('/projects');

// Deploy application
const deployment = await api.post('/deploy', {
  project_id: 1,
  environment_id: 2,
  version: 'v1.1.0'
});

// Monitor status
const status = await api.get('/status?detailed=true');
```

## Supported Runtimes

Plate automatically detects and supports the following runtimes:

### Node.js
- **Detection**: `package.json`
- **Build**: `npm install`
- **Start**: `npm start`
- **Port**: 3000

### Python
- **Detection**: `requirements.txt` or `Pipfile`
- **Build**: `pip install -r requirements.txt`
- **Start**: `python app.py`
- **Port**: 8000

### Go
- **Detection**: `go.mod`
- **Build**: `go build -o app .`
- **Start**: `./app`
- **Port**: 8080

### Java
- **Detection**: `pom.xml` or `build.gradle`
- **Build**: `mvn clean package`
- **Start**: `java -jar target/*.jar`
- **Port**: 8080

### PHP
- **Detection**: `composer.json`
- **Build**: `composer install`
- **Start**: `php -S 0.0.0.0:8000`
- **Port**: 8000

### Ruby
- **Detection**: `Gemfile`
- **Build**: `bundle install`
- **Start**: `bundle exec ruby app.rb`
- **Port**: 3000

### Rust
- **Detection**: `Cargo.toml`
- **Build**: `cargo build --release`
- **Start**: `./target/release/app`
- **Port**: 8080

## Project Structure

After importing a project, Plate creates:

```
your-project/
├── .plate/
│   └── config.yaml      # Project configuration
├── Dockerfile           # Generated if not present
└── ... your project files
```

### Configuration File

The `.plate/config.yaml` file contains:

```yaml
name: my-app
runtime: nodejs
environment: development
port: 3000
build_cmd: npm install
start_cmd: npm start
env_vars:
  NODE_ENV: production
  API_URL: https://api.example.com
```

## Environment Management

Plate supports multiple deployment environments:

### Default Environments

1. **Development** (`development`)
   - Namespace: `plate-dev`
   - Domain: `dev.plate.local`
   - Purpose: Active development and testing

2. **Staging** (`staging`)
   - Namespace: `plate-staging`
   - Domain: `staging.plate.local`
   - Purpose: Pre-production testing

3. **Production** (`production`)
   - Namespace: `plate-prod`
   - Domain: `plate.local`
   - Purpose: Live production environment

### Custom Environments

You can create additional environments through the API:

```bash
curl -X POST http://localhost:8080/api/v1/environments \
  -H "Content-Type: application/json" \
  -d '{
    "name": "testing",
    "namespace": "plate-test",
    "domain": "test.plate.local"
  }'
```

## Best Practices

### Project Organization
- Use descriptive project names
- Include meaningful descriptions
- Set appropriate environment variables
- Use semantic versioning for deployments

### Development Workflow
1. Import project with `plate import`
2. Deploy to development: `plate deploy --env development`
3. Test and validate
4. Deploy to staging: `plate deploy --env staging`
5. Production deployment: `plate deploy --env production`

### Monitoring
- Regularly check deployment status
- Monitor application logs
- Set up alerts for failed deployments
- Use the web dashboard for overview

## Troubleshooting

### Common Issues

**CLI Not Found:**
```bash
# Make sure the CLI is built and in PATH
cd cli
go build -o plate .
sudo mv plate /usr/local/bin/
```

**API Connection Failed:**
```bash
# Check if service is running
curl http://localhost:8080/health

# Check service logs
cd service
go run main.go serve
```

**UI Not Loading:**
```bash
# Check if Node.js dependencies are installed
cd ui
npm install

# Start development server
npm run dev
```

**Kubernetes Cluster Issues:**
```bash
# Check cluster status
kubectl cluster-info

# Check node status
kubectl get nodes

# Restart k3d cluster if needed
k3d cluster delete plate-dev
k3d cluster create plate-dev --port "8080:80@loadbalancer"
```

### Getting Help

- **Documentation**: Check the `/docs` folder
- **API Reference**: `/docs/api/README.md`
- **Architecture**: `/docs/architecture/README.md`
- **Deployment Guide**: `/docs/deployment/README.md`

## Next Steps

1. **Explore the Web UI** - Navigate through different sections
2. **Import Real Projects** - Try importing your existing applications
3. **Set up Production** - Follow the deployment guide for production setup
4. **Customize Configurations** - Modify project settings as needed
5. **Monitor Deployments** - Use the dashboard to track your applications

## Examples

Check the `/examples` directory (to be created) for sample projects demonstrating different runtimes and deployment scenarios.
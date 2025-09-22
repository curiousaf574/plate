# Plate - Developer Platform

Plate is a modern developer platform that simplifies application deployment. Focus on writing code while Plate handles the deployment complexity.

## 🚀 What is Plate?

Plate transforms the deployment experience for developers by providing:

- **One-Click Deployment** - Deploy your applications to any environment with a single command
- **Auto-Detection** - Automatically recognizes your project type and configures builds
- **Multi-Environment Support** - Seamlessly deploy to development, staging, and production
- **Real-Time Monitoring** - Monitor your applications with live status updates
- **Developer-First Design** - No infrastructure knowledge required

## ✨ Key Features

### For Developers
- 🔍 **Smart Detection** - Automatically detects Node.js, Python, Go, Java, and more
- 🚀 **Instant Deploy** - Deploy with `plate deploy` - that's it!
- 🌐 **Live URLs** - Get instant URLs for every deployment
- 📊 **Visual Dashboard** - Beautiful web interface to monitor your apps
- 🔄 **Environment Management** - Easy switching between dev, staging, and production

### For Teams
- 👥 **Collaborative** - Share deployments and manage environments together
- 🔒 **Secure** - Built-in security best practices
- 📈 **Scalable** - Handles projects of any size
- 🔔 **Notifications** - Get notified about deployment status
- 📝 **Audit Trail** - Complete deployment history and logs

## 🏃‍♂️ Quick Start

### 1. Install Plate CLI

```bash
# Build from source
git clone https://github.com/plate/cli
cd cli && go build -o plate .
```

### 2. Import Your Application

```bash
# Navigate to your project
cd my-awesome-app

# Import into Plate
plate import

# Deploy to development
plate deploy
```

### 3. Access Your App

Your application is now live! Visit the generated URL or check the web dashboard at `http://localhost:3000`.

## 📱 Web Dashboard

Access the beautiful web interface to:

- View all your applications
- Monitor deployment status
- Manage environments
- View deployment logs
- Configure settings

## 🛠️ Supported Technologies

Plate automatically detects and supports:

| Runtime | Detection | Build Command | Start Command |
|---------|-----------|---------------|---------------|
| **Node.js** | `package.json` | `npm install` | `npm start` |
| **Python** | `requirements.txt` | `pip install -r requirements.txt` | `python app.py` |
| **Go** | `go.mod` | `go build` | `./app` |
| **Java** | `pom.xml` | `mvn package` | `java -jar target/*.jar` |
| **PHP** | `composer.json` | `composer install` | `php -S 0.0.0.0:8000` |
| **Ruby** | `Gemfile` | `bundle install` | `bundle exec ruby app.rb` |
| **Rust** | `Cargo.toml` | `cargo build --release` | `./target/release/app` |

## 🎯 Use Cases

### Individual Developers
- **Personal Projects** - Deploy side projects and experiments
- **Portfolio Sites** - Showcase your work with live demos
- **API Testing** - Quickly deploy APIs for testing
- **Learning** - Experiment with new technologies

### Development Teams
- **Feature Branches** - Deploy every branch for review
- **Staging Environments** - Test before production
- **Client Demos** - Show progress to stakeholders
- **Microservices** - Manage multiple services easily

### Companies
- **Product Development** - Streamline deployment workflows
- **DevOps Automation** - Reduce deployment complexity
- **Multi-Environment** - Manage dev, staging, production
- **Team Collaboration** - Centralized deployment platform

## 📚 Documentation

- [**Getting Started**](docs/getting-started/README.md) - Complete setup guide
- [**CLI Reference**](docs/cli/README.md) - All CLI commands and options
- [**Web Dashboard**](docs/dashboard/README.md) - Using the web interface
- [**Deployment Guide**](docs/deployment/README.md) - Production deployment
- [**API Documentation**](docs/api/README.md) - REST API reference

## 🏗️ Architecture

Plate consists of three main components:

1. **CLI Tool** (`./cli`) - Command-line interface for developers
2. **Web Dashboard** (`./ui`) - Vue.js web interface for visual management
3. **API Service** (`./service`) - Go backend service handling deployments

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Plate CLI     │    │  Web Dashboard  │    │   API Service   │
│                 │    │                 │    │                 │
│ • plate import  │───▶│ • Applications  │───▶│ • Deployments   │
│ • plate deploy  │    │ • Deployments   │    │ • Environments  │
│ • plate status  │    │ • Monitoring    │    │ • Build System  │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

## 🔧 Development Setup

### Prerequisites

- Go 1.21+
- Node.js 18+
- Docker Desktop

### Local Development

1. **Start the API service**:
   ```bash
   cd service
   go run main.go serve
   ```

2. **Start the web dashboard**:
   ```bash
   cd ui
   npm install && npm run dev
   ```

3. **Build the CLI**:
   ```bash
   cd cli
   go build -o plate .
   ```

Access points:
- Web Dashboard: http://localhost:3000
- API Service: http://localhost:8080
- CLI: `./plate` command

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

## 📄 License

Plate is open source software licensed under the [MIT License](LICENSE).

## 🆘 Support

- **Documentation**: [docs/](docs/)
- **Issues**: [GitHub Issues](https://github.com/plate/plate/issues)
- **Discussions**: [GitHub Discussions](https://github.com/plate/plate/discussions)

---

**Ready to deploy?** Start with `plate import` in your project directory! 🚀
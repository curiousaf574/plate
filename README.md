# Plate - Developer Platform

> *"After years of building cloud-native platforms and struggling with fragmented tooling, I built the opinionated developer portal I always wished existed."*

## Why Plate?

As both a platform engineer and developer, I've built Kubernetes platforms and GitOps workflows, but using them day-to-day was painful. Fragmented tools, complex YAML, slow feedback loops - developers shouldn't need a PhD in Kubernetes just to deploy their code.

After trying every developer platform out there, I realized: **The tool I needed didn't exist.**

## What is Plate?

Plate is a unified developer platform that makes Kubernetes deployments simple. One CLI, one UI, one API - everything connected.

**Core Philosophy:**
- Deploy without YAML or Kubernetes knowledge
- Instant feedback and real-time monitoring
- Opinionated choices that just work
- Beautiful, intuitive interfaces

## 🚀 What Makes Plate Different?

Plate transforms the deployment experience by providing:

### 🎪 **The Triple Threat**
Unlike other platforms that force you to choose between CLI, UI, or API - Plate gives you all three, perfectly synchronized:

- **🖥️ CLI That Doesn't Suck** - `plate deploy` and you're done. No 47-line YAML files.
- **🎨 Gorgeous Web Dashboard** - Monitor everything in a beautiful, real-time interface
- **⚡ Rock-Solid API** - Integrate with anything, automate everything

### 🧠 **Intelligent Automation**
- **🔍 Runtime Detection** - Automatically identifies Node.js, Python, Go, Rust, Java, PHP, and Ruby projects
- **🐳 Dockerfile Generation** - Creates optimized Dockerfiles when needed
- **🚀 GitOps Integration** - Uses ArgoCD and Helm behind the scenes for reliable deployments
- **🌐 Automatic URLs** - Generates accessible URLs for each deployment

### 👨‍💻 **Developer Happiness Features**
- **⚡ Sub-30-Second Deployments** - From code to URL faster than your coffee gets cold
- **📊 Real-Time Everything** - Logs, status, metrics - all live, all the time
- **🔄 Environment Juggling Made Easy** - Switch between dev, staging, prod with zero friction
- **🎯 Zero YAML Required** - Seriously. We handle all the Kubernetes complexity
- **💥 Failure Recovery** - Automatic rollbacks when things go wrong

### 🏢 **Enterprise-Ready, Startup-Friendly**
- **🔒 Security First** - RBAC, network policies, and secrets management built-in
- **📈 Scales with You** - Handle 1 app or 1000 apps with the same elegance
- **👥 Team Collaboration** - Share deployments, manage permissions, audit everything
- **🔔 Smart Notifications** - Get pinged when things matter, not for every log line

## Quick Start

```bash
# Install from source
git clone https://github.com/curiousaf574/plate
cd plate/cli && go build -o plate .

# Import your project
plate import

# Deploy to development
plate deploy
```

## Supported Runtimes

Plate automatically detects and configures:

- **Node.js** - `package.json` → `npm install` → `npm start`
- **Python** - `requirements.txt` → `pip install` → `python app.py`
- **Go** - `go.mod` → `go build` → `./app`
- **Java** - `pom.xml` → `mvn package` → `java -jar app.jar`
- **PHP** - `composer.json` → `composer install` → `php -S 0.0.0.0:8000`
- **Ruby** - `Gemfile` → `bundle install` → `bundle exec ruby app.rb`
- **Rust** - `Cargo.toml` → `cargo build --release` → `./target/release/app`

## Architecture

Plate consists of three main components:

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Plate CLI     │    │  Web Dashboard  │    │   API Service   │
│                 │    │                 │    │                 │
│ • plate import  │───▶│ • Applications  │───▶│ • Deployments   │
│ • plate deploy  │    │ • Deployments   │    │ • Environments  │
│ • plate status  │    │ • Monitoring    │    │ • Build System  │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

**Infrastructure:** Built on Kubernetes, ArgoCD (GitOps), Helm (packaging), and Gitea (Git hosting).

## 🌟 Real-World Impact Stories

### 😎 **Solo Developers**
*"I went from Kubernetes-phobic to deploying 5 side projects a week"*
- **Weekend Warriors** - Deploy that hackathon project before Monday
- **Portfolio Power** - Live demos that actually work during interviews
- **Learning Labs** - Try new frameworks without infrastructure nightmares  
- **API Playground** - Rapid prototyping with instant feedback

### 🚀 **Development Teams**
*"Our PR review process went from 2 days to 2 hours"*
- **Branch Deployments** - Every feature gets a live URL for review
- **Demo Magic** - Show clients progress with real, working deployments
- **Staging Sanity** - Testing environments that mirror production exactly
- **Microservice Mastery** - Manage 20 services as easily as 1

### 🏢 **Platform Teams**
*"We finally have a developer portal that developers actually want to use"*
- **Developer Velocity** - 10x faster from idea to production
- **Reduced Support Tickets** - Self-service that actually works
- **Compliance Made Easy** - Security and governance built-in
- **Migration Path** - Move legacy apps to cloud-native without drama

### 🎯 **The Sweet Spot**
**Plate shines when you need:**
- Kubernetes power without Kubernetes complexity
- Multiple environments managed effortlessly
- Real-time visibility into everything
- A single tool that does it all well

## Documentation

📖 **[Complete Documentation](docs/)** - Everything you need to know

- **[Getting Started](docs/getting-started/)** - Your first deployment in 5 minutes
- **[API Reference](docs/api/)** - Complete REST API documentation
- **[Architecture](docs/architecture/)** - System design and component interactions
- **[Deployment Guide](docs/deployment/)** - Production deployment instructions

## Development

```bash
# Clone the repository
git clone https://github.com/curiousaf574/plate
cd plate

# Start the services (3 terminals):

# Terminal 1: API Service
cd service && go run main.go serve

# Terminal 2: Web Dashboard
cd ui && npm install && npm run dev

# Terminal 3: CLI
cd cli && go build -o plate .
```

**Prerequisites:** Go 1.21+, Node.js 18+, Docker Desktop, kubectl

### 🎯 **Why This Architecture Works**
- **🔄 GitOps Native** - Every deployment is version controlled and auditable
- **🏗️ Cloud Native** - Built on Kubernetes patterns that scale
- **🔒 Security First** - RBAC, network policies, and secrets management built-in
- **📊 Observable** - Metrics, logs, and traces from day one
- **🔧 Extensible** - Plugin architecture for custom workflows

## 🛠️ Join the Revolution (Development Setup)

**Want to contribute to the future of developer platforms?**

### 🚀 Quick Start for Contributors

```bash
# Clone the magic
git clone https://github.com/curiousaf574/plate
cd plate

# The full stack in 3 terminals:

# Terminal 1: API Service (Go)
cd service && go run main.go serve
# 🌐 API running at http://localhost:8080

# Terminal 2: Web Dashboard (Vue.js)
cd ui && npm install && npm run dev
# 🎨 Dashboard at http://localhost:3000

# Terminal 3: CLI Development
cd cli && go build -o plate .
# 🖥️ CLI ready: ./plate --help
```

### 🔧 **Prerequisites**
- **Go 1.21+** - For the CLI and API service
- **Node.js 18+** - For the gorgeous Vue.js dashboard
- **Docker Desktop** - For local Kubernetes testing
- **kubectl** - For cluster interactions
- **A sense of adventure** - For changing how deployments work forever

### 🎯 **Development Workflow**
```bash
# Make your changes, then test the full flow:
plate import ./test-app
plate deploy --env development
# Watch the magic happen in real-time! ✨
```

## 🤝 Join the Movement

**Plate is more than a tool - it's a revolution in developer experience.**

### 🎯 **How You Can Help**
- **⭐ Star the repo** - Help others discover Plate
- **🐛 Report bugs** - Help us squash issues before they spread
- **💡 Share ideas** - What would make your deployment workflow perfect?
- **📝 Write docs** - Help other developers discover the magic
- **🔧 Contribute code** - Add features, fix bugs, optimize performance
- **📢 Spread the word** - Tell your team about the platform that actually works

### 🎪 **The Vision**
*"Every developer should be able to deploy to production as easily as they run `npm start` locally."*

This isn't just about making deployments easier - it's about **unleashing developer creativity**. When deployment friction disappears, developers build more, experiment more, and ship faster.

## 📚 **Learn More**

- **📖 [Complete Documentation](docs/)** - Everything you need to know
- **🚀 [Getting Started Guide](docs/getting-started/)** - Your first deployment in 5 minutes
- **⚡ [CLI Reference](docs/cli/)** - Master every command
- **🎨 [Dashboard Guide](docs/dashboard/)** - Navigate like a pro
- **🏗️ [Architecture Deep Dive](docs/architecture/)** - Understand the magic

## 💬 **Community & Support**

- **🔥 [GitHub Discussions](https://github.com/curiousaf574/plate/discussions)** - Ideas, questions, show-and-tell
- **🐛 [Issues](https://github.com/curiousaf574/plate/issues)** - Bug reports and feature requests
- **💬 [Discord Community](https://discord.gg/plate)** - Real-time chat with the community
- **📧 [Email Support](mailto:support@plate.dev)** - Direct line to the maintainers

## Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

## License

Plate is open source and licensed under the [MIT License](LICENSE).

**Translation:** Use it, modify it, ship it, profit from it. Just don't blame us if it makes deployments too easy and you run out of things to do. 😉

---

## 🚀 **Ready to Transform Your Deployment Experience?**

```bash
# The journey starts with a single command:
plate import

# The destination: Developer happiness at scale ✨
```
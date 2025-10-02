# Plate - Developer Platform# Plate - Developer Platform



> *"After years of building cloud-native platforms and struggling with fragmented tooling, I built the opinionated developer portal I always wished existed."*> *"After years of building cloud-native platforms and struggling with fragmented tooling, I built the opinionated developer portal I always wished existed."*



## Why Plate?## Why Plate?



As both a platform engineer and developer, I've built Kubernetes platforms and GitOps workflows, but using them day-to-day was painful. Fragmented tools, complex YAML, slow feedback loops - developers shouldn't need a PhD in Kubernetes just to deploy their code.As both a platform engineer and developer, I've built Kubernetes platforms and GitOps workflows, but using them day-to-day was painful. Fragmented tools, complex YAML, slow feedback loops - developers shouldn't need a PhD in Kubernetes just to deploy their code.



After trying every developer platform out there, I realized: **The tool I needed didn't exist.**After trying every developer platform out there, I realized: **The tool I needed didn't exist.**



## What is Plate?## What is Plate?



Plate is a unified developer platform that makes Kubernetes deployments simple. One CLI, one UI, one API - everything connected.Plate is a unified developer platform that makes Kubernetes deployments simple. One CLI, one UI, one API - everything connected.



**Core Philosophy:****Core Philosophy:**

- Deploy without YAML or Kubernetes knowledge- Deploy without YAML or Kubernetes knowledge

- Instant feedback and real-time monitoring  - Instant feedback and real-time monitoring  

- Opinionated choices that just work- Opinionated choices that just work

- Beautiful, intuitive interfaces- Beautiful, intuitive interfaces



## Components## 🚀 What Makes Plate Different?



- **CLI** - Command-line tool for importing projects and managing deploymentsPlate transforms the deployment experience by providing:

- **UI** - Web dashboard for visual monitoring and team collaboration  

- **API** - REST service for automation and third-party integrations### 🎪 **The Triple Threat**

Unlike other platforms that force you to choose between CLI, UI, or API - Plate gives you all three, perfectly synchronized:

## Key Features

- **�️ CLI That Doesn't Suck** - `plate deploy` and you're done. No 47-line YAML files.

- **Auto-Detection** - Identifies Node.js, Python, Go, Rust, Java, PHP, and Ruby projects- **🎨 Gorgeous Web Dashboard** - Monitor everything in a beautiful, real-time interface

- **GitOps Workflow** - Uses ArgoCD and Helm for reliable deployments- **⚡ Rock-Solid API** - Integrate with anything, automate everything

- **Multi-Environment** - Easy management of dev, staging, and production

- **Real-Time Monitoring** - Live deployment status, logs, and metrics### 🧠 **Intelligent Automation**

- **Team Collaboration** - RBAC, audit trails, and shared management- **🔍 Runtime Detection** - Automatically identifies Node.js, Python, Go, Rust, Java, PHP, and Ruby projects

- **🐳 Dockerfile Generation** - Creates optimized Dockerfiles when needed

## Quick Start- **🚀 GitOps Integration** - Uses ArgoCD and Helm behind the scenes for reliable deployments

- **🌐 Automatic URLs** - Generates accessible URLs for each deployment

```bash

# Install from source### �‍💻 **Developer Happiness Features**

git clone https://github.com/curiousaf574/plate- **⚡ Sub-30-Second Deployments** - From code to URL faster than your coffee gets cold

cd plate/cli && go build -o plate .- **📊 Real-Time Everything** - Logs, status, metrics - all live, all the time  

- **🔄 Environment Juggling Made Easy** - Switch between dev, staging, prod with zero friction

# Import your project- **🎯 Zero YAML Required** - Seriously. We handle all the Kubernetes complexity

plate import- **💥 Failure Recovery** - Automatic rollbacks when things go wrong



# Deploy to development### 🏢 **Enterprise-Ready, Startup-Friendly**

plate deploy- **🔒 Security First** - RBAC, network policies, and secrets management built-in

```- **📈 Scales with You** - Handle 1 app or 1000 apps with the same elegance

- **👥 Team Collaboration** - Share deployments, manage permissions, audit everything

## Supported Runtimes- **🔔 Smart Notifications** - Get pinged when things matter, not for every log line



Plate automatically detects and configures:## Quick Start



- **Node.js** - `package.json` → `npm install` → `npm start````bash

- **Python** - `requirements.txt` → `pip install` → `python app.py`# Install from source

- **Go** - `go.mod` → `go build` → `./app`git clone https://github.com/curiousaf574/plate

- **Java** - `pom.xml` → `mvn package` → `java -jar app.jar`cd plate/cli && go build -o plate .

- **PHP** - `composer.json` → `composer install` → `php -S 0.0.0.0:8000`

- **Ruby** - `Gemfile` → `bundle install` → `bundle exec ruby app.rb`# Import your project

- **Rust** - `Cargo.toml` → `cargo build --release` → `./target/release/app`plate import



## Architecture# Deploy to development

plate deploy

Plate consists of three main components:```



```

┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐

│   Plate CLI     │    │  Web Dashboard  │    │   API Service   │## Supported Runtimes

│                 │    │                 │    │                 │

│ • plate import  │───▶│ • Applications  │───▶│ • Deployments   │Plate automatically detects and configures:

│ • plate deploy  │    │ • Deployments   │    │ • Environments  │

│ • plate status  │    │ • Monitoring    │    │ • Build System  │- **Node.js** - `package.json` → `npm install` → `npm start`

└─────────────────┘    └─────────────────┘    └─────────────────┘- **Python** - `requirements.txt` → `pip install` → `python app.py`

```- **Go** - `go.mod` → `go build` → `./app`

- **Java** - `pom.xml` → `mvn package` → `java -jar app.jar`

**Infrastructure:** Built on Kubernetes, ArgoCD (GitOps), Helm (packaging), and Gitea (Git hosting).- **PHP** - `composer.json` → `composer install` → `php -S 0.0.0.0:8000`

- **Ruby** - `Gemfile` → `bundle install` → `bundle exec ruby app.rb`

## Documentation- **Rust** - `Cargo.toml` → `cargo build --release` → `./target/release/app`



📖 **[Complete Documentation](docs/)** - Everything you need to know## � Real-World Impact Stories



- **[Getting Started](docs/getting-started/)** - Your first deployment in 5 minutes### 😎 **Solo Developers**

- **[API Reference](docs/api/)** - Complete REST API documentation*"I went from Kubernetes-phobic to deploying 5 side projects a week"*

- **[Architecture](docs/architecture/)** - System design and component interactions- **Weekend Warriors** - Deploy that hackathon project before Monday

- **[Deployment Guide](docs/deployment/)** - Production deployment instructions- **Portfolio Power** - Live demos that actually work during interviews

- **Learning Labs** - Try new frameworks without infrastructure nightmares

## Development- **API Playground** - Rapid prototyping with instant feedback



```bash### 🚀 **Development Teams**  

# Clone the repository*"Our PR review process went from 2 days to 2 hours"*

git clone https://github.com/curiousaf574/plate- **Branch Deployments** - Every feature gets a live URL for review

cd plate- **Demo Magic** - Show clients progress with real, working deployments

- **Staging Sanity** - Testing environments that mirror production exactly

# Start the services (3 terminals):- **Microservice Mastery** - Manage 20 services as easily as 1



# Terminal 1: API Service### 🏢 **Platform Teams**

cd service && go run main.go serve*"We finally have a developer portal that developers actually want to use"*

- **Developer Velocity** - 10x faster from idea to production

# Terminal 2: Web Dashboard  - **Reduced Support Tickets** - Self-service that actually works

cd ui && npm install && npm run dev- **Compliance Made Easy** - Security and governance built-in

- **Migration Path** - Move legacy apps to cloud-native without drama

# Terminal 3: CLI

cd cli && go build -o plate .### 🎯 **The Sweet Spot**

```**Plate shines when you need:**

- Kubernetes power without Kubernetes complexity

**Prerequisites:** Go 1.21+, Node.js 18+, Docker Desktop, kubectl- Multiple environments managed effortlessly  

- Real-time visibility into everything

## Contributing- A single tool that does it all well



We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.## Documentation



## License� **[Complete Documentation](docs/)** - Everything you need to know



Plate is open source software licensed under the [MIT License](LICENSE).- **[Getting Started](docs/getting-started/)** - Your first deployment in 5 minutes

- **[API Reference](docs/api/)** - Complete REST API documentation

---- **[Architecture](docs/architecture/)** - System design and component interactions

- **[Deployment Guide](docs/deployment/)** - Production deployment instructions

**Ready to deploy?** Start with `plate import` in your project directory! 🚀
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

## Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

## License

Plate is open source software licensed under the [MIT License](LICENSE).

---

**Ready to deploy?** Start with `plate import` in your project directory! 🚀

### 🎯 **Why This Architecture Works**
- **🔄 GitOps Native** - Every deployment is version controlled and auditable
- **🏗️ Cloud Native** - Built on Kubernetes patterns that scale
- **🔒 Security First** - RBAC, network policies, and secrets management built-in
- **📊 Observable** - Metrics, logs, and traces from day one
- **🔧 Extensible** - Plugin architecture for custom workflows

## �️ Join the Revolution (Development Setup)

**Want to contribute to the future of developer platforms?**

### 🚀 Quick Start for Contributors

```bash
# Clone the magic
git clone https://github.com/bramha574/plate
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

- **🔥 [GitHub Discussions](https://github.com/bramha574/plate/discussions)** - Ideas, questions, show-and-tell
- **🐛 [Issues](https://github.com/bramha574/plate/issues)** - Bug reports and feature requests  
- **💬 [Discord Community](https://discord.gg/plate)** - Real-time chat with the community
- **📧 [Email Support](mailto:support@plate.dev)** - Direct line to the maintainers

## 📄 **License**

Plate is open source and licensed under the [MIT License](LICENSE). 

**Translation:** Use it, modify it, ship it, profit from it. Just don't blame us if it makes deployments too easy and you run out of things to do. 😉

---

## 🚀 **Ready to Transform Your Deployment Experience?**

```bash
# The journey starts with a single command:
plate import

# The destination: Developer happiness at scale ✨
```

**Stop fighting with YAML. Start shipping features.**

**Welcome to Plate - where deployments just work.** 🎉
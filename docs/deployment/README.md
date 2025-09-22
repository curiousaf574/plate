# Plate Platform Deployment Guide

## Overview

This guide covers deploying the Plate platform in various environments, from local development to production Kubernetes clusters.

## Prerequisites

### Required Software
- **Kubernetes cluster** (v1.24+)
- **ArgoCD** (v2.8+)
- **Helm** (v3.12+)
- **Gitea** (v1.20+) or any Git repository hosting
- **PostgreSQL** (v13+)
- **Docker** (for building images)

### Infrastructure Requirements

#### Development
- Local Kubernetes (Docker Desktop, minikube, or k3d)
- 4GB RAM minimum
- 20GB storage

#### Production
- Kubernetes cluster with:
  - 3+ nodes
  - 8GB RAM per node minimum
  - 100GB storage per node
  - LoadBalancer support
  - Ingress controller

## Local Development Setup

### 1. Prerequisites Installation

#### Install k3d (Lightweight Kubernetes)
```bash
# Install k3d
curl -s https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | bash

# Create local cluster
k3d cluster create plate-dev --port "8080:80@loadbalancer" --port "8443:443@loadbalancer"
```

#### Install ArgoCD
```bash
# Create namespace
kubectl create namespace argocd

# Install ArgoCD
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml

# Get admin password
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d
```

#### Install Gitea
```bash
# Add Gitea Helm repository
helm repo add gitea-charts https://dl.gitea.io/charts/
helm repo update

# Install Gitea
helm install gitea gitea-charts/gitea \
  --namespace gitea \
  --create-namespace \
  --set service.http.type=LoadBalancer \
  --set gitea.admin.username=admin \
  --set gitea.admin.password=admin123
```

### 2. Run Plate Services

#### Start Backend Service
```bash
cd service
go mod download
go run main.go serve
```

#### Start Frontend UI
```bash
cd ui
npm install
npm run dev
```

#### Build CLI Tool
```bash
cd cli
go build -o plate .
./plate --help
```

### 3. Verify Installation

```bash
# Check Kubernetes cluster
kubectl get nodes

# Check ArgoCD
kubectl get pods -n argocd

# Check Gitea
kubectl get pods -n gitea

# Test Plate API
curl http://localhost:8080/health

# Test Plate UI
open http://localhost:3000
```

## Production Deployment

### 1. Infrastructure Setup

#### Kubernetes Cluster
Choose one of these options:

**Managed Kubernetes:**
- AWS EKS
- Google GKE
- Azure AKS
- DigitalOcean Kubernetes

**Self-managed:**
- kubeadm
- kops
- Rancher

#### Storage
```yaml
# StorageClass for PostgreSQL
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: fast-ssd
provisioner: kubernetes.io/aws-ebs
parameters:
  type: gp3
  fsType: ext4
```

#### Ingress Controller
```bash
# Install NGINX Ingress Controller
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm repo update

helm install ingress-nginx ingress-nginx/ingress-nginx \
  --namespace ingress-nginx \
  --create-namespace \
  --set controller.service.type=LoadBalancer
```

### 2. Install Core Components

#### PostgreSQL Database
```bash
# Add Bitnami repository
helm repo add bitnami https://charts.bitnami.com/bitnami

# Install PostgreSQL
helm install postgresql bitnami/postgresql \
  --namespace plate-system \
  --create-namespace \
  --set auth.postgresPassword=securepassword \
  --set auth.database=plate \
  --set primary.persistence.enabled=true \
  --set primary.persistence.size=100Gi \
  --set primary.persistence.storageClass=fast-ssd
```

#### ArgoCD
```bash
# Install ArgoCD with custom values
cat << EOF > argocd-values.yaml
server:
  ingress:
    enabled: true
    hosts:
      - argocd.yourdomain.com
    tls:
      - secretName: argocd-tls
        hosts:
          - argocd.yourdomain.com

configs:
  params:
    server.insecure: false
    server.grpc.web: true
EOF

helm repo add argo https://argoproj.github.io/argo-helm
helm install argocd argo/argo-cd \
  --namespace argocd \
  --create-namespace \
  --values argocd-values.yaml
```

#### Gitea
```bash
cat << EOF > gitea-values.yaml
gitea:
  admin:
    username: admin
    password: securepassword
    email: admin@yourdomain.com

postgresql:
  enabled: true
  auth:
    password: giteapassword
    database: gitea

service:
  http:
    type: LoadBalancer
  ssh:
    type: LoadBalancer

ingress:
  enabled: true
  hosts:
    - host: git.yourdomain.com
      paths:
        - path: /
          pathType: Prefix
  tls:
    - secretName: gitea-tls
      hosts:
        - git.yourdomain.com
EOF

helm install gitea gitea-charts/gitea \
  --namespace gitea \
  --create-namespace \
  --values gitea-values.yaml
```

### 3. Deploy Plate Services

#### Plate Service (Backend)
```bash
# Create deployment manifest
cat << EOF > plate-service.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: plate-service
  namespace: plate-system
spec:
  replicas: 3
  selector:
    matchLabels:
      app: plate-service
  template:
    metadata:
      labels:
        app: plate-service
    spec:
      containers:
      - name: plate-service
        image: plate/service:latest
        ports:
        - containerPort: 8080
        env:
        - name: DB_HOST
          value: postgresql.plate-system.svc.cluster.local
        - name: DB_PASSWORD
          valueFrom:
            secretKeyRef:
              name: postgresql
              key: postgres-password
        - name: ARGOCD_SERVER
          value: argocd-server.argocd.svc.cluster.local
        - name: GITEA_URL
          value: http://gitea-http.gitea.svc.cluster.local:3000
---
apiVersion: v1
kind: Service
metadata:
  name: plate-service
  namespace: plate-system
spec:
  selector:
    app: plate-service
  ports:
  - port: 8080
    targetPort: 8080
---
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: plate-service
  namespace: plate-system
  annotations:
    nginx.ingress.kubernetes.io/rewrite-target: /
spec:
  tls:
  - hosts:
    - api.yourdomain.com
    secretName: plate-api-tls
  rules:
  - host: api.yourdomain.com
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: plate-service
            port:
              number: 8080
EOF

kubectl apply -f plate-service.yaml
```

#### Plate UI (Frontend)
```bash
cat << EOF > plate-ui.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: plate-ui
  namespace: plate-system
spec:
  replicas: 3
  selector:
    matchLabels:
      app: plate-ui
  template:
    metadata:
      labels:
        app: plate-ui
    spec:
      containers:
      - name: plate-ui
        image: plate/ui:latest
        ports:
        - containerPort: 80
        env:
        - name: VITE_API_BASE_URL
          value: https://api.yourdomain.com
---
apiVersion: v1
kind: Service
metadata:
  name: plate-ui
  namespace: plate-system
spec:
  selector:
    app: plate-ui
  ports:
  - port: 80
    targetPort: 80
---
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: plate-ui
  namespace: plate-system
spec:
  tls:
  - hosts:
    - plate.yourdomain.com
    secretName: plate-ui-tls
  rules:
  - host: plate.yourdomain.com
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: plate-ui
            port:
              number: 80
EOF

kubectl apply -f plate-ui.yaml
```

### 4. Configuration

#### ArgoCD Configuration
```bash
# Create ArgoCD project for Plate
cat << EOF > argocd-project.yaml
apiVersion: argoproj.io/v1alpha1
kind: AppProject
metadata:
  name: plate
  namespace: argocd
spec:
  description: Plate platform applications
  sourceRepos:
  - 'https://git.yourdomain.com/plate/*'
  destinations:
  - namespace: 'plate-*'
    server: https://kubernetes.default.svc
  clusterResourceWhitelist:
  - group: ''
    kind: Namespace
  namespaceResourceWhitelist:
  - group: ''
    kind: Service
  - group: apps
    kind: Deployment
  - group: networking.k8s.io
    kind: Ingress
EOF

kubectl apply -f argocd-project.yaml
```

#### Gitea Configuration
```bash
# Create Plate organization in Gitea
curl -X POST "https://git.yourdomain.com/api/v1/orgs" \
  -H "Authorization: token YOUR_GITEA_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "plate",
    "full_name": "Plate Platform",
    "description": "Plate deployment platform repositories"
  }'
```

### 5. SSL/TLS Setup

#### Using cert-manager
```bash
# Install cert-manager
helm repo add jetstack https://charts.jetstack.io
helm repo update

helm install cert-manager jetstack/cert-manager \
  --namespace cert-manager \
  --create-namespace \
  --set installCRDs=true

# Create ClusterIssuer
cat << EOF > cluster-issuer.yaml
apiVersion: cert-manager.io/v1
kind: ClusterIssuer
metadata:
  name: letsencrypt-prod
spec:
  acme:
    server: https://acme-v02.api.letsencrypt.org/directory
    email: admin@yourdomain.com
    privateKeySecretRef:
      name: letsencrypt-prod
    solvers:
    - http01:
        ingress:
          class: nginx
EOF

kubectl apply -f cluster-issuer.yaml
```

## Monitoring & Observability

### Prometheus & Grafana
```bash
# Add Prometheus Helm repository
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts

# Install Prometheus stack
helm install monitoring prometheus-community/kube-prometheus-stack \
  --namespace monitoring \
  --create-namespace \
  --set grafana.adminPassword=admin123
```

### Logging with ELK Stack
```bash
# Add Elastic Helm repository
helm repo add elastic https://helm.elastic.co

# Install Elasticsearch
helm install elasticsearch elastic/elasticsearch \
  --namespace logging \
  --create-namespace

# Install Kibana
helm install kibana elastic/kibana \
  --namespace logging

# Install Filebeat
helm install filebeat elastic/filebeat \
  --namespace logging
```

## Backup & Recovery

### Database Backup
```bash
# Create backup CronJob
cat << EOF > postgres-backup.yaml
apiVersion: batch/v1
kind: CronJob
metadata:
  name: postgres-backup
  namespace: plate-system
spec:
  schedule: "0 2 * * *"
  jobTemplate:
    spec:
      template:
        spec:
          containers:
          - name: postgres-backup
            image: postgres:13
            command:
            - /bin/bash
            - -c
            - |
              pg_dump -h postgresql.plate-system.svc.cluster.local \
                -U postgres \
                -d plate \
                > /backup/plate-\$(date +%Y%m%d_%H%M%S).sql
            env:
            - name: PGPASSWORD
              valueFrom:
                secretKeyRef:
                  name: postgresql
                  key: postgres-password
            volumeMounts:
            - name: backup-storage
              mountPath: /backup
          volumes:
          - name: backup-storage
            persistentVolumeClaim:
              claimName: backup-pvc
          restartPolicy: OnFailure
EOF

kubectl apply -f postgres-backup.yaml
```

### Git Repository Backup
```bash
# Create backup script for Gitea repositories
cat << EOF > gitea-backup.sh
#!/bin/bash
GITEA_URL="https://git.yourdomain.com"
GITEA_TOKEN="YOUR_TOKEN"
BACKUP_DIR="/backup/git"

# Get all repositories
repos=\$(curl -s -H "Authorization: token \$GITEA_TOKEN" \
  "\$GITEA_URL/api/v1/orgs/plate/repos" | jq -r '.[].clone_url')

# Clone/update each repository
for repo in \$repos; do
  repo_name=\$(basename \$repo .git)
  if [ -d "\$BACKUP_DIR/\$repo_name" ]; then
    cd "\$BACKUP_DIR/\$repo_name" && git pull
  else
    git clone \$repo "\$BACKUP_DIR/\$repo_name"
  fi
done
EOF
```

## Scaling

### Horizontal Pod Autoscaler
```bash
cat << EOF > hpa.yaml
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: plate-service-hpa
  namespace: plate-system
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: plate-service
  minReplicas: 3
  maxReplicas: 10
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 70
  - type: Resource
    resource:
      name: memory
      target:
        type: Utilization
        averageUtilization: 80
EOF

kubectl apply -f hpa.yaml
```

### Cluster Autoscaler
```bash
# Enable cluster autoscaler (AWS example)
cat << EOF > cluster-autoscaler.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: cluster-autoscaler
  namespace: kube-system
spec:
  selector:
    matchLabels:
      app: cluster-autoscaler
  template:
    metadata:
      labels:
        app: cluster-autoscaler
    spec:
      serviceAccountName: cluster-autoscaler
      containers:
      - image: k8s.gcr.io/autoscaling/cluster-autoscaler:v1.21.0
        name: cluster-autoscaler
        command:
        - ./cluster-autoscaler
        - --v=4
        - --stderrthreshold=info
        - --cloud-provider=aws
        - --skip-nodes-with-local-storage=false
        - --expander=least-waste
        - --node-group-auto-discovery=asg:tag=k8s.io/cluster-autoscaler/enabled,k8s.io/cluster-autoscaler/plate-cluster
EOF

kubectl apply -f cluster-autoscaler.yaml
```

## Troubleshooting

### Common Issues

#### ArgoCD Sync Failures
```bash
# Check ArgoCD application status
kubectl get applications -n argocd

# View application details
kubectl describe application <app-name> -n argocd

# Check ArgoCD logs
kubectl logs -n argocd deployment/argocd-application-controller
```

#### Service Discovery Issues
```bash
# Check DNS resolution
kubectl run -it --rm debug --image=busybox --restart=Never -- nslookup plate-service.plate-system.svc.cluster.local

# Check service endpoints
kubectl get endpoints -n plate-system
```

#### Database Connection Issues
```bash
# Test PostgreSQL connection
kubectl run -it --rm psql --image=postgres:13 --restart=Never -- \
  psql -h postgresql.plate-system.svc.cluster.local -U postgres -d plate
```

### Log Collection
```bash
# Collect logs from all Plate components
kubectl logs -n plate-system deployment/plate-service > plate-service.log
kubectl logs -n plate-system deployment/plate-ui > plate-ui.log
kubectl logs -n argocd deployment/argocd-server > argocd-server.log
```

## Security Hardening

### Network Policies
```bash
cat << EOF > network-policies.yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: plate-service-netpol
  namespace: plate-system
spec:
  podSelector:
    matchLabels:
      app: plate-service
  policyTypes:
  - Ingress
  - Egress
  ingress:
  - from:
    - namespaceSelector:
        matchLabels:
          name: ingress-nginx
    ports:
    - protocol: TCP
      port: 8080
  egress:
  - to:
    - namespaceSelector:
        matchLabels:
          name: plate-system
    ports:
    - protocol: TCP
      port: 5432
EOF

kubectl apply -f network-policies.yaml
```

### Pod Security Standards
```bash
cat << EOF > pod-security.yaml
apiVersion: v1
kind: Namespace
metadata:
  name: plate-system
  labels:
    pod-security.kubernetes.io/enforce: restricted
    pod-security.kubernetes.io/audit: restricted
    pod-security.kubernetes.io/warn: restricted
EOF

kubectl apply -f pod-security.yaml
```
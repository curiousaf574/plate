<template>
  <div class="space-y-8 fade-in">
    <!-- Page Header -->
    <div class="page-header">
      <div class="flex items-center justify-between">
        <div>
          <h1 class="page-title">Dashboard</h1>
          <p class="page-subtitle">Monitor your applications and deployment activity</p>
        </div>
        <div class="flex items-center space-x-3">
          <button class="btn btn-secondary btn-sm">
            <ArrowPathIcon class="h-4 w-4 mr-2" />
            Refresh
          </button>
          <button class="btn btn-primary btn-sm">
            <PlusIcon class="h-4 w-4 mr-2" />
            New Deployment
          </button>
        </div>
      </div>
    </div>

    <!-- Quick Stats Grid -->
    <div class="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-4">
      <div 
        v-for="(stat, index) in stats" 
        :key="stat.name" 
        class="stat-card slide-up"
        :style="{ animationDelay: `${index * 100}ms` }"
      >
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-secondary-600">{{ stat.name }}</p>
            <p class="text-2xl font-bold text-secondary-900 mt-1">{{ stat.value }}</p>
            <p class="text-xs text-secondary-500 mt-1">{{ stat.change }}</p>
          </div>
          <div :class="[
            'stat-icon',
            stat.trend === 'up' ? 'stat-icon-success' : 
            stat.trend === 'down' ? 'stat-icon-warning' : 'stat-icon-primary'
          ]">
            <component :is="stat.icon" class="h-6 w-6" />
          </div>
        </div>
      </div>
    </div>

    <!-- Main Content Grid -->
    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <!-- Recent Activity -->
      <div class="lg:col-span-2 space-y-6">
        <!-- Recent Deployments -->
        <div class="card">
          <div class="card-header">
            <div class="flex items-center justify-between">
              <h2 class="text-lg font-semibold text-secondary-900">Recent Deployments</h2>
              <router-link to="/deployments" class="text-sm text-primary-600 hover:text-primary-700 font-medium">
                View all
              </router-link>
            </div>
          </div>
          <div class="card-body">
            <div class="space-y-4">
              <div 
                v-for="deployment in recentDeployments" 
                :key="deployment.id" 
                class="flex items-center justify-between p-3 rounded-lg border border-secondary-100 hover:border-secondary-200 transition-colors"
              >
                <div class="flex items-center space-x-3">
                  <div :class="[
                    'h-8 w-8 rounded-lg flex items-center justify-center text-xs font-medium',
                    deployment.status === 'success' ? 'bg-success-100 text-success-700' :
                    deployment.status === 'failed' ? 'bg-danger-100 text-danger-700' :
                    'bg-warning-100 text-warning-700'
                  ]">
                    <component :is="getStatusIcon(deployment.status)" class="h-4 w-4" />
                  </div>
                  <div>
                    <p class="font-medium text-secondary-900">{{ deployment.project }}</p>
                    <p class="text-sm text-secondary-500">{{ deployment.environment }} • {{ deployment.version }}</p>
                  </div>
                </div>
                <div class="text-right">
                  <span :class="[
                    'status-indicator',
                    deployment.status === 'success' ? 'status-success' :
                    deployment.status === 'failed' ? 'status-danger' :
                    'status-warning'
                  ]">
                    {{ deployment.status }}
                  </span>
                  <p class="text-xs text-secondary-500 mt-1">{{ deployment.time }}</p>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Performance Metrics -->
        <div class="card">
          <div class="card-header">
            <h2 class="text-lg font-semibold text-secondary-900">Performance Overview</h2>
          </div>
          <div class="card-body">
            <div class="grid grid-cols-1 sm:grid-cols-3 gap-4">
              <div v-for="metric in performanceMetrics" :key="metric.name" class="text-center">
                <div class="text-2xl font-bold text-secondary-900">{{ metric.value }}</div>
                <div class="text-sm text-secondary-600">{{ metric.name }}</div>
                <div class="mt-2 w-full bg-secondary-200 rounded-full h-2">
                  <div 
                    :class="[
                      'h-2 rounded-full transition-all duration-300',
                      metric.status === 'good' ? 'bg-success-500' :
                      metric.status === 'warning' ? 'bg-warning-500' :
                      'bg-danger-500'
                    ]"
                    :style="{ width: `${metric.percentage}%` }"
                  ></div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Sidebar Widgets -->
      <div class="space-y-6">
        <!-- System Health -->
        <div class="card">
          <div class="card-header">
            <h2 class="text-lg font-semibold text-secondary-900">System Health</h2>
          </div>
          <div class="card-body">
            <div class="space-y-3">
              <div v-for="service in systemStatus" :key="service.name" class="flex items-center justify-between">
                <div class="flex items-center space-x-3">
                  <div :class="[
                    'h-2 w-2 rounded-full',
                    service.status === 'healthy' ? 'bg-success-500' :
                    service.status === 'warning' ? 'bg-warning-500' :
                    'bg-danger-500'
                  ]"></div>
                  <span class="text-sm font-medium text-secondary-900">{{ service.name }}</span>
                </div>
                <span :class="[
                  'text-xs font-medium px-2 py-1 rounded-full',
                  service.status === 'healthy' ? 'bg-success-100 text-success-700' :
                  service.status === 'warning' ? 'bg-warning-100 text-warning-700' :
                  'bg-danger-100 text-danger-700'
                ]">
                  {{ service.status }}
                </span>
              </div>
            </div>
          </div>
        </div>

        <!-- Quick Actions -->
        <div class="card">
          <div class="card-header">
            <h2 class="text-lg font-semibold text-secondary-900">Quick Actions</h2>
          </div>
          <div class="card-body">
            <div class="space-y-2">
              <button class="w-full btn btn-secondary btn-sm justify-start">
                <RocketLaunchIcon class="h-4 w-4 mr-2" />
                Deploy Application
              </button>
              <button class="w-full btn btn-secondary btn-sm justify-start">
                <FolderIcon class="h-4 w-4 mr-2" />
                Import Project
              </button>
              <button class="w-full btn btn-secondary btn-sm justify-start">
                <CogIcon class="h-4 w-4 mr-2" />
                System Settings
              </button>
            </div>
          </div>
        </div>

        <!-- Environment Status -->
        <div class="card">
          <div class="card-header">
            <h2 class="text-lg font-semibold text-secondary-900">Environments</h2>
          </div>
          <div class="card-body">
            <div class="space-y-3">
              <div v-for="env in environments" :key="env.name" class="flex items-center justify-between">
                <div>
                  <p class="text-sm font-medium text-secondary-900">{{ env.name }}</p>
                  <p class="text-xs text-secondary-500">{{ env.apps }} apps</p>
                </div>
                <div :class="[
                  'h-8 w-8 rounded-lg flex items-center justify-center',
                  env.status === 'active' ? 'bg-success-100' :
                  'bg-secondary-100'
                ]">
                  <div :class="[
                    'h-2 w-2 rounded-full',
                    env.status === 'active' ? 'bg-success-500' : 'bg-secondary-400'
                  ]"></div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { 
  FolderIcon, 
  RocketLaunchIcon, 
  ServerIcon, 
  ClockIcon,
  CheckCircleIcon,
  XCircleIcon,
  ExclamationTriangleIcon,
  PlusIcon,
  ArrowPathIcon,
  CogIcon
} from '@heroicons/vue/24/outline'

const stats = ref([
  { 
    name: 'Applications', 
    value: '12', 
    change: '+2 this week',
    trend: 'up',
    icon: FolderIcon 
  },
  { 
    name: 'Active Deployments', 
    value: '8', 
    change: '+1 today',
    trend: 'up',
    icon: RocketLaunchIcon 
  },
  { 
    name: 'Healthy Services', 
    value: '24/25', 
    change: '96% uptime',
    trend: 'stable',
    icon: ServerIcon 
  },
  { 
    name: 'Avg Deploy Time', 
    value: '2.4m', 
    change: '-30s improved',
    trend: 'up',
    icon: ClockIcon 
  },
])

const recentDeployments = ref([
  { 
    id: 1, 
    project: 'web-frontend', 
    environment: 'production', 
    version: 'v1.2.3',
    status: 'success', 
    time: '2 minutes ago' 
  },
  { 
    id: 2, 
    project: 'api-gateway', 
    environment: 'staging', 
    version: 'v2.1.0',
    status: 'running', 
    time: '5 minutes ago' 
  },
  { 
    id: 3, 
    project: 'worker-service', 
    environment: 'production', 
    version: 'v1.0.8',
    status: 'success', 
    time: '12 minutes ago' 
  },
  { 
    id: 4, 
    project: 'auth-service', 
    environment: 'development', 
    version: 'v3.0.1',
    status: 'failed', 
    time: '18 minutes ago' 
  },
])

const systemStatus = ref([
  { name: 'Deployment Engine', status: 'healthy' },
  { name: 'Git Repository', status: 'healthy' },
  { name: 'Build Pipeline', status: 'healthy' },
  { name: 'Monitoring', status: 'warning' },
  { name: 'Load Balancer', status: 'healthy' },
])

const performanceMetrics = ref([
  { name: 'Success Rate', value: '98.5%', percentage: 98.5, status: 'good' },
  { name: 'Avg Response', value: '120ms', percentage: 85, status: 'good' },
  { name: 'Error Rate', value: '0.2%', percentage: 15, status: 'good' },
])

const environments = ref([
  { name: 'Production', apps: 8, status: 'active' },
  { name: 'Staging', apps: 12, status: 'active' },
  { name: 'Development', apps: 15, status: 'active' },
  { name: 'Testing', apps: 5, status: 'inactive' },
])

const getStatusIcon = (status) => {
  switch (status) {
    case 'success':
      return CheckCircleIcon
    case 'failed':
      return XCircleIcon
    default:
      return ExclamationTriangleIcon
  }
}
</script>
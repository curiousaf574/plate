<template>
  <div class="space-y-8 fade-in">
    <!-- Page Header -->
    <div class="page-header">
      <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4">
        <div>
          <h1 class="page-title">Deployments</h1>
          <p class="page-subtitle">Monitor your application deployments across all environments</p>
        </div>
        <div class="flex items-center space-x-3">
          <button @click="fetchDeployments" class="btn btn-secondary btn-sm">
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

    <!-- Stats Summary -->
    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="stat-card">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-secondary-600">Total Deployments</p>
            <p class="text-2xl font-bold text-secondary-900">{{ filteredDeployments.length }}</p>
          </div>
          <div class="stat-icon-primary">
            <RocketLaunchIcon class="h-5 w-5" />
          </div>
        </div>
      </div>
      <div class="stat-card">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-secondary-600">Running</p>
            <p class="text-2xl font-bold text-secondary-900">{{ runningDeployments }}</p>
          </div>
          <div class="stat-icon-success">
            <CheckCircleIcon class="h-5 w-5" />
          </div>
        </div>
      </div>
      <div class="stat-card">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-secondary-600">Pending</p>
            <p class="text-2xl font-bold text-secondary-900">{{ pendingDeployments }}</p>
          </div>
          <div class="stat-icon-warning">
            <ClockIcon class="h-5 w-5" />
          </div>
        </div>
      </div>
      <div class="stat-card">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-secondary-600">Failed</p>
            <p class="text-2xl font-bold text-secondary-900">{{ failedDeployments }}</p>
          </div>
          <div class="stat-icon stat-icon bg-danger-100 text-danger-600">
            <ExclamationTriangleIcon class="h-5 w-5" />
          </div>
        </div>
      </div>
    </div>

    <!-- Filter Bar -->
    <div class="flex flex-col sm:flex-row gap-4">
      <div class="relative flex-1 max-w-md">
        <MagnifyingGlassIcon class="absolute left-3 top-1/2 transform -translate-y-1/2 h-4 w-4 text-secondary-400" />
        <input 
          v-model="searchQuery"
          type="text" 
          placeholder="Search deployments..." 
          class="form-input pl-10"
        />
      </div>
      <select v-model="selectedEnvironment" class="form-input max-w-48">
        <option value="">All Environments</option>
        <option value="development">Development</option>
        <option value="staging">Staging</option>
        <option value="production">Production</option>
      </select>
      <select v-model="selectedStatus" class="form-input max-w-32">
        <option value="">All Status</option>
        <option value="running">Running</option>
        <option value="success">Success</option>
        <option value="failed">Failed</option>
        <option value="pending">Pending</option>
      </select>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="flex flex-col items-center justify-center py-16">
      <div class="animate-spin rounded-full h-8 w-8 border-2 border-primary-600 border-t-transparent"></div>
      <p class="mt-4 text-sm text-secondary-500">Loading deployments...</p>
    </div>

    <!-- Error State -->
    <div v-else-if="error" class="text-center py-16">
      <div class="mx-auto h-12 w-12 rounded-full bg-danger-100 flex items-center justify-center">
        <ExclamationTriangleIcon class="h-6 w-6 text-danger-600" />
      </div>
      <h3 class="mt-4 text-lg font-medium text-secondary-900">Failed to load deployments</h3>
      <p class="mt-2 text-sm text-secondary-500">{{ error }}</p>
      <button @click="fetchDeployments" class="mt-4 btn btn-primary btn-sm">
        <ArrowPathIcon class="h-4 w-4 mr-2" />
        Retry
      </button>
    </div>

    <!-- Deployments Grid -->
    <div v-else-if="filteredDeployments.length > 0" class="grid grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 gap-6">
      <div 
        v-for="(deployment, index) in filteredDeployments" 
        :key="`${deployment.project}-${deployment.environment}`" 
        class="card-hover slide-up"
        :style="{ animationDelay: `${index * 50}ms` }"
      >
        <div class="card-body">
          <!-- Deployment Header -->
          <div class="flex items-start justify-between mb-4">
            <div class="flex items-center space-x-3">
              <div :class="[
                'h-10 w-10 rounded-xl flex items-center justify-center text-sm font-medium',
                getEnvironmentColor(deployment.environment)
              ]">
                {{ deployment.environment.substring(0, 2).toUpperCase() }}
              </div>
              <div>
                <h3 class="text-lg font-semibold text-secondary-900">{{ deployment.project }}</h3>
                <p class="text-sm text-secondary-500">{{ deployment.environment }}</p>
              </div>
            </div>
            <Menu as="div" class="relative">
              <MenuButton class="p-1 hover:bg-secondary-100 rounded-lg transition-colors">
                <EllipsisVerticalIcon class="h-4 w-4 text-secondary-400" />
              </MenuButton>
              <MenuItems class="absolute right-0 z-10 mt-2 w-48 origin-top-right rounded-xl bg-white py-2 shadow-large ring-1 ring-secondary-900/5 focus:outline-none">
                <MenuItem v-slot="{ active }">
                  <button :class="[active ? 'bg-secondary-50' : '', 'block w-full text-left px-4 py-2 text-sm text-secondary-700']">
                    View Logs
                  </button>
                </MenuItem>
                <MenuItem v-slot="{ active }">
                  <button :class="[active ? 'bg-secondary-50' : '', 'block w-full text-left px-4 py-2 text-sm text-secondary-700']">
                    Redeploy
                  </button>
                </MenuItem>
                <MenuItem v-slot="{ active }">
                  <button :class="[active ? 'bg-danger-50' : '', 'block w-full text-left px-4 py-2 text-sm text-danger-700']">
                    Stop
                  </button>
                </MenuItem>
              </MenuItems>
            </Menu>
          </div>

          <!-- Status Badge -->
          <div class="mb-4">
            <span :class="[
              'status-indicator',
              deployment.status === 'running' ? 'status-success' :
              deployment.status === 'success' ? 'status-info' :
              deployment.status === 'failed' ? 'status-danger' :
              'status-warning'
            ]">
              {{ deployment.status === 'running' ? 'live' : deployment.status }}
            </span>
          </div>
          
          <!-- Deployment Details -->
          <div class="space-y-3 mb-6">
            <div class="flex items-center justify-between text-sm">
              <span class="text-secondary-500">URL:</span>
              <div class="text-right">
                <a v-if="deployment.url" :href="deployment.url" target="_blank" class="text-primary-600 hover:text-primary-700 font-medium">
                  Visit Site
                  <ArrowTopRightOnSquareIcon class="h-3 w-3 inline ml-1" />
                </a>
                <span v-else class="text-secondary-400">No URL</span>
              </div>
            </div>
            <div class="flex items-center justify-between text-sm">
              <span class="text-secondary-500">Deployed:</span>
              <span class="text-secondary-900 font-medium">{{ deployment.deployedAt }}</span>
            </div>
            <div class="flex items-center justify-between text-sm">
              <span class="text-secondary-500">Version:</span>
              <span class="text-secondary-900 font-medium font-mono text-xs">{{ deployment.version }}</span>
            </div>
            <div class="flex items-center justify-between text-sm">
              <span class="text-secondary-500">Duration:</span>
              <span class="text-secondary-900 font-medium">{{ deployment.duration || 'N/A' }}</span>
            </div>
          </div>
          
          <!-- Action Buttons -->
          <div class="flex space-x-2">
            <button class="btn btn-primary btn-sm flex-1">
              <DocumentTextIcon class="h-4 w-4 mr-2" />
              View Logs
            </button>
            <button class="btn btn-secondary btn-sm">
              <ArrowPathIcon class="h-4 w-4" />
            </button>
            <button class="btn btn-secondary btn-sm">
              <ChartBarIcon class="h-4 w-4" />
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Empty State -->
    <div v-else class="text-center py-16">
      <div class="mx-auto h-16 w-16 rounded-full bg-secondary-100 flex items-center justify-center mb-4">
        <RocketLaunchIcon class="h-8 w-8 text-secondary-400" />
      </div>
      <h3 class="text-lg font-medium text-secondary-900">No deployments found</h3>
      <p class="mt-2 text-sm text-secondary-500">
        {{ searchQuery ? 'Try adjusting your search terms or filters.' : 'Deploy your first application to get started.' }}
      </p>
      <div class="mt-6">
        <button v-if="!searchQuery" class="btn btn-primary btn-md">
          <PlusIcon class="h-4 w-4 mr-2" />
          Create Deployment
        </button>
        <button v-else @click="clearSearch" class="btn btn-secondary btn-md">
          Clear Search
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { Menu, MenuButton, MenuItem, MenuItems } from '@headlessui/vue'
import { 
  RocketLaunchIcon,
  PlusIcon,
  ArrowPathIcon,
  CheckCircleIcon,
  ExclamationTriangleIcon,
  ClockIcon,
  MagnifyingGlassIcon,
  DocumentTextIcon,
  ChartBarIcon,
  EllipsisVerticalIcon,
  ArrowTopRightOnSquareIcon
} from '@heroicons/vue/24/outline'

const selectedEnvironment = ref('')
const selectedStatus = ref('')
const searchQuery = ref('')
const deployments = ref([])
const loading = ref(true)
const error = ref(null)

// Mock data for demo - replace with API call
const mockDeployments = [
  {
    project: 'web-frontend',
    environment: 'production',
    status: 'running',
    url: 'https://app.plate.dev',
    deployedAt: '2 hours ago',
    version: 'v1.2.3',
    duration: '2m 34s'
  },
  {
    project: 'api-gateway',
    environment: 'staging',
    status: 'running',
    url: 'https://staging-api.plate.dev',
    deployedAt: '5 minutes ago',
    version: 'v2.1.0',
    duration: '1m 45s'
  },
  {
    project: 'auth-service',
    environment: 'development',
    status: 'failed',
    url: null,
    deployedAt: '1 day ago',
    version: 'v1.5.2',
    duration: '45s'
  },
  {
    project: 'data-processor',
    environment: 'production',
    status: 'success',
    url: null,
    deployedAt: '3 days ago',
    version: 'v0.8.1',
    duration: '3m 12s'
  },
  {
    project: 'notification-service',
    environment: 'staging',
    status: 'pending',
    url: null,
    deployedAt: 'deploying...',
    version: 'v1.0.6',
    duration: '1m 20s'
  }
]

const fetchDeployments = async () => {
  try {
    loading.value = true
    error.value = null
    
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 800))
    
    // In a real app, replace this with:
    // const response = await fetch('http://localhost:8080/api/v1/deployments')
    // if (!response.ok) throw new Error(`HTTP error! status: ${response.status}`)
    // const data = await response.json()
    
    deployments.value = mockDeployments
  } catch (err) {
    error.value = err.message
    console.error('Failed to fetch deployments:', err)
  } finally {
    loading.value = false
  }
}

const filteredDeployments = computed(() => {
  let filtered = deployments.value

  // Filter by search query
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    filtered = filtered.filter(deployment => 
      deployment.project.toLowerCase().includes(query) ||
      deployment.environment.toLowerCase().includes(query)
    )
  }

  // Filter by environment and status
  return filtered.filter(deployment => {
    const environmentMatch = !selectedEnvironment.value || deployment.environment === selectedEnvironment.value
    const statusMatch = !selectedStatus.value || deployment.status === selectedStatus.value
    return environmentMatch && statusMatch
  })
})

const runningDeployments = computed(() => deployments.value.filter(d => d.status === 'running').length)
const pendingDeployments = computed(() => deployments.value.filter(d => d.status === 'pending').length)
const failedDeployments = computed(() => deployments.value.filter(d => d.status === 'failed').length)

const clearSearch = () => {
  searchQuery.value = ''
  selectedEnvironment.value = ''
  selectedStatus.value = ''
}

const getEnvironmentColor = (environment) => {
  const colors = {
    production: 'bg-danger-100 text-danger-700',
    staging: 'bg-warning-100 text-warning-700',
    development: 'bg-primary-100 text-primary-700',
    testing: 'bg-secondary-100 text-secondary-700'
  }
  return colors[environment] || 'bg-secondary-100 text-secondary-700'
}

onMounted(() => {
  fetchDeployments()
})
</script>
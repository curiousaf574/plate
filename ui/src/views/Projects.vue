<template>
  <div class="space-y-8 fade-in">
    <!-- Page Header -->
    <div class="page-header">
      <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4">
        <div>
          <h1 class="page-title">Applications</h1>
          <p class="page-subtitle">Manage and deploy your applications across environments</p>
        </div>
        <div class="flex items-center space-x-3">
          <div class="relative">
            <MagnifyingGlassIcon class="absolute left-3 top-1/2 transform -translate-y-1/2 h-4 w-4 text-secondary-400" />
            <input 
              v-model="searchQuery"
              type="text" 
              placeholder="Search applications..." 
              class="form-input pl-10 w-64"
            />
          </div>
          <Menu as="div" class="relative">
            <MenuButton class="btn btn-secondary btn-sm">
              <FunnelIcon class="h-4 w-4 mr-2" />
              Filter
              <ChevronDownIcon class="h-4 w-4 ml-1" />
            </MenuButton>
            <MenuItems class="absolute right-0 z-10 mt-2 w-48 origin-top-right rounded-xl bg-white py-2 shadow-large ring-1 ring-secondary-900/5 focus:outline-none">
              <MenuItem v-for="filter in filters" :key="filter.key" v-slot="{ active }">
                <button 
                  @click="selectedFilter = filter.key"
                  :class="[
                    active ? 'bg-secondary-50' : '', 
                    selectedFilter === filter.key ? 'text-primary-700 font-medium' : 'text-secondary-700',
                    'block w-full text-left px-4 py-2 text-sm'
                  ]"
                >
                  {{ filter.label }}
                </button>
              </MenuItem>
            </MenuItems>
          </Menu>
          <button class="btn btn-primary btn-sm">
            <PlusIcon class="h-4 w-4 mr-2" />
            Import Application
          </button>
        </div>
      </div>
    </div>

    <!-- Stats Summary -->
    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="stat-card">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-secondary-600">Total Apps</p>
            <p class="text-2xl font-bold text-secondary-900">{{ filteredProjects.length }}</p>
          </div>
          <div class="stat-icon-primary">
            <FolderIcon class="h-5 w-5" />
          </div>
        </div>
      </div>
      <div class="stat-card">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-secondary-600">Active</p>
            <p class="text-2xl font-bold text-secondary-900">{{ activeApps }}</p>
          </div>
          <div class="stat-icon-success">
            <CheckCircleIcon class="h-5 w-5" />
          </div>
        </div>
      </div>
      <div class="stat-card">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-secondary-600">Deploying</p>
            <p class="text-2xl font-bold text-secondary-900">{{ deployingApps }}</p>
          </div>
          <div class="stat-icon-warning">
            <ArrowPathIcon class="h-5 w-5" />
          </div>
        </div>
      </div>
      <div class="stat-card">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-secondary-600">Issues</p>
            <p class="text-2xl font-bold text-secondary-900">{{ failedApps }}</p>
          </div>
          <div class="stat-icon stat-icon bg-danger-100 text-danger-600">
            <ExclamationTriangleIcon class="h-5 w-5" />
          </div>
        </div>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="flex flex-col items-center justify-center py-16">
      <div class="animate-spin rounded-full h-8 w-8 border-2 border-primary-600 border-t-transparent"></div>
      <p class="mt-4 text-sm text-secondary-500">Loading applications...</p>
    </div>

    <!-- Error State -->
    <div v-else-if="error" class="text-center py-16">
      <div class="mx-auto h-12 w-12 rounded-full bg-danger-100 flex items-center justify-center">
        <ExclamationTriangleIcon class="h-6 w-6 text-danger-600" />
      </div>
      <h3 class="mt-4 text-lg font-medium text-secondary-900">Failed to load applications</h3>
      <p class="mt-2 text-sm text-secondary-500">{{ error }}</p>
      <button @click="fetchProjects" class="mt-4 btn btn-primary btn-sm">
        <ArrowPathIcon class="h-4 w-4 mr-2" />
        Retry
      </button>
    </div>

    <!-- Applications Grid -->
    <div v-else-if="filteredProjects.length > 0" class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-6">
      <div 
        v-for="(project, index) in filteredProjects" 
        :key="project.id" 
        class="card-hover slide-up"
        :style="{ animationDelay: `${index * 50}ms` }"
      >
        <div class="card-body">
          <!-- Project Header -->
          <div class="flex items-start justify-between mb-4">
            <div class="flex items-center space-x-3">
              <div :class="[
                'h-10 w-10 rounded-xl flex items-center justify-center text-sm font-medium',
                getProjectTypeColor(project.runtime)
              ]">
                {{ getProjectIcon(project.runtime) }}
              </div>
              <div>
                <h3 class="text-lg font-semibold text-secondary-900">{{ project.name }}</h3>
                <p class="text-sm text-secondary-500">{{ project.runtime }}</p>
              </div>
            </div>
            <Menu as="div" class="relative">
              <MenuButton class="p-1 hover:bg-secondary-100 rounded-lg transition-colors">
                <EllipsisVerticalIcon class="h-4 w-4 text-secondary-400" />
              </MenuButton>
              <MenuItems class="absolute right-0 z-10 mt-2 w-48 origin-top-right rounded-xl bg-white py-2 shadow-large ring-1 ring-secondary-900/5 focus:outline-none">
                <MenuItem v-slot="{ active }">
                  <button :class="[active ? 'bg-secondary-50' : '', 'block w-full text-left px-4 py-2 text-sm text-secondary-700']">
                    View Details
                  </button>
                </MenuItem>
                <MenuItem v-slot="{ active }">
                  <button :class="[active ? 'bg-secondary-50' : '', 'block w-full text-left px-4 py-2 text-sm text-secondary-700']">
                    Configure
                  </button>
                </MenuItem>
                <MenuItem v-slot="{ active }">
                  <button :class="[active ? 'bg-danger-50' : '', 'block w-full text-left px-4 py-2 text-sm text-danger-700']">
                    Delete
                  </button>
                </MenuItem>
              </MenuItems>
            </Menu>
          </div>

          <!-- Status Badge -->
          <div class="mb-4">
            <span :class="[
              'status-indicator',
              project.status === 'active' ? 'status-success' :
              project.status === 'deploying' ? 'status-warning' :
              project.status === 'failed' ? 'status-danger' :
              'status-info'
            ]">
              {{ project.status }}
            </span>
          </div>
          
          <p class="text-sm text-secondary-600 mb-4 line-clamp-2">{{ project.description }}</p>
          
          <!-- Project Details -->
          <div class="space-y-3 mb-6">
            <div class="flex items-center justify-between text-sm">
              <span class="text-secondary-500">Last Deploy:</span>
              <span class="text-secondary-900 font-medium">{{ project.lastDeploy }}</span>
            </div>
            <div class="flex items-center justify-between text-sm">
              <span class="text-secondary-500">Environments:</span>
              <div class="flex space-x-1">
                <span 
                  v-for="env in project.environments.slice(0, 2)" 
                  :key="env" 
                  class="px-2 py-1 bg-secondary-100 text-secondary-700 rounded-md text-xs font-medium"
                >
                  {{ env }}
                </span>
                <span 
                  v-if="project.environments.length > 2" 
                  class="px-2 py-1 bg-secondary-100 text-secondary-700 rounded-md text-xs font-medium"
                >
                  +{{ project.environments.length - 2 }}
                </span>
              </div>
            </div>
            <div class="flex items-center justify-between text-sm">
              <span class="text-secondary-500">Version:</span>
              <span class="text-secondary-900 font-medium font-mono text-xs">{{ project.version }}</span>
            </div>
          </div>
          
          <!-- Action Buttons -->
          <div class="flex space-x-2">
            <button class="btn btn-primary btn-sm flex-1">
              <RocketLaunchIcon class="h-4 w-4 mr-2" />
              Deploy
            </button>
            <button class="btn btn-secondary btn-sm">
              <CogIcon class="h-4 w-4" />
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
        <FolderIcon class="h-8 w-8 text-secondary-400" />
      </div>
      <h3 class="text-lg font-medium text-secondary-900">No applications found</h3>
      <p class="mt-2 text-sm text-secondary-500">
        {{ searchQuery ? 'Try adjusting your search terms or filters.' : 'Get started by importing your first application.' }}
      </p>
      <div class="mt-6">
        <button v-if="!searchQuery" class="btn btn-primary btn-md">
          <PlusIcon class="h-4 w-4 mr-2" />
          Import Application
        </button>
        <button v-else @click="clearSearch" class="btn btn-secondary btn-md">
          Clear Search
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { Menu, MenuButton, MenuItem, MenuItems } from '@headlessui/vue'
import { 
  FolderIcon,
  PlusIcon,
  MagnifyingGlassIcon,
  FunnelIcon,
  ChevronDownIcon,
  ArrowPathIcon,
  CheckCircleIcon,
  ExclamationTriangleIcon,
  RocketLaunchIcon,
  CogIcon,
  ChartBarIcon,
  EllipsisVerticalIcon
} from '@heroicons/vue/24/outline'

const projects = ref([])
const loading = ref(true)
const error = ref(null)
const searchQuery = ref('')
const selectedFilter = ref('all')

const filters = [
  { key: 'all', label: 'All Applications' },
  { key: 'active', label: 'Active' },
  { key: 'deploying', label: 'Deploying' },
  { key: 'failed', label: 'Issues' },
  { key: 'inactive', label: 'Inactive' },
]

// Mock data for demo - replace with API call
const mockProjects = [
  {
    id: 1,
    name: 'web-frontend',
    description: 'React-based frontend application with modern UI components and responsive design.',
    runtime: 'nodejs',
    status: 'active',
    lastDeploy: '2 hours ago',
    version: 'v1.2.3',
    environments: ['production', 'staging', 'development']
  },
  {
    id: 2,
    name: 'api-gateway',
    description: 'Central API gateway handling authentication, rate limiting, and service routing.',
    runtime: 'golang',
    status: 'deploying',
    lastDeploy: '5 minutes ago',
    version: 'v2.1.0',
    environments: ['staging', 'development']
  },
  {
    id: 3,
    name: 'auth-service',
    description: 'Microservice handling user authentication and authorization with JWT tokens.',
    runtime: 'python',
    status: 'failed',
    lastDeploy: '1 day ago',
    version: 'v1.5.2',
    environments: ['development']
  },
  {
    id: 4,
    name: 'data-processor',
    description: 'Background job processor for handling data transformation and analytics.',
    runtime: 'python',
    status: 'active',
    lastDeploy: '3 days ago',
    version: 'v0.8.1',
    environments: ['production', 'staging']
  },
  {
    id: 5,
    name: 'notification-service',
    description: 'Service responsible for sending emails, SMS, and push notifications.',
    runtime: 'nodejs',
    status: 'active',
    lastDeploy: '1 week ago',
    version: 'v1.0.5',
    environments: ['production']
  },
  {
    id: 6,
    name: 'file-storage',
    description: 'Distributed file storage service with CDN integration and automatic scaling.',
    runtime: 'golang',
    status: 'inactive',
    lastDeploy: '2 weeks ago',
    version: 'v0.3.0',
    environments: ['development']
  }
]

const filteredProjects = computed(() => {
  let filtered = projects.value

  // Filter by search query
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    filtered = filtered.filter(project => 
      project.name.toLowerCase().includes(query) ||
      project.description.toLowerCase().includes(query) ||
      project.runtime.toLowerCase().includes(query)
    )
  }

  // Filter by status
  if (selectedFilter.value !== 'all') {
    filtered = filtered.filter(project => project.status === selectedFilter.value)
  }

  return filtered
})

const activeApps = computed(() => projects.value.filter(p => p.status === 'active').length)
const deployingApps = computed(() => projects.value.filter(p => p.status === 'deploying').length)
const failedApps = computed(() => projects.value.filter(p => p.status === 'failed').length)

const fetchProjects = async () => {
  try {
    loading.value = true
    error.value = null
    
    // Use the real API
    const response = await fetch('http://localhost:8080/api/v1/projects')
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }
    const data = await response.json()
    
    // Transform API data to match UI expectations
    projects.value = data.map(project => ({
      id: project.id,
      name: project.name,
      description: project.description,
      runtime: project.runtime,
      status: project.status,
      lastDeploy: project.last_deploy,
      version: 'v1.0.0', // Default version since API doesn't provide it
      environments: project.environments
    }))
  } catch (err) {
    error.value = err.message
    console.error('Failed to fetch projects:', err)
    
    // Fallback to mock data if API fails
    projects.value = mockProjects
  } finally {
    loading.value = false
  }
}

const clearSearch = () => {
  searchQuery.value = ''
  selectedFilter.value = 'all'
}

const getProjectIcon = (runtime) => {
  const icons = {
    nodejs: 'JS',
    python: 'PY',
    golang: 'GO',
    java: 'JV',
    rust: 'RS',
    php: 'PHP'
  }
  return icons[runtime] || runtime.substring(0, 2).toUpperCase()
}

const getProjectTypeColor = (runtime) => {
  const colors = {
    nodejs: 'bg-success-100 text-success-700',
    python: 'bg-primary-100 text-primary-700',
    golang: 'bg-warning-100 text-warning-700',
    java: 'bg-danger-100 text-danger-700',
    rust: 'bg-secondary-100 text-secondary-700',
    php: 'bg-purple-100 text-purple-700'
  }
  return colors[runtime] || 'bg-secondary-100 text-secondary-700'
}

onMounted(() => {
  fetchProjects()
})
</script>
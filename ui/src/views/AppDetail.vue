<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Header -->
    <div class="bg-white shadow">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex items-center justify-between h-16">
          <div class="flex items-center">
            <button 
              @click="$router.go(-1)"
              class="mr-4 p-2 rounded-md text-gray-400 hover:text-gray-500 hover:bg-gray-100"
            >
              <ArrowLeftIcon class="h-5 w-5" />
            </button>
            <div class="flex items-center space-x-3">
              <div 
                class="w-10 h-10 rounded-lg flex items-center justify-center text-white text-sm font-medium"
                :class="getProjectTypeColor(app?.runtime || 'unknown')"
              >
                {{ getProjectIcon(app?.runtime || 'unknown') }}
              </div>
              <div>
                <h1 class="text-2xl font-bold text-gray-900">{{ app?.name || 'Loading...' }}</h1>
                <p class="text-sm text-gray-500">{{ app?.description }}</p>
              </div>
            </div>
          </div>
          <div class="flex items-center space-x-4">
            <span 
              class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
              :class="getStatusColor(app?.status || 'unknown')"
            >
              <div class="w-1.5 h-1.5 rounded-full mr-1" :class="getStatusDotColor(app?.status || 'unknown')"></div>
              {{ app?.status || 'Unknown' }}
            </span>
          </div>
        </div>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="flex items-center justify-center h-64">
      <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-indigo-600"></div>
    </div>

    <!-- Error State -->
    <div v-else-if="error" class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <div class="bg-red-50 border border-red-200 rounded-md p-4">
        <div class="flex">
          <ExclamationTriangleIcon class="h-5 w-5 text-red-400" />
          <div class="ml-3">
            <h3 class="text-sm font-medium text-red-800">Error loading application</h3>
            <div class="mt-2 text-sm text-red-700">{{ error }}</div>
          </div>
        </div>
      </div>
    </div>

    <!-- Main Content -->
    <div v-else class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
        
        <!-- Main Info -->
        <div class="lg:col-span-2 space-y-6">
          
          <!-- Quick Stats -->
          <div class="bg-white rounded-lg shadow p-6">
            <h2 class="text-lg font-medium text-gray-900 mb-4">Quick Overview</h2>
            <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
              <div class="text-center">
                <div class="text-2xl font-bold text-indigo-600">{{ app?.environments?.length || 0 }}</div>
                <div class="text-sm text-gray-500">Environments</div>
              </div>
              <div class="text-center">
                <div class="text-2xl font-bold text-green-600">{{ deployments.filter(d => d.status === 'live').length }}</div>
                <div class="text-sm text-gray-500">Live Deployments</div>
              </div>
              <div class="text-center">
                <div class="text-2xl font-bold text-red-600">{{ deployments.filter(d => d.status === 'failed').length }}</div>
                <div class="text-sm text-gray-500">Failed</div>
              </div>
              <div class="text-center">
                <div class="text-2xl font-bold text-gray-600">{{ app?.runtime || 'Unknown' }}</div>
                <div class="text-sm text-gray-500">Runtime</div>
              </div>
            </div>
          </div>

          <!-- Deployments by Environment -->
          <div class="bg-white rounded-lg shadow p-6">
            <h2 class="text-lg font-medium text-gray-900 mb-4">Deployments</h2>
            <div v-if="deployments && deployments.length > 0" class="space-y-4">
              <div 
                v-for="deployment in deployments" 
                :key="deployment.id"
                class="border border-gray-200 rounded-lg p-4"
              >
                <div class="flex items-center justify-between mb-3">
                  <div class="flex items-center space-x-3">
                    <div 
                      class="w-8 h-8 rounded-full flex items-center justify-center text-xs font-medium text-white"
                      :class="getEnvironmentColor(deployment.environment)"
                    >
                      {{ deployment.environment.charAt(0).toUpperCase() }}
                    </div>
                    <div>
                      <div class="font-medium text-gray-900">{{ deployment.environment }}</div>
                      <div class="text-sm text-gray-500">Version {{ deployment.version }}</div>
                    </div>
                  </div>
                  <div class="flex items-center space-x-4">
                    <span 
                      class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
                      :class="getStatusColor(deployment.status)"
                    >
                      <div class="w-1.5 h-1.5 rounded-full mr-1" :class="getStatusDotColor(deployment.status)"></div>
                      {{ deployment.status }}
                    </span>
                    <div class="text-sm text-gray-500">{{ deployment.deployed_at }}</div>
                  </div>
                </div>

                <!-- Replica Information -->
                <div class="mb-3 p-3 bg-gray-50 rounded-lg">
                  <div class="flex items-center justify-between mb-3">
                    <span class="text-sm font-medium text-gray-700">Replicas</span>
                    <div class="flex items-center space-x-2">
                      <span class="text-sm text-gray-600">{{ deployment.ready_replicas || 0 }}/{{ deployment.desired_replicas || 0 }}</span>
                      <div 
                        class="w-3 h-3 rounded-full animate-pulse"
                        :class="getReplicaStatusColor(deployment)"
                      ></div>
                    </div>
                  </div>
                  
                  <!-- Replica Status Grid -->
                  <div class="grid grid-cols-5 gap-1 mb-3" v-if="deployment.desired_replicas > 0">
                    <div 
                      v-for="i in Math.max(deployment.desired_replicas || 0, deployment.ready_replicas || 0)"
                      :key="i"
                      class="h-6 rounded flex items-center justify-center text-xs font-medium transition-all duration-300"
                      :class="getReplicaBoxColor(i, deployment)"
                      :title="getReplicaTooltip(i, deployment)"
                    >
                      {{ i }}
                    </div>
                  </div>
                  
                  <!-- Replica Details -->
                  <div class="text-xs text-gray-500 space-y-1">
                    <div class="flex justify-between">
                      <span>Running:</span>
                      <span class="font-medium text-green-600">{{ deployment.ready_replicas || 0 }}</span>
                    </div>
                    <div class="flex justify-between">
                      <span>Pending:</span>
                      <span class="font-medium text-yellow-600">{{ Math.max(0, (deployment.desired_replicas || 0) - (deployment.ready_replicas || 0)) }}</span>
                    </div>
                    <div class="flex justify-between">
                      <span>Age:</span>
                      <span class="font-medium">{{ getDeploymentAge(deployment.deployed_at) }}</span>
                    </div>
                  </div>
                  
                  <!-- Replica Controls -->
                  <div class="flex items-center space-x-2 mt-2">
                    <button 
                      @click="scaleApp(deployment.environment, Math.max(0, (deployment.desired_replicas || 1) - 1))"
                      :disabled="(deployment.desired_replicas || 0) <= 0 || scaling[deployment.environment]"
                      class="px-2 py-1 text-xs font-medium text-gray-700 bg-white border border-gray-300 rounded hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed transition-all duration-200"
                    >
                      <MinusIcon class="h-3 w-3" />
                    </button>
                    
                    <div class="relative">
                      <input 
                        type="number" 
                        :value="deployment.desired_replicas || 0"
                        @input="scaleApp(deployment.environment, parseInt($event.target.value))"
                        @keypress.enter="$event.preventDefault()"
                        :disabled="scaling[deployment.environment]"
                        min="0" 
                        max="10"
                        class="w-16 px-2 py-1 text-xs text-center border border-gray-300 rounded focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 disabled:opacity-50 transition-all duration-200"
                      />
                      <div 
                        v-if="scaling[deployment.environment]"
                        class="absolute inset-0 flex items-center justify-center"
                      >
                        <div class="animate-spin rounded-full h-3 w-3 border border-indigo-500 border-t-transparent"></div>
                      </div>
                    </div>
                    
                    <button 
                      @click="scaleApp(deployment.environment, (deployment.desired_replicas || 0) + 1)"
                      :disabled="(deployment.desired_replicas || 0) >= 10 || scaling[deployment.environment]"
                      class="px-2 py-1 text-xs font-medium text-gray-700 bg-white border border-gray-300 rounded hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed transition-all duration-200"
                    >
                      <PlusIcon class="h-3 w-3" />
                    </button>
                    
                    <span class="text-xs text-gray-500 ml-2">
                      {{ scaling[deployment.environment] ? 'Scaling...' : 'replicas' }}
                    </span>
                  </div>
                </div>

                <!-- Management Controls -->
                <div class="flex items-center space-x-2 mb-3">
                  <button 
                    @click="startApp(deployment.environment)"
                    :disabled="deployment.status === 'live'"
                    class="flex items-center px-3 py-1 text-xs font-medium text-green-700 bg-green-100 rounded hover:bg-green-200 disabled:opacity-50 disabled:cursor-not-allowed"
                  >
                    <PlayIcon class="h-3 w-3 mr-1" />
                    Start
                  </button>
                  
                  <button 
                    @click="stopApp(deployment.environment)"
                    :disabled="deployment.status === 'failed'"
                    class="flex items-center px-3 py-1 text-xs font-medium text-red-700 bg-red-100 rounded hover:bg-red-200 disabled:opacity-50 disabled:cursor-not-allowed"
                  >
                    <StopIcon class="h-3 w-3 mr-1" />
                    Stop
                  </button>
                  
                  <button 
                    @click="restartApp(deployment.environment)"
                    class="flex items-center px-3 py-1 text-xs font-medium text-blue-700 bg-blue-100 rounded hover:bg-blue-200"
                  >
                    <ArrowPathIcon class="h-3 w-3 mr-1" />
                    Restart
                  </button>
                </div>
                
                <!-- Routes if available -->
                <div v-if="deployment.routes && deployment.routes.length > 0" class="pt-3 border-t border-gray-100">
                  <div class="space-y-2">
                    <div class="flex items-center space-x-2 mb-2">
                      <LinkIcon class="h-4 w-4 text-gray-400" />
                      <span class="text-sm font-medium text-gray-700">Routes ({{ deployment.routes.length }})</span>
                    </div>
                    <div class="space-y-1">
                      <div 
                        v-for="route in deployment.routes" 
                        :key="`${route.host}${route.path}`"
                        class="flex items-center justify-between p-2 bg-gray-50 rounded-md"
                      >
                        <div class="flex items-center space-x-2">
                          <a 
                            :href="route.url" 
                            target="_blank"
                            class="text-sm text-indigo-600 hover:text-indigo-800 font-medium"
                          >
                            {{ route.host }}{{ route.path }}
                          </a>
                          <ArrowTopRightOnSquareIcon class="h-3 w-3 text-gray-400" />
                        </div>
                        <div class="flex items-center space-x-2 text-xs text-gray-500">
                          <span>{{ route.service_name }}:{{ route.service_port }}</span>
                          <span class="px-1.5 py-0.5 bg-gray-200 rounded text-xs">{{ route.path_type }}</span>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
                
                <!-- Legacy URL fallback for deployments without routes -->
                <div v-else-if="deployment.url" class="pt-3 border-t border-gray-100">
                  <div class="flex items-center space-x-2">
                    <LinkIcon class="h-4 w-4 text-gray-400" />
                    <a 
                      :href="deployment.url" 
                      target="_blank"
                      class="text-sm text-indigo-600 hover:text-indigo-800 font-medium"
                    >
                      {{ deployment.url }}
                    </a>
                    <ArrowTopRightOnSquareIcon class="h-3 w-3 text-gray-400" />
                  </div>
                </div>
              </div>
            </div>
            <div v-else class="text-center py-8 text-gray-500">
              <div class="text-lg font-medium mb-2">No deployments found</div>
              <div class="text-sm">Deploy this application to see deployments here.</div>
            </div>
          </div>

          <!-- Recent Activity -->
          <div class="bg-white rounded-lg shadow p-6">
            <h2 class="text-lg font-medium text-gray-900 mb-4">Recent Activity</h2>
            <div v-if="deployments && deployments.length > 0" class="space-y-4">
              <div 
                v-for="deployment in deployments.slice(0, 5)" 
                :key="`activity-${deployment.id}`"
                class="flex items-start space-x-3"
              >
                <div class="flex-shrink-0">
                  <div 
                    class="w-2 h-2 rounded-full mt-2"
                    :class="deployment.status === 'live' ? 'bg-green-400' : deployment.status === 'failed' ? 'bg-red-400' : 'bg-yellow-400'"
                  ></div>
                </div>
                <div class="flex-1 min-w-0">
                  <p class="text-sm text-gray-900">
                    <span class="font-medium">{{ deployment.environment }}</span>
                    deployment {{ deployment.status === 'live' ? 'succeeded' : deployment.status }}
                  </p>
                  <p class="text-sm text-gray-500">{{ deployment.deployed_at }}</p>
                </div>
              </div>
            </div>
            <div v-else class="text-center py-8 text-gray-500">
              <div class="text-sm">No recent activity</div>
            </div>
          </div>
        </div>

        <!-- Sidebar -->
        <div class="space-y-6">
          
          <!-- Application Info -->
          <div class="bg-white rounded-lg shadow p-6">
            <h2 class="text-lg font-medium text-gray-900 mb-4">Application Info</h2>
            <div class="space-y-3">
              <div>
                <dt class="text-sm font-medium text-gray-500">Runtime</dt>
                <dd class="mt-1 text-sm text-gray-900">{{ app?.runtime || 'Unknown' }}</dd>
              </div>
              <div>
                <dt class="text-sm font-medium text-gray-500">Repository</dt>
                <dd class="mt-1">
                  <a 
                    :href="app?.repository" 
                    target="_blank"
                    class="text-sm text-indigo-600 hover:text-indigo-800 flex items-center space-x-1"
                  >
                    <span>{{ getRepoDisplayName(app?.repository) }}</span>
                    <ArrowTopRightOnSquareIcon class="h-3 w-3" />
                  </a>
                </dd>
              </div>
              <div>
                <dt class="text-sm font-medium text-gray-500">Last Deploy</dt>
                <dd class="mt-1 text-sm text-gray-900">{{ app?.last_deploy || 'Unknown' }}</dd>
              </div>
              <div>
                <dt class="text-sm font-medium text-gray-500">Environments</dt>
                <dd class="mt-1">
                  <div class="flex flex-wrap gap-1">
                    <span 
                      v-for="env in app?.environments" 
                      :key="env"
                      class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
                      :class="getEnvironmentBadgeColor(env)"
                    >
                      {{ env }}
                    </span>
                  </div>
                </dd>
              </div>
            </div>
          </div>

          <!-- Quick Actions -->
          <div class="bg-white rounded-lg shadow p-6">
            <h2 class="text-lg font-medium text-gray-900 mb-4">Quick Actions</h2>
            <div class="space-y-3">
              <button 
                @click="deployLatest"
                :disabled="deploying"
                class="w-full flex items-center justify-center px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 disabled:opacity-50 disabled:cursor-not-allowed"
              >
                <div v-if="deploying" class="animate-spin rounded-full h-4 w-4 border border-white border-t-transparent mr-2"></div>
                <RocketLaunchIcon v-else class="h-4 w-4 mr-2" />
                {{ deploying ? 'Deploying...' : 'Deploy Latest' }}
              </button>
              <button 
                @click="viewLogs"
                class="w-full flex items-center justify-center px-4 py-2 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
              >
                <DocumentTextIcon class="h-4 w-4 mr-2" />
                View Logs
              </button>
              <button 
                @click="openSettings"
                class="w-full flex items-center justify-center px-4 py-2 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
              >
                <Cog6ToothIcon class="h-4 w-4 mr-2" />
                Settings
              </button>
            </div>
          </div>

          <!-- Environment Health -->
          <div class="bg-white rounded-lg shadow p-6">
            <h2 class="text-lg font-medium text-gray-900 mb-4">Environment Health</h2>
            <div class="space-y-3">
              <div 
                v-for="env in app?.environments" 
                :key="`health-${env}`"
                class="flex items-center justify-between"
              >
                <div class="flex items-center space-x-2">
                  <div 
                    class="w-3 h-3 rounded-full"
                    :class="getEnvironmentHealthColor(env)"
                  ></div>
                  <span class="text-sm font-medium text-gray-900">{{ env }}</span>
                </div>
                <span class="text-sm text-gray-500">
                  {{ getEnvironmentHealth(env) }}
                </span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRoute } from 'vue-router'
import {
  ArrowLeftIcon,
  ExclamationTriangleIcon,
  LinkIcon,
  ArrowTopRightOnSquareIcon,
  RocketLaunchIcon,
  DocumentTextIcon,
  Cog6ToothIcon,
  PlayIcon,
  StopIcon,
  ArrowPathIcon,
  PlusIcon,
  MinusIcon
} from '@heroicons/vue/24/outline'

const route = useRoute()
const appName = route.params.name

const app = ref(null)
const deployments = ref([])
const loading = ref(true)
const error = ref(null)
const scaling = ref({}) // Track scaling state per environment
const deploying = ref(false) // Track deployment state

// Debounce timer for scaling
let scaleTimeout = null

const fetchAppData = async () => {
  try {
    loading.value = true
    error.value = null

    // Fetch projects to find the specific app
    const projectsResponse = await fetch('http://localhost:8080/api/v1/projects')
    if (!projectsResponse.ok) {
      throw new Error(`Failed to fetch projects: ${projectsResponse.status}`)
    }
    const projects = await projectsResponse.json()
    
    // Find the specific app
    const foundApp = projects.find(p => p.name === appName)
    if (!foundApp) {
      throw new Error(`Application "${appName}" not found`)
    }
    app.value = foundApp

    // Fetch deployments for this app
    const deploymentsResponse = await fetch('http://localhost:8080/api/v1/deployments')
    if (!deploymentsResponse.ok) {
      throw new Error(`Failed to fetch deployments: ${deploymentsResponse.status}`)
    }
    const allDeployments = await deploymentsResponse.json()
    
    // Filter deployments for this app and enhance with Kubernetes data
    const appDeployments = allDeployments.filter(d => d.application === appName)
    
    // Enhance each deployment with real Kubernetes status and routes
    const enhancedDeployments = await Promise.all(appDeployments.map(async (deployment) => {
      try {
        // Fetch Kubernetes deployment status which includes routes
        const statusResponse = await fetch(`http://localhost:8080/api/v1/manage/${deployment.environment}/${appName}/status`)
        if (statusResponse.ok) {
          const statusData = await statusResponse.json()
          
          // Merge the deployment data with Kubernetes status
          return {
            ...deployment,
            desired_replicas: statusData.desired_replicas,
            ready_replicas: statusData.ready_replicas,
            available_replicas: statusData.available_replicas,
            pods: statusData.pods || [],
            routes: statusData.routes || [],
            // Set URL from routes if available
            url: statusData.routes && statusData.routes.length > 0 ? statusData.routes[0].url : deployment.url
          }
        }
      } catch (err) {
        console.warn(`Failed to fetch Kubernetes status for ${appName} in ${deployment.environment}:`, err)
      }
      
      // Return original deployment if Kubernetes status fetch fails
      return {
        ...deployment,
        routes: [],
        pods: []
      }
    }))
    
    deployments.value = enhancedDeployments

  } catch (err) {
    error.value = err.message
    console.error('Failed to fetch app data:', err)
  } finally {
    loading.value = false
  }
}

const getProjectIcon = (runtime) => {
  const icons = {
    nodejs: 'JS',
    python: 'PY',
    go: 'GO',
    java: 'JV',
    rust: 'RS',
    php: 'PHP'
  }
  return icons[runtime] || runtime.substring(0, 2).toUpperCase()
}

const getProjectTypeColor = (runtime) => {
  const colors = {
    nodejs: 'bg-green-500',
    python: 'bg-blue-500',
    go: 'bg-cyan-500',
    java: 'bg-red-500',
    rust: 'bg-orange-500',
    php: 'bg-purple-500'
  }
  return colors[runtime] || 'bg-gray-500'
}

const getStatusColor = (status) => {
  if (!status) return 'bg-gray-100 text-gray-800'
  
  const colors = {
    active: 'bg-green-100 text-green-800',
    live: 'bg-green-100 text-green-800',
    inactive: 'bg-gray-100 text-gray-800',
    building: 'bg-yellow-100 text-yellow-800',
    failed: 'bg-red-100 text-red-800'
  }
  return colors[status] || 'bg-gray-100 text-gray-800'
}

const getStatusDotColor = (status) => {
  if (!status) return 'bg-gray-400'
  
  const colors = {
    active: 'bg-green-400',
    live: 'bg-green-400',
    inactive: 'bg-gray-400',
    building: 'bg-yellow-400',
    failed: 'bg-red-400'
  }
  return colors[status] || 'bg-gray-400'
}

const getEnvironmentColor = (env) => {
  if (!env) return 'bg-gray-500'
  
  const colors = {
    dev: 'bg-blue-500',
    development: 'bg-blue-500',
    staging: 'bg-yellow-500',
    production: 'bg-green-500',
    prod: 'bg-green-500'
  }
  return colors[env] || 'bg-gray-500'
}

const getEnvironmentBadgeColor = (env) => {
  if (!env) return 'bg-gray-100 text-gray-800'
  
  const colors = {
    dev: 'bg-blue-100 text-blue-800',
    development: 'bg-blue-100 text-blue-800',
    staging: 'bg-yellow-100 text-yellow-800',
    production: 'bg-green-100 text-green-800',
    prod: 'bg-green-100 text-green-800'
  }
  return colors[env] || 'bg-gray-100 text-gray-800'
}

const getEnvironmentHealthColor = (env) => {
  if (!env || !deployments.value) return 'bg-gray-400'
  
  const deployment = deployments.value.find(d => d && d.environment === env)
  if (!deployment) return 'bg-gray-400'
  
  return deployment.status === 'live' ? 'bg-green-400' : 
         deployment.status === 'failed' ? 'bg-red-400' : 'bg-yellow-400'
}

const getEnvironmentHealth = (env) => {
  if (!env || !deployments.value) return 'Unknown'
  
  const deployment = deployments.value.find(d => d && d.environment === env)
  if (!deployment) return 'Unknown'
  
  return deployment.status === 'live' ? 'Healthy' : 
         deployment.status === 'failed' ? 'Failed' : 'Building'
}

const getRepoDisplayName = (url) => {
  if (!url) return 'Unknown'
  const parts = url.split('/')
  return parts[parts.length - 1].replace('.git', '')
}

// Replica status helpers
const getReplicaStatusColor = (deployment) => {
  const ready = deployment.ready_replicas || 0
  const desired = deployment.desired_replicas || 0
  
  if (desired === 0) return 'bg-gray-400'
  if (ready === desired) return 'bg-green-400'
  if (ready === 0) return 'bg-red-400'
  return 'bg-yellow-400'
}

const getReplicaBoxColor = (index, deployment) => {
  const ready = deployment.ready_replicas || 0
  const desired = deployment.desired_replicas || 0
  
  if (index <= ready) {
    return 'bg-green-100 text-green-800 border border-green-200'
  } else if (index <= desired) {
    return 'bg-yellow-100 text-yellow-800 border border-yellow-200 animate-pulse'
  } else {
    return 'bg-gray-100 text-gray-400 border border-gray-200'
  }
}

const getReplicaTooltip = (index, deployment) => {
  const ready = deployment.ready_replicas || 0
  const desired = deployment.desired_replicas || 0
  
  if (index <= ready) {
    return `Replica ${index}: Running`
  } else if (index <= desired) {
    return `Replica ${index}: Starting...`
  } else {
    return `Replica ${index}: Terminated`
  }
}

const getDeploymentAge = (deployedAt) => {
  if (!deployedAt) return 'Unknown'
  
  const now = new Date()
  const deployed = new Date(deployedAt)
  const diffMs = now - deployed
  const diffMins = Math.floor(diffMs / 60000)
  const diffHours = Math.floor(diffMins / 60)
  const diffDays = Math.floor(diffHours / 24)
  
  if (diffDays > 0) return `${diffDays}d ${diffHours % 24}h`
  if (diffHours > 0) return `${diffHours}h ${diffMins % 60}m`
  if (diffMins > 0) return `${diffMins}m`
  return 'Just now'
}

// Management functions with debouncing and optimistic updates
const scaleApp = async (environment, replicas) => {
  // Clear existing timeout
  if (scaleTimeout) {
    clearTimeout(scaleTimeout)
  }
  
  // Set scaling state
  scaling.value[environment] = true
  
  // Optimistically update the UI immediately
  const deployment = deployments.value.find(d => d.environment === environment)
  if (deployment) {
    deployment.desired_replicas = parseInt(replicas)
  }
  
  // Debounce the API call
  scaleTimeout = setTimeout(async () => {
    try {
      const response = await fetch(`http://localhost:8080/api/v1/apps/${appName}/scale?env=${environment}`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ replicas: parseInt(replicas) })
      })
      
      if (!response.ok) {
        throw new Error(`Failed to scale application: ${response.status}`)
      }
      
      // Refresh data after a short delay to see the actual state
      setTimeout(() => {
        fetchAppData()
      }, 1500)
      
    } catch (err) {
      console.error('Failed to scale application:', err)
      error.value = err.message
      // Revert optimistic update on error
      fetchAppData()
    } finally {
      scaling.value[environment] = false
    }
  }, 300) // 300ms debounce
}

const startApp = async (environment) => {
  try {
    const response = await fetch(`http://localhost:8080/api/v1/apps/${appName}/start?env=${environment}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ replicas: 1 })
    })
    
    if (!response.ok) {
      throw new Error(`Failed to start application: ${response.status}`)
    }
    
    // Refresh data after starting
    setTimeout(() => {
      fetchAppData()
    }, 1000)
    
  } catch (err) {
    console.error('Failed to start application:', err)
    error.value = err.message
  }
}

const stopApp = async (environment) => {
  try {
    const response = await fetch(`http://localhost:8080/api/v1/apps/${appName}/stop?env=${environment}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      }
    })
    
    if (!response.ok) {
      throw new Error(`Failed to stop application: ${response.status}`)
    }
    
    // Refresh data after stopping
    setTimeout(() => {
      fetchAppData()
    }, 1000)
    
  } catch (err) {
    console.error('Failed to stop application:', err)
    error.value = err.message
  }
}

const restartApp = async (environment) => {
  try {
    const response = await fetch(`http://localhost:8080/api/v1/apps/${appName}/restart?env=${environment}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      }
    })
    
    if (!response.ok) {
      throw new Error(`Failed to restart application: ${response.status}`)
    }
    
    // Refresh data after restarting
    setTimeout(() => {
      fetchAppData()
    }, 2000) // Wait a bit longer for restart
    
  } catch (err) {
    console.error('Failed to restart application:', err)
    error.value = err.message
  }
}

// Quick Actions functions
const deployLatest = async () => {
  try {
    deploying.value = true
    
    // Deploy to all environments
    for (const env of app.value?.environments || []) {
      const response = await fetch('http://localhost:8080/api/v1/deploy', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          project_id: app.value.id,
          environment: env,
          version: 'latest'
        })
      })
      
      if (!response.ok) {
        throw new Error(`Failed to deploy to ${env}: ${response.status}`)
      }
    }
    
    // Refresh data after deployment
    setTimeout(() => {
      fetchAppData()
    }, 2000)
    
  } catch (err) {
    console.error('Failed to deploy:', err)
    error.value = err.message
  } finally {
    deploying.value = false
  }
}

const viewLogs = () => {
  // Open logs in a new window/tab
  const logsWindow = window.open('', '_blank', 'width=1000,height=700')
  if (logsWindow) {
    logsWindow.document.title = `${app.value?.name} Logs`
    logsWindow.document.body.innerHTML = `
      <div style="font-family: 'Monaco', 'Menlo', 'Consolas', monospace; padding: 20px; background: #1a1a1a; color: #00ff00; height: 100vh; overflow-y: auto; font-size: 14px; line-height: 1.4;">
        <div style="color: #ffffff; margin-bottom: 20px; border-bottom: 1px solid #333; padding-bottom: 10px;">
          <h2 style="margin: 0; font-size: 18px;">${app.value?.name} - Application Logs</h2>
          <p style="margin: 5px 0 0 0; color: #888; font-size: 12px;">Real-time logs from all environments</p>
        </div>
        <div id="logs-content">
          <div style="margin-bottom: 8px; color: #888;">[${new Date().toISOString()}] <span style="color: #00aaff;">INFO</span> Fetching application logs...</div>
          <div style="margin-bottom: 8px; color: #888;">[${new Date().toISOString()}] <span style="color: #00aaff;">INFO</span> Application: <span style="color: #ffffff;">${app.value?.name}</span></div>
          <div style="margin-bottom: 8px; color: #888;">[${new Date().toISOString()}] <span style="color: #00aaff;">INFO</span> Runtime: <span style="color: #ffffff;">${app.value?.runtime}</span></div>
          <div style="margin-bottom: 8px; color: #888;">[${new Date().toISOString()}] <span style="color: #00aaff;">INFO</span> Status: <span style="color: #00ff00;">${app.value?.status}</span></div>
          <div style="margin-bottom: 8px; color: #888;">[${new Date().toISOString()}] <span style="color: #00aaff;">INFO</span> Environments: <span style="color: #ffffff;">${app.value?.environments?.join(', ')}</span></div>
          <div style="margin-bottom: 8px; color: #888;">[${new Date().toISOString()}] <span style="color: #00aaff;">INFO</span> Last Deploy: <span style="color: #ffffff;">${app.value?.last_deploy}</span></div>
          <div style="margin-bottom: 8px; color: #888;">[${new Date().toISOString()}] <span style="color: #ffaa00;">WARN</span> Live log streaming will be implemented in next version</div>
          <div style="margin-bottom: 8px; color: #888;">[${new Date().toISOString()}] <span style="color: #ff6600;">TODO</span> Connect to Kubernetes pod logs API for real-time streaming</div>
          <div style="margin-bottom: 8px; color: #888;">[${new Date().toISOString()}] <span style="color: #00aaff;">INFO</span> Sample log entries:</div>
          <div style="margin: 10px 0; padding: 10px; background: #222; border-left: 3px solid #00ff00;">
            <div style="margin-bottom: 5px;">[2025-09-22T21:30:01.123Z] <span style="color: #00aaff;">INFO</span> Server starting on port 3000</div>
            <div style="margin-bottom: 5px;">[2025-09-22T21:30:01.456Z] <span style="color: #00aaff;">INFO</span> Database connection established</div>
            <div style="margin-bottom: 5px;">[2025-09-22T21:30:02.789Z] <span style="color: #00aaff;">INFO</span> Routes initialized</div>
            <div style="margin-bottom: 5px;">[2025-09-22T21:30:03.012Z] <span style="color: #00ff00;">SUCCESS</span> Application ready to accept connections</div>
            <div style="margin-bottom: 5px;">[2025-09-22T21:35:15.234Z] <span style="color: #00aaff;">INFO</span> GET /api/health - 200 OK (5ms)</div>
            <div style="margin-bottom: 5px;">[2025-09-22T21:35:16.567Z] <span style="color: #00aaff;">INFO</span> GET /api/status - 200 OK (12ms)</div>
          </div>
          <div style="margin-bottom: 8px; color: #888;">[${new Date().toISOString()}] <span style="color: #00aaff;">INFO</span> End of available logs</div>
        </div>
        <div style="position: fixed; bottom: 20px; right: 20px;">
          <button onclick="window.close()" style="padding: 8px 16px; background: #333; color: white; border: 1px solid #555; border-radius: 4px; cursor: pointer; font-family: inherit;">Close</button>
        </div>
      </div>
    `
  }
}

const openSettings = () => {
  // Create a simple settings modal
  const settingsModal = `
    <div style="position: fixed; top: 0; left: 0; right: 0; bottom: 0; background: rgba(0,0,0,0.8); display: flex; align-items: center; justify-content: center; z-index: 10000;">
      <div style="background: white; border-radius: 8px; padding: 24px; max-width: 500px; width: 90%; max-height: 80%; overflow-y: auto;">
        <h2 style="margin: 0 0 20px 0; color: #333;">Settings for ${app.value?.name}</h2>
        <div style="margin-bottom: 16px;">
          <h3 style="margin: 0 0 8px 0; color: #555; font-size: 14px;">Environment Configuration</h3>
          <p style="margin: 0; color: #666; font-size: 12px;">Configure deployment settings for each environment</p>
        </div>
        <div style="margin-bottom: 16px;">
          <h3 style="margin: 0 0 8px 0; color: #555; font-size: 14px;">Resource Limits</h3>
          <p style="margin: 0; color: #666; font-size: 12px;">CPU and memory allocation settings</p>
        </div>
        <div style="margin-bottom: 16px;">
          <h3 style="margin: 0 0 8px 0; color: #555; font-size: 14px;">Auto-scaling Rules</h3>
          <p style="margin: 0; color: #666; font-size: 12px;">Horizontal pod autoscaler configuration</p>
        </div>
        <div style="margin-bottom: 16px;">
          <h3 style="margin: 0 0 8px 0; color: #555; font-size: 14px;">Monitoring & Alerts</h3>
          <p style="margin: 0; color: #666; font-size: 12px;">Health checks and notification settings</p>
        </div>
        <div style="margin-bottom: 24px;">
          <h3 style="margin: 0 0 8px 0; color: #555; font-size: 14px;">Git Integration</h3>
          <p style="margin: 0; color: #666; font-size: 12px;">Repository and deployment webhook settings</p>
        </div>
        <div style="text-align: center; padding-top: 16px; border-top: 1px solid #eee;">
          <p style="margin: 0 0 16px 0; color: #888; font-size: 14px;">Full settings UI coming soon!</p>
          <button onclick="this.closest('div[style*=fixed]').remove()" style="padding: 8px 24px; background: #4f46e5; color: white; border: none; border-radius: 4px; cursor: pointer; font-size: 14px;">Close</button>
        </div>
      </div>
    </div>
  `
  
  const modalDiv = document.createElement('div')
  modalDiv.innerHTML = settingsModal
  document.body.appendChild(modalDiv.firstElementChild)
}

onMounted(() => {
  fetchAppData()
})
</script>
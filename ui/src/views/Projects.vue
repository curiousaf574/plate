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
          <!-- Search -->
          <div class="relative">
            <MagnifyingGlassIcon class="absolute left-3 top-1/2 transform -translate-y-1/2 h-4 w-4 text-secondary-400" />
            <input 
              v-model="searchQuery"
              type="text" 
              placeholder="Search applications..." 
              class="form-input pl-10 w-64"
            />
          </div>
          
          <!-- Sort Dropdown -->
          <Menu as="div" class="relative">
            <MenuButton class="btn btn-secondary btn-sm">
              <ArrowsUpDownIcon class="h-4 w-4 mr-2" />
              Sort
              <ChevronDownIcon class="h-4 w-4 ml-1" />
            </MenuButton>
            <MenuItems class="absolute right-0 z-10 mt-2 w-56 origin-top-right rounded-xl bg-white py-2 shadow-large ring-1 ring-secondary-900/5 focus:outline-none">
              <MenuItem v-for="sort in sortOptions" :key="sort.key" v-slot="{ active }">
                <button 
                  @click="selectedSort = sort.key"
                  :class="[
                    active ? 'bg-secondary-50' : '', 
                    selectedSort === sort.key ? 'text-primary-700 font-medium' : 'text-secondary-700',
                    'flex items-center w-full text-left px-4 py-2 text-sm'
                  ]"
                >
                  <component :is="sort.icon" class="h-4 w-4 mr-3" />
                  {{ sort.label }}
                </button>
              </MenuItem>
            </MenuItems>
          </Menu>
          
          <!-- Filter Dropdown -->
          <Menu as="div" class="relative">
            <MenuButton class="btn btn-secondary btn-sm">
              <FunnelIcon class="h-4 w-4 mr-2" />
              {{ getFilterLabel() }}
              <ChevronDownIcon class="h-4 w-4 ml-1" />
            </MenuButton>
            <MenuItems class="absolute right-0 z-10 mt-2 w-48 origin-top-right rounded-xl bg-white py-2 shadow-large ring-1 ring-secondary-900/5 focus:outline-none">
              <div class="px-4 py-2 text-xs font-medium text-secondary-500 uppercase tracking-wide">Status</div>
              <MenuItem v-for="filter in statusFilters" :key="filter.key" v-slot="{ active }">
                <button 
                  @click="selectedStatusFilter = filter.key"
                  :class="[
                    active ? 'bg-secondary-50' : '', 
                    selectedStatusFilter === filter.key ? 'text-primary-700 font-medium' : 'text-secondary-700',
                    'flex items-center w-full text-left px-4 py-2 text-sm'
                  ]"
                >
                  <div :class="['w-2 h-2 rounded-full mr-3', filter.color]"></div>
                  {{ filter.label }}
                </button>
              </MenuItem>
              <div class="border-t border-secondary-100 my-1"></div>
              <div class="px-4 py-2 text-xs font-medium text-secondary-500 uppercase tracking-wide">Runtime</div>
              <MenuItem v-for="runtime in runtimeFilters" :key="runtime" v-slot="{ active }">
                <button 
                  @click="selectedRuntimeFilter = runtime"
                  :class="[
                    active ? 'bg-secondary-50' : '', 
                    selectedRuntimeFilter === runtime ? 'text-primary-700 font-medium' : 'text-secondary-700',
                    'block w-full text-left px-4 py-2 text-sm'
                  ]"
                >
                  {{ runtime || 'All Runtimes' }}
                </button>
              </MenuItem>
            </MenuItems>
          </Menu>
          
          <!-- View Toggle -->
          <div class="flex rounded-lg border border-secondary-200 bg-secondary-50 p-1">
            <button 
              @click="viewMode = 'tile'"
              :class="[
                'p-1.5 rounded-md transition-all duration-200',
                viewMode === 'tile' ? 'bg-white shadow-sm text-primary-600' : 'text-secondary-400 hover:text-secondary-600'
              ]"
            >
              <Squares2X2Icon class="h-4 w-4" />
            </button>
            <button 
              @click="viewMode = 'list'"
              :class="[
                'p-1.5 rounded-md transition-all duration-200',
                viewMode === 'list' ? 'bg-white shadow-sm text-primary-600' : 'text-secondary-400 hover:text-secondary-600'
              ]"
            >
              <Bars3Icon class="h-4 w-4" />
            </button>
          </div>
          
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
            <p class="text-2xl font-bold text-secondary-900">{{ filteredAndSortedProjects.length }}</p>
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

    <!-- Applications View -->
    <div v-else-if="filteredAndSortedProjects.length > 0">
      <!-- Tile View -->
      <div v-if="viewMode === 'tile'" class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-6">
        <div 
          v-for="(project, index) in filteredAndSortedProjects" 
          :key="project.id" 
          class="card-hover slide-up cursor-pointer"
          :style="{ animationDelay: `${index * 50}ms` }"
          @click="$router.push(`/projects/${project.name}`)"
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
              <MenuButton 
                class="p-1 hover:bg-secondary-100 rounded-lg transition-colors"
                @click.stop
              >
                <EllipsisVerticalIcon class="h-4 w-4 text-secondary-400" />
              </MenuButton>
              <MenuItems class="absolute right-0 z-10 mt-2 w-48 origin-top-right rounded-xl bg-white py-2 shadow-large ring-1 ring-secondary-900/5 focus:outline-none">
                <MenuItem v-slot="{ active }">
                  <button 
                    :class="[active ? 'bg-secondary-50' : '', 'block w-full text-left px-4 py-2 text-sm text-secondary-700']"
                    @click="$router.push(`/projects/${project.name}`)"
                  >
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
            <button 
              class="btn btn-primary btn-sm flex-1"
              @click.stop
            >
              <RocketLaunchIcon class="h-4 w-4 mr-2" />
              Deploy
            </button>
            <button 
              class="btn btn-secondary btn-sm"
              @click.stop="$router.push(`/projects/${project.name}`)"
            >
              <CogIcon class="h-4 w-4" />
            </button>
            <button 
              class="btn btn-secondary btn-sm"
              @click.stop="$router.push(`/projects/${project.name}`)"
            >
              <ChartBarIcon class="h-4 w-4" />
            </button>
          </div>
        </div>
        </div>
      </div>
      
      <!-- List View -->
      <div v-else class="bg-white rounded-xl border border-secondary-200 overflow-hidden">
        <!-- Table Header -->
        <div class="px-6 py-4 border-b border-secondary-200 bg-secondary-50">
          <div class="grid grid-cols-12 gap-4 items-center">
            <div class="col-span-4">
              <span class="text-sm font-medium text-secondary-700">Application</span>
            </div>
            <div class="col-span-1 text-center">
              <span class="text-sm font-medium text-secondary-700">Status</span>
            </div>
            <div class="col-span-2 text-center">
              <span class="text-sm font-medium text-secondary-700">Environments</span>
            </div>
            <div class="col-span-1 text-center">
              <span class="text-sm font-medium text-secondary-700">Routes</span>
            </div>
            <div class="col-span-2 text-center">
              <span class="text-sm font-medium text-secondary-700">Last Deploy</span>
            </div>
            <div class="col-span-2 text-center">
              <span class="text-sm font-medium text-secondary-700">Actions</span>
            </div>
          </div>
        </div>

        <!-- Table Body -->
        <div class="divide-y divide-secondary-100">
          <div 
            v-for="(project, index) in filteredAndSortedProjects" 
            :key="project.id" 
            class="px-6 py-5 hover:bg-secondary-50 cursor-pointer transition-all duration-200 slide-up group"
            :style="{ animationDelay: `${index * 25}ms` }"
            @click="$router.push(`/projects/${project.name}`)"
          >
            <div class="grid grid-cols-12 gap-4 items-center">
              <!-- Application Info -->
              <div class="col-span-4">
                <div class="flex items-center space-x-4">
                  <div :class="[
                    'h-11 w-11 rounded-xl flex items-center justify-center text-sm font-semibold shadow-sm transition-transform group-hover:scale-105',
                    getProjectTypeColor(project.runtime)
                  ]">
                    {{ getProjectIcon(project.runtime) }}
                  </div>
                  <div class="min-w-0 flex-1">
                    <div class="flex items-center space-x-2">
                      <h4 class="font-semibold text-secondary-900 truncate">{{ project.name }}</h4>
                      <span class="text-xs font-mono text-secondary-500 bg-secondary-100 px-2 py-1 rounded-md">
                        {{ project.version }}
                      </span>
                    </div>
                    <p class="text-sm text-secondary-600 truncate">{{ project.description }}</p>
                    <div class="flex items-center space-x-2 mt-1">
                      <span class="text-xs text-secondary-500">{{ project.runtime }}</span>
                      <span class="text-xs text-secondary-400">•</span>
                      <span class="text-xs text-secondary-500">Updated {{ project.lastDeploy }}</span>
                    </div>
                  </div>
                </div>
              </div>
              
              <!-- Status -->
              <div class="col-span-1 text-center">
                <div class="flex flex-col items-center space-y-1">
                  <span :class="[
                    'inline-flex items-center px-2.5 py-1 rounded-full text-xs font-medium',
                    project.status === 'active' ? 'bg-success-100 text-success-800' :
                    project.status === 'deploying' ? 'bg-warning-100 text-warning-800' :
                    project.status === 'failed' ? 'bg-danger-100 text-danger-800' :
                    'bg-secondary-100 text-secondary-800'
                  ]">
                    <div :class="[
                      'w-1.5 h-1.5 rounded-full mr-1.5',
                      project.status === 'active' ? 'bg-success-400' :
                      project.status === 'deploying' ? 'bg-warning-400 animate-pulse' :
                      project.status === 'failed' ? 'bg-danger-400' :
                      'bg-secondary-400'
                    ]"></div>
                    {{ project.status }}
                  </span>
                </div>
              </div>
              
              <!-- Environments -->
              <div class="col-span-2 text-center">
                <div class="flex flex-wrap justify-center gap-1">
                  <span 
                    v-for="env in project.environments.slice(0, 3)" 
                    :key="env" 
                    :class="[
                      'inline-flex items-center px-2 py-1 rounded-md text-xs font-medium',
                      env === 'production' ? 'bg-danger-100 text-danger-700' :
                      env === 'staging' ? 'bg-warning-100 text-warning-700' :
                      'bg-primary-100 text-primary-700'
                    ]"
                  >
                    {{ env }}
                  </span>
                  <span 
                    v-if="project.environments.length > 3" 
                    class="inline-flex items-center px-2 py-1 rounded-md text-xs font-medium bg-secondary-100 text-secondary-700"
                  >
                    +{{ project.environments.length - 3 }}
                  </span>
                </div>
              </div>
              
              <!-- Routes -->
              <div class="col-span-1 text-center">
                <div class="flex flex-col items-center space-y-1">
                  <div v-if="project.routes && project.routes.length > 0" class="flex items-center justify-center space-x-1">
                    <GlobeAltIcon class="h-4 w-4 text-success-500" />
                    <span class="text-sm font-semibold text-success-700">{{ project.routes.length }}</span>
                  </div>
                  <div v-else class="flex items-center justify-center space-x-1">
                    <GlobeAltIcon class="h-4 w-4 text-secondary-300" />
                    <span class="text-sm text-secondary-400">0</span>
                  </div>
                  <span class="text-xs text-secondary-500">
                    {{ project.routes && project.routes.length ? 'active' : 'none' }}
                  </span>
                </div>
              </div>
              
              <!-- Last Deploy -->
              <div class="col-span-2 text-center">
                <div class="flex flex-col items-center space-y-1">
                  <div class="flex items-center space-x-2">
                    <ClockIcon class="h-4 w-4 text-secondary-400" />
                    <span class="text-sm font-medium text-secondary-900">{{ project.lastDeploy }}</span>
                  </div>
                  <div class="flex items-center space-x-1 text-xs text-secondary-500">
                    <span>via</span>
                    <span class="font-mono bg-secondary-100 px-1.5 py-0.5 rounded">{{ project.runtime }}</span>
                  </div>
                </div>
              </div>
              
              <!-- Actions -->
              <div class="col-span-2 text-center">
                <div class="flex items-center justify-center space-x-2">
                  <button 
                    class="btn btn-primary btn-sm opacity-0 group-hover:opacity-100 transition-opacity duration-200"
                    @click.stop
                    title="Deploy Application"
                  >
                    <RocketLaunchIcon class="h-4 w-4 mr-1" />
                    Deploy
                  </button>
                  
                  <button 
                    class="btn btn-secondary btn-sm opacity-0 group-hover:opacity-100 transition-opacity duration-200"
                    @click.stop="$router.push(`/projects/${project.name}`)"
                    title="View Details"
                  >
                    <ChartBarIcon class="h-4 w-4" />
                  </button>
                  
                  <Menu as="div" class="relative">
                    <MenuButton 
                      class="btn btn-secondary btn-sm opacity-0 group-hover:opacity-100 transition-opacity duration-200"
                      @click.stop
                      title="More Actions"
                    >
                      <EllipsisVerticalIcon class="h-4 w-4" />
                    </MenuButton>
                    <MenuItems class="absolute right-0 z-10 mt-2 w-56 origin-top-right rounded-xl bg-white py-2 shadow-large ring-1 ring-secondary-900/5 focus:outline-none">
                      <MenuItem v-slot="{ active }">
                        <button 
                          :class="[active ? 'bg-secondary-50' : '', 'flex items-center w-full text-left px-4 py-2 text-sm text-secondary-700']"
                          @click="$router.push(`/projects/${project.name}`)"
                        >
                          <ChartBarIcon class="h-4 w-4 mr-3 text-secondary-400" />
                          View Details
                        </button>
                      </MenuItem>
                      <MenuItem v-slot="{ active }">
                        <button :class="[active ? 'bg-secondary-50' : '', 'flex items-center w-full text-left px-4 py-2 text-sm text-secondary-700']">
                          <CogIcon class="h-4 w-4 mr-3 text-secondary-400" />
                          Configure
                        </button>
                      </MenuItem>
                      <MenuItem v-slot="{ active }">
                        <button :class="[active ? 'bg-secondary-50' : '', 'flex items-center w-full text-left px-4 py-2 text-sm text-secondary-700']">
                          <ArrowPathIcon class="h-4 w-4 mr-3 text-secondary-400" />
                          Redeploy
                        </button>
                      </MenuItem>
                      <MenuItem v-slot="{ active }">
                        <button :class="[active ? 'bg-secondary-50' : '', 'flex items-center w-full text-left px-4 py-2 text-sm text-secondary-700']">
                          <GlobeAltIcon class="h-4 w-4 mr-3 text-secondary-400" />
                          View Routes
                        </button>
                      </MenuItem>
                      <div class="border-t border-secondary-100 my-1"></div>
                      <MenuItem v-slot="{ active }">
                        <button :class="[active ? 'bg-danger-50' : '', 'flex items-center w-full text-left px-4 py-2 text-sm text-danger-700']">
                          <ExclamationTriangleIcon class="h-4 w-4 mr-3 text-danger-400" />
                          Delete Application
                        </button>
                      </MenuItem>
                    </MenuItems>
                  </Menu>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Footer with summary -->
        <div class="px-6 py-4 border-t border-secondary-200 bg-secondary-50">
          <div class="flex items-center justify-between text-sm text-secondary-600">
            <span>Showing {{ filteredAndSortedProjects.length }} application{{ filteredAndSortedProjects.length !== 1 ? 's' : '' }}</span>
            <div class="flex items-center space-x-4">
              <span class="flex items-center space-x-1">
                <div class="w-2 h-2 rounded-full bg-success-400"></div>
                <span>{{ activeApps }} Active</span>
              </span>
              <span class="flex items-center space-x-1">
                <div class="w-2 h-2 rounded-full bg-warning-400"></div>
                <span>{{ deployingApps }} Deploying</span>
              </span>
              <span class="flex items-center space-x-1">
                <div class="w-2 h-2 rounded-full bg-danger-400"></div>
                <span>{{ failedApps }} Issues</span>
              </span>
            </div>
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
  EllipsisVerticalIcon,
  ArrowsUpDownIcon,
  Squares2X2Icon,
  Bars3Icon,
  GlobeAltIcon,
  ClockIcon,
  TagIcon,
  FireIcon
} from '@heroicons/vue/24/outline'

const projects = ref([])
const loading = ref(true)
const error = ref(null)
const searchQuery = ref('')
const selectedStatusFilter = ref('all')
const selectedRuntimeFilter = ref('')
const selectedSort = ref('name')
const viewMode = ref('tile')

const statusFilters = [
  { key: 'all', label: 'All Status', color: 'bg-secondary-400' },
  { key: 'active', label: 'Active', color: 'bg-success-400' },
  { key: 'deploying', label: 'Deploying', color: 'bg-warning-400' },
  { key: 'failed', label: 'Issues', color: 'bg-danger-400' },
  { key: 'inactive', label: 'Inactive', color: 'bg-secondary-400' },
]

const sortOptions = [
  { key: 'name', label: 'Name (A-Z)', icon: TagIcon },
  { key: 'name-desc', label: 'Name (Z-A)', icon: TagIcon },
  { key: 'status', label: 'Status', icon: CheckCircleIcon },
  { key: 'lastDeploy', label: 'Last Deploy', icon: ClockIcon },
  { key: 'runtime', label: 'Runtime', icon: CogIcon },
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

const runtimeFilters = computed(() => {
  const runtimes = [...new Set(projects.value.map(p => p.runtime))].sort()
  return ['', ...runtimes]
})

const filteredAndSortedProjects = computed(() => {
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
  if (selectedStatusFilter.value !== 'all') {
    filtered = filtered.filter(project => project.status === selectedStatusFilter.value)
  }

  // Filter by runtime
  if (selectedRuntimeFilter.value) {
    filtered = filtered.filter(project => project.runtime === selectedRuntimeFilter.value)
  }

  // Sort
  filtered.sort((a, b) => {
    switch (selectedSort.value) {
      case 'name':
        return a.name.localeCompare(b.name)
      case 'name-desc':
        return b.name.localeCompare(a.name)
      case 'status':
        return a.status.localeCompare(b.status)
      case 'lastDeploy':
        // Simple string comparison for demo - you'd parse dates in real app
        return b.lastDeploy.localeCompare(a.lastDeploy)
      case 'runtime':
        return a.runtime.localeCompare(b.runtime)
      default:
        return 0
    }
  })

  return filtered
})

const getFilterLabel = () => {
  const statusLabel = statusFilters.find(f => f.key === selectedStatusFilter.value)?.label || 'All'
  const runtimeLabel = selectedRuntimeFilter.value || 'All Runtimes'
  if (selectedStatusFilter.value === 'all' && !selectedRuntimeFilter.value) {
    return 'Filter'
  }
  return `${statusLabel}${selectedRuntimeFilter.value ? ` • ${runtimeLabel}` : ''}`
}

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
    
    // Transform API data and fetch route information for each project
    const projectsWithRoutes = await Promise.all(data.map(async (project) => {
      let routes = []
      
      // Fetch route information for each environment
      try {
        // Try to get deployment status which now includes routes
        const statusResponse = await fetch(`http://localhost:8080/api/v1/manage/dev/${project.name}/status`)
        if (statusResponse.ok) {
          const statusData = await statusResponse.json()
          routes = statusData.routes || []
        }
      } catch (err) {
        console.warn(`Failed to fetch routes for ${project.name}:`, err)
      }
      
      return {
        id: project.id,
        name: project.name,
        description: project.description,
        runtime: project.runtime,
        status: project.status,
        lastDeploy: project.last_deploy,
        version: 'v1.0.0', // Default version since API doesn't provide it
        environments: project.environments,
        routes: routes
      }
    }))
    
    projects.value = projectsWithRoutes
  } catch (err) {
    error.value = err.message
    console.error('Failed to fetch projects:', err)
    
    // Fallback to mock data if API fails
    projects.value = mockProjects.map(project => ({
      ...project,
      routes: [] // Add empty routes to mock data
    }))
  } finally {
    loading.value = false
  }
}

const clearSearch = () => {
  searchQuery.value = ''
  selectedStatusFilter.value = 'all'
  selectedRuntimeFilter.value = ''
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
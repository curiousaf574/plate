<template>
  <!-- Off-canvas menu for mobile -->
  <div v-if="open" class="relative z-50 lg:hidden">
    <div class="fixed inset-0 bg-secondary-900/80 backdrop-blur-sm" @click="$emit('close')" />
    <div class="fixed inset-0 flex">
      <div class="relative mr-16 flex w-full max-w-xs flex-1">
        <div class="absolute left-full top-0 flex w-16 justify-center pt-5">
          <button type="button" class="-m-2.5 p-2.5 hover:bg-white/20 rounded-lg transition-colors" @click="$emit('close')">
            <span class="sr-only">Close sidebar</span>
            <XMarkIcon class="h-6 w-6 text-white" aria-hidden="true" />
          </button>
        </div>
        <div class="flex grow flex-col gap-y-5 overflow-y-auto bg-white px-6 pb-4 shadow-xl">
          <div class="flex h-18 shrink-0 items-center">
            <div class="flex items-center">
              <div class="h-8 w-8 bg-gradient-to-br from-primary-500 to-primary-600 rounded-lg flex items-center justify-center shadow-sm">
                <span class="text-white font-bold text-sm">P</span>
              </div>
              <h2 class="ml-3 text-xl font-bold text-secondary-900">Plate</h2>
            </div>
          </div>
          <nav class="flex flex-1 flex-col">
            <ul role="list" class="flex flex-1 flex-col gap-y-7">
              <li>
                <ul role="list" class="-mx-2 space-y-1">
                  <li v-for="item in navigation" :key="item.name">
                    <router-link
                      :to="item.href"
                      :class="[
                        $route.name === item.current
                          ? 'nav-link-active'
                          : 'nav-link-inactive'
                      ]"
                    >
                      <component
                        :is="item.icon"
                        :class="[
                          $route.name === item.current ? 'text-primary-600' : 'text-secondary-400 group-hover:text-primary-600',
                          'h-5 w-5 shrink-0'
                        ]"
                        aria-hidden="true"
                      />
                      {{ item.name }}
                    </router-link>
                  </li>
                </ul>
              </li>
              
              <!-- Quick Stats -->
              <li class="mt-auto">
                <div class="rounded-xl bg-gradient-to-br from-primary-50 to-primary-100 p-4 border border-primary-200">
                  <h3 class="text-sm font-medium text-primary-900 mb-2">Quick Stats</h3>
                  <div class="space-y-2">
                    <div class="flex justify-between text-xs">
                      <span class="text-primary-700">Active Apps</span>
                      <span class="font-medium text-primary-900">12</span>
                    </div>
                    <div class="flex justify-between text-xs">
                      <span class="text-primary-700">Deployments</span>
                      <span class="font-medium text-primary-900">8</span>
                    </div>
                  </div>
                </div>
              </li>
            </ul>
          </nav>
        </div>
      </div>
    </div>
  </div>

  <!-- Static sidebar for desktop -->
  <div class="hidden lg:fixed lg:inset-y-0 lg:z-50 lg:flex lg:w-72 lg:flex-col">
    <div class="flex grow flex-col gap-y-5 overflow-y-auto border-r border-secondary-200/60 bg-white px-6 pb-4 shadow-soft">
      <div class="flex h-18 shrink-0 items-center">
        <div class="flex items-center">
          <div class="h-10 w-10 bg-gradient-to-br from-primary-500 to-primary-600 rounded-xl flex items-center justify-center shadow-sm">
            <span class="text-white font-bold text-lg">P</span>
          </div>
          <div class="ml-3">
            <h2 class="text-xl font-bold text-secondary-900">Plate Platform</h2>
            <p class="text-xs text-secondary-500">Ready, Set & Deploy!</p>
          </div>
        </div>
      </div>
      
      <nav class="flex flex-1 flex-col">
        <ul role="list" class="flex flex-1 flex-col gap-y-7">
          <li>
            <ul role="list" class="-mx-2 space-y-1">
              <li v-for="item in navigation" :key="item.name">
                <router-link
                  :to="item.href"
                  :class="[
                    $route.name === item.current
                      ? 'nav-link-active'
                      : 'nav-link-inactive'
                  ]"
                >
                  <component
                    :is="item.icon"
                    :class="[
                      $route.name === item.current ? 'text-primary-600' : 'text-secondary-400 group-hover:text-primary-600',
                      'h-5 w-5 shrink-0 transition-colors'
                    ]"
                    aria-hidden="true"
                  />
                  {{ item.name }}
                  <span v-if="item.badge" :class="[
                    'ml-auto px-2 py-0.5 text-xs font-medium rounded-full',
                    $route.name === item.current 
                      ? 'bg-primary-100 text-primary-700' 
                      : 'bg-secondary-100 text-secondary-600 group-hover:bg-primary-100 group-hover:text-primary-700'
                  ]">
                    {{ item.badge }}
                  </span>
                </router-link>
              </li>
            </ul>
          </li>
          
          <!-- Quick Actions -->
          <li>
            <div class="text-xs font-semibold leading-6 text-secondary-400 mb-2">Quick Actions</div>
            <ul role="list" class="-mx-2 space-y-1">
              <li>
                <button class="nav-link-inactive w-full text-left group">
                  <PlusIcon class="h-5 w-5 shrink-0 text-secondary-400 group-hover:text-primary-600 transition-colors" />
                  New Deployment
                </button>
              </li>
              <li>
                <button class="nav-link-inactive w-full text-left group">
                  <CloudArrowUpIcon class="h-5 w-5 shrink-0 text-secondary-400 group-hover:text-primary-600 transition-colors" />
                  Import Project
                </button>
              </li>
            </ul>
          </li>
          
          <!-- Quick Stats -->
          <li class="mt-auto">
            <div class="rounded-xl bg-gradient-to-br from-primary-50 to-primary-100 p-4 border border-primary-200">
              <h3 class="text-sm font-medium text-primary-900 mb-3">System Health</h3>
              <div class="space-y-3">
                <div class="flex items-center justify-between">
                  <div class="flex items-center">
                    <div class="h-2 w-2 bg-success-500 rounded-full mr-2"></div>
                    <span class="text-xs text-primary-700">All Systems</span>
                  </div>
                  <span class="text-xs font-medium text-success-700">Healthy</span>
                </div>
                <div class="space-y-2">
                  <div class="flex justify-between text-xs">
                    <span class="text-primary-700">CPU Usage</span>
                    <span class="font-medium text-primary-900">45%</span>
                  </div>
                  <div class="w-full bg-primary-200 rounded-full h-1">
                    <div class="bg-primary-600 h-1 rounded-full" style="width: 45%"></div>
                  </div>
                </div>
              </div>
            </div>
          </li>
        </ul>
      </nav>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import {
  HomeIcon,
  FolderIcon,
  RocketLaunchIcon,
  CogIcon,
  XMarkIcon,
  PlusIcon,
  CloudArrowUpIcon,
} from '@heroicons/vue/24/outline'

defineProps({
  open: Boolean
})

defineEmits(['close'])

const route = useRoute()

const navigation = computed(() => [
  { name: 'Dashboard', href: '/', icon: HomeIcon, current: 'dashboard' },
  { name: 'Applications', href: '/projects', icon: FolderIcon, current: 'projects', badge: '12' },
  { name: 'Deployments', href: '/deployments', icon: RocketLaunchIcon, current: 'deployments', badge: '3' },
  { name: 'Settings', href: '/settings', icon: CogIcon, current: 'settings' },
])
</script>
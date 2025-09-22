<template>
  <div class="min-h-screen bg-secondary-50">
    <Sidebar :open="sidebarOpen" @close="sidebarOpen = false" />
    
    <div class="lg:pl-72">
      <!-- Top navigation -->
      <div class="sticky top-0 z-40 flex h-18 shrink-0 items-center gap-x-4 border-b border-white/10 bg-white/80 backdrop-blur-md px-4 shadow-soft sm:gap-x-6 sm:px-6 lg:px-8">
        <button type="button" class="-m-2.5 p-2.5 text-secondary-600 hover:text-secondary-900 transition-colors lg:hidden" @click="sidebarOpen = true">
          <span class="sr-only">Open sidebar</span>
          <Bars3Icon class="h-6 w-6" aria-hidden="true" />
        </button>

        <div class="h-6 w-px bg-secondary-200 lg:hidden" aria-hidden="true" />

        <div class="flex flex-1 gap-x-4 self-stretch lg:gap-x-6">
          <div class="relative flex flex-1 items-center">
            <h1 class="text-xl font-bold text-secondary-900">Plate Platform</h1>
            <div class="ml-3 px-2 py-1 bg-primary-100 text-primary-700 text-xs font-medium rounded-full">
              v1.0.0
            </div>
          </div>
          <div class="flex items-center gap-x-4 lg:gap-x-6">
            <!-- Notifications -->
            <button type="button" class="relative p-2 text-secondary-500 hover:text-secondary-700 transition-colors">
              <span class="sr-only">View notifications</span>
              <BellIcon class="h-5 w-5" aria-hidden="true" />
              <span class="absolute top-1 right-1 h-2 w-2 bg-danger-500 rounded-full"></span>
            </button>
            
            <!-- Profile dropdown -->
            <Menu as="div" class="relative">
              <MenuButton class="-m-1.5 flex items-center p-1.5 hover:bg-secondary-100 rounded-lg transition-colors">
                <span class="sr-only">Open user menu</span>
                <div class="h-8 w-8 bg-gradient-to-br from-primary-500 to-primary-600 rounded-lg flex items-center justify-center shadow-sm">
                  <UserIcon class="h-4 w-4 text-white" />
                </div>
                <div class="ml-3 hidden sm:block text-left">
                  <p class="text-sm font-medium text-secondary-900">John Doe</p>
                  <p class="text-xs text-secondary-500">Administrator</p>
                </div>
                <ChevronDownIcon class="ml-2 h-4 w-4 text-secondary-500" aria-hidden="true" />
              </MenuButton>
              <transition 
                enter-active-class="transition ease-out duration-100" 
                enter-from-class="transform opacity-0 scale-95" 
                enter-to-class="transform opacity-100 scale-100" 
                leave-active-class="transition ease-in duration-75" 
                leave-from-class="transform opacity-100 scale-100" 
                leave-to-class="transform opacity-0 scale-95"
              >
                <MenuItems class="absolute right-0 z-10 mt-2.5 w-56 origin-top-right rounded-xl bg-white py-2 shadow-large ring-1 ring-secondary-900/5 focus:outline-none">
                  <MenuItem v-slot="{ active }">
                    <a href="#" :class="[active ? 'bg-secondary-50' : '', 'flex items-center px-4 py-2 text-sm text-secondary-700']">
                      <UserIcon class="mr-3 h-4 w-4 text-secondary-400" />
                      Your Profile
                    </a>
                  </MenuItem>
                  <MenuItem v-slot="{ active }">
                    <a href="#" :class="[active ? 'bg-secondary-50' : '', 'flex items-center px-4 py-2 text-sm text-secondary-700']">
                      <CogIcon class="mr-3 h-4 w-4 text-secondary-400" />
                      Settings
                    </a>
                  </MenuItem>
                  <MenuItem v-slot="{ active }">
                    <a href="#" :class="[active ? 'bg-secondary-50' : '', 'flex items-center px-4 py-2 text-sm text-secondary-700']">
                      <QuestionMarkCircleIcon class="mr-3 h-4 w-4 text-secondary-400" />
                      Support
                    </a>
                  </MenuItem>
                  <div class="my-1 h-px bg-secondary-200" />
                  <MenuItem v-slot="{ active }">
                    <a href="#" :class="[active ? 'bg-danger-50' : '', 'flex items-center px-4 py-2 text-sm text-danger-700']">
                      <ArrowRightOnRectangleIcon class="mr-3 h-4 w-4 text-danger-400" />
                      Sign out
                    </a>
                  </MenuItem>
                </MenuItems>
              </transition>
            </Menu>
          </div>
        </div>
      </div>

      <!-- Main content -->
      <main class="py-8">
        <div class="px-4 sm:px-6 lg:px-8">
          <slot />
        </div>
      </main>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { Menu, MenuButton, MenuItem, MenuItems } from '@headlessui/vue'
import { 
  Bars3Icon, 
  UserIcon, 
  ChevronDownIcon, 
  BellIcon,
  CogIcon,
  QuestionMarkCircleIcon,
  ArrowRightOnRectangleIcon
} from '@heroicons/vue/24/outline'
import Sidebar from './Sidebar.vue'

const sidebarOpen = ref(false)
</script>
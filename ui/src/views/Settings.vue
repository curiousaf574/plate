<template>
  <div>
    <div class="mb-8">
      <h1 class="text-2xl font-bold text-gray-900">Settings</h1>
      <p class="mt-2 text-gray-600">Configure your development platform preferences</p>
    </div>

    <div class="space-y-8">
      <!-- API Configuration -->
      <div class="card">
        <h2 class="text-lg font-medium text-gray-900 mb-4">API Configuration</h2>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700">API Base URL</label>
            <input v-model="settings.apiUrl" type="text" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 sm:text-sm" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">Authentication Token</label>
            <input v-model="settings.authToken" type="password" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 sm:text-sm" />
          </div>
        </div>
      </div>

      <!-- Deployment Settings -->
      <div class="card">
        <h2 class="text-lg font-medium text-gray-900 mb-4">Deployment Settings</h2>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700">Default Environment</label>
            <input v-model="settings.clusterName" type="text" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 sm:text-sm" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">Default Region</label>
            <input v-model="settings.namespace" type="text" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 sm:text-sm" />
          </div>
        </div>
      </div>

      <!-- Build & Deploy Configuration -->
      <div class="card">
        <h2 class="text-lg font-medium text-gray-900 mb-4">Build & Deploy</h2>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700">Build Timeout (minutes)</label>
            <input v-model="settings.argoCdUrl" type="text" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 sm:text-sm" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">Auto-deploy on push</label>
            <input v-model="settings.argoCdToken" type="checkbox" class="mt-1 h-4 w-4 text-primary-600 focus:ring-primary-500 border-gray-300 rounded" />
          </div>
        </div>
      </div>

      <!-- Repository Configuration -->
      <div class="card">
        <h2 class="text-lg font-medium text-gray-900 mb-4">Repository Settings</h2>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700">Default Branch</label>
            <input v-model="settings.giteaUrl" type="text" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 sm:text-sm" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">Personal Access Token</label>
            <input v-model="settings.giteaToken" type="password" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 sm:text-sm" />
          </div>
        </div>
      </div>

      <!-- Notification Settings -->
      <div class="card">
        <h2 class="text-lg font-medium text-gray-900 mb-4">Notifications</h2>
        <div class="space-y-4">
          <div class="flex items-center">
            <input v-model="settings.emailNotifications" type="checkbox" class="h-4 w-4 text-primary-600 focus:ring-primary-500 border-gray-300 rounded" />
            <label class="ml-2 block text-sm text-gray-900">Email me when deployments complete</label>
          </div>
          <div class="flex items-center">
            <input v-model="settings.slackNotifications" type="checkbox" class="h-4 w-4 text-primary-600 focus:ring-primary-500 border-gray-300 rounded" />
            <label class="ml-2 block text-sm text-gray-900">Send alerts for failed deployments</label>
          </div>
        </div>
      </div>

      <!-- Actions -->
      <div class="flex justify-end space-x-4">
        <button class="btn-secondary">Reset</button>
        <button class="btn-primary" @click="saveSettings">Save Settings</button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const settings = ref({
  apiUrl: 'http://localhost:8080',
  authToken: '',
  clusterName: 'development',
  namespace: 'us-west-2',
  argoCdUrl: '10',
  argoCdToken: false,
  giteaUrl: 'main',
  giteaToken: '',
  emailNotifications: true,
  slackNotifications: false
})

const saveSettings = () => {
  // Implementation would save to backend
  alert('Settings saved successfully!')
}
</script>
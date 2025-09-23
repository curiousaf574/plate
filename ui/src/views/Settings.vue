<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Header -->
    <div class="bg-white shadow-sm border-b border-gray-200">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="py-6">
          <div class="flex items-center justify-between">
            <div>
              <h1 class="text-3xl font-bold text-gray-900">Settings</h1>
              <p class="mt-2 text-gray-600">Manage your platform configuration and preferences</p>
            </div>
            <div class="flex items-center space-x-3">
              <button 
                @click="resetToDefaults"
                class="inline-flex items-center px-4 py-2 border border-gray-300 rounded-lg text-sm font-medium text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 transition-colors duration-200"
              >
                <ArrowPathIcon class="h-4 w-4 mr-2" />
                Reset to Defaults
              </button>
              <button 
                @click="saveSettings"
                :disabled="saving"
                class="inline-flex items-center px-6 py-2 border border-transparent rounded-lg text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 disabled:opacity-50 disabled:cursor-not-allowed transition-colors duration-200"
              >
                <div v-if="saving" class="animate-spin rounded-full h-4 w-4 border border-white border-t-transparent mr-2"></div>
                <CheckIcon v-else class="h-4 w-4 mr-2" />
                {{ saving ? 'Saving...' : 'Save Changes' }}
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Content -->
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <div class="lg:grid lg:grid-cols-12 lg:gap-x-8">
        
        <!-- Sidebar Navigation -->
        <div class="lg:col-span-3">
          <nav class="space-y-2 sticky top-8">
            <button
              v-for="section in sections"
              :key="section.id"
              @click="activeSection = section.id"
              :class="[
                'w-full flex items-center px-3 py-2 text-sm font-medium rounded-lg transition-colors duration-200',
                activeSection === section.id
                  ? 'bg-indigo-50 text-indigo-700 border-l-4 border-indigo-700'
                  : 'text-gray-600 hover:bg-gray-50 hover:text-gray-900'
              ]"
            >
              <component :is="section.icon" class="h-5 w-5 mr-3" />
              {{ section.title }}
            </button>
          </nav>
        </div>

        <!-- Main Content -->
        <div class="mt-8 lg:mt-0 lg:col-span-9">
          
          <!-- API Configuration -->
          <div v-show="activeSection === 'api'" class="space-y-6">
            <div class="bg-white rounded-xl shadow-sm border border-gray-200 overflow-hidden">
              <div class="px-6 py-4 bg-gradient-to-r from-blue-50 to-indigo-50 border-b border-gray-200">
                <div class="flex items-center">
                  <GlobeAltIcon class="h-6 w-6 text-indigo-600 mr-3" />
                  <div>
                    <h3 class="text-lg font-semibold text-gray-900">API Configuration</h3>
                    <p class="text-sm text-gray-600">Configure your backend API connection settings</p>
                  </div>
                </div>
              </div>
              <div class="p-6 space-y-6">
                <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                  <div>
                    <label class="block text-sm font-semibold text-gray-700 mb-2">API Base URL</label>
                    <div class="relative">
                      <input 
                        v-model="settings.apiUrl" 
                        type="text" 
                        placeholder="https://api.example.com"
                        class="block w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition-colors duration-200" 
                      />
                      <div class="absolute inset-y-0 right-0 pr-3 flex items-center">
                        <span :class="settings.apiUrl ? 'text-green-500' : 'text-gray-400'">
                          <CheckCircleIcon v-if="settings.apiUrl" class="h-5 w-5" />
                          <ExclamationCircleIcon v-else class="h-5 w-5" />
                        </span>
                      </div>
                    </div>
                    <p class="mt-1 text-xs text-gray-500">The base URL for your Plate API server</p>
                  </div>
                  <div>
                    <label class="block text-sm font-semibold text-gray-700 mb-2">Authentication Token</label>
                    <div class="relative">
                      <input 
                        v-model="settings.authToken" 
                        :type="showToken ? 'text' : 'password'"
                        placeholder="Enter your API token"
                        class="block w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition-colors duration-200" 
                      />
                      <button
                        @click="showToken = !showToken"
                        class="absolute inset-y-0 right-0 pr-3 flex items-center text-gray-400 hover:text-gray-600"
                      >
                        <EyeIcon v-if="!showToken" class="h-5 w-5" />
                        <EyeSlashIcon v-else class="h-5 w-5" />
                      </button>
                    </div>
                    <p class="mt-1 text-xs text-gray-500">Personal access token for API authentication</p>
                  </div>
                </div>
                <div class="bg-blue-50 rounded-lg p-4">
                  <div class="flex">
                    <InformationCircleIcon class="h-5 w-5 text-blue-400 mt-0.5" />
                    <div class="ml-3">
                      <h4 class="text-sm font-medium text-blue-800">Connection Status</h4>
                      <p class="mt-1 text-sm text-blue-600">
                        <span v-if="settings.apiUrl && settings.authToken" class="flex items-center">
                          <span class="inline-block w-2 h-2 bg-green-400 rounded-full mr-2"></span>
                          Configuration complete
                        </span>
                        <span v-else class="flex items-center">
                          <span class="inline-block w-2 h-2 bg-yellow-400 rounded-full mr-2"></span>
                          Please configure both URL and token
                        </span>
                      </p>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Kubernetes Configuration -->
          <div v-show="activeSection === 'kubernetes'" class="space-y-6">
            <div class="bg-white rounded-xl shadow-sm border border-gray-200 overflow-hidden">
              <div class="px-6 py-4 bg-gradient-to-r from-green-50 to-emerald-50 border-b border-gray-200">
                <div class="flex items-center">
                  <CloudIcon class="h-6 w-6 text-green-600 mr-3" />
                  <div>
                    <h3 class="text-lg font-semibold text-gray-900">Kubernetes Configuration</h3>
                    <p class="text-sm text-gray-600">Configure your Kubernetes cluster connection and deployment settings</p>
                  </div>
                </div>
              </div>
              <div class="p-6 space-y-6">
                <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                  <div>
                    <label class="block text-sm font-semibold text-gray-700 mb-2">Default Environment</label>
                    <select 
                      v-model="settings.defaultEnvironment" 
                      class="block w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition-colors duration-200"
                    >
                      <option value="dev">Development</option>
                      <option value="staging">Staging</option>
                      <option value="production">Production</option>
                    </select>
                    <p class="mt-1 text-xs text-gray-500">Default environment for new deployments</p>
                  </div>
                  <div>
                    <label class="block text-sm font-semibold text-gray-700 mb-2">Default Namespace</label>
                    <input 
                      v-model="settings.defaultNamespace" 
                      type="text"
                      placeholder="default"
                      class="block w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition-colors duration-200" 
                    />
                    <p class="mt-1 text-xs text-gray-500">Kubernetes namespace for deployments</p>
                  </div>
                </div>
                <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                  <div>
                    <label class="block text-sm font-semibold text-gray-700 mb-2">Cluster Context</label>
                    <input 
                      v-model="settings.clusterContext" 
                      type="text"
                      placeholder="docker-desktop"
                      class="block w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition-colors duration-200" 
                    />
                    <p class="mt-1 text-xs text-gray-500">kubectl context name</p>
                  </div>
                  <div>
                    <label class="block text-sm font-semibold text-gray-700 mb-2">Default Replicas</label>
                    <input 
                      v-model.number="settings.defaultReplicas" 
                      type="number"
                      min="1"
                      max="10"
                      class="block w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition-colors duration-200" 
                    />
                    <p class="mt-1 text-xs text-gray-500">Default number of replicas for new deployments</p>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Build & Deploy -->
          <div v-show="activeSection === 'build'" class="space-y-6">
            <div class="bg-white rounded-xl shadow-sm border border-gray-200 overflow-hidden">
              <div class="px-6 py-4 bg-gradient-to-r from-purple-50 to-pink-50 border-b border-gray-200">
                <div class="flex items-center">
                  <CogIcon class="h-6 w-6 text-purple-600 mr-3" />
                  <div>
                    <h3 class="text-lg font-semibold text-gray-900">Build & Deploy Configuration</h3>
                    <p class="text-sm text-gray-600">Configure build and deployment automation settings</p>
                  </div>
                </div>
              </div>
              <div class="p-6 space-y-6">
                <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                  <div>
                    <label class="block text-sm font-semibold text-gray-700 mb-2">Build Timeout</label>
                    <div class="relative">
                      <input 
                        v-model.number="settings.buildTimeout" 
                        type="number"
                        min="1"
                        max="60"
                        class="block w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition-colors duration-200" 
                      />
                      <div class="absolute inset-y-0 right-0 pr-3 flex items-center pointer-events-none">
                        <span class="text-gray-500 text-sm">minutes</span>
                      </div>
                    </div>
                    <p class="mt-1 text-xs text-gray-500">Maximum time to wait for builds to complete</p>
                  </div>
                  <div>
                    <label class="block text-sm font-semibold text-gray-700 mb-2">Deploy Timeout</label>
                    <div class="relative">
                      <input 
                        v-model.number="settings.deployTimeout" 
                        type="number"
                        min="1"
                        max="30"
                        class="block w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition-colors duration-200" 
                      />
                      <div class="absolute inset-y-0 right-0 pr-3 flex items-center pointer-events-none">
                        <span class="text-gray-500 text-sm">minutes</span>
                      </div>
                    </div>
                    <p class="mt-1 text-xs text-gray-500">Maximum time to wait for deployments to complete</p>
                  </div>
                </div>
                
                <div class="space-y-4">
                  <div class="flex items-center justify-between p-4 bg-gray-50 rounded-lg">
                    <div class="flex items-center">
                      <div class="flex-shrink-0">
                        <RocketLaunchIcon class="h-6 w-6 text-indigo-600" />
                      </div>
                      <div class="ml-3">
                        <h4 class="text-sm font-medium text-gray-900">Auto-deploy on push</h4>
                        <p class="text-sm text-gray-500">Automatically deploy when code is pushed to main branch</p>
                      </div>
                    </div>
                    <div class="flex-shrink-0">
                      <button
                        @click="settings.autoDeploy = !settings.autoDeploy"
                        :class="[
                          'relative inline-flex h-6 w-11 flex-shrink-0 cursor-pointer rounded-full border-2 border-transparent transition-colors duration-200 ease-in-out focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2',
                          settings.autoDeploy ? 'bg-indigo-600' : 'bg-gray-200'
                        ]"
                      >
                        <span
                          :class="[
                            'pointer-events-none inline-block h-5 w-5 transform rounded-full bg-white shadow ring-0 transition duration-200 ease-in-out',
                            settings.autoDeploy ? 'translate-x-5' : 'translate-x-0'
                          ]"
                        />
                      </button>
                    </div>
                  </div>

                  <div class="flex items-center justify-between p-4 bg-gray-50 rounded-lg">
                    <div class="flex items-center">
                      <div class="flex-shrink-0">
                        <ArrowPathIcon class="h-6 w-6 text-green-600" />
                      </div>
                      <div class="ml-3">
                        <h4 class="text-sm font-medium text-gray-900">Auto-restart on failure</h4>
                        <p class="text-sm text-gray-500">Automatically restart failed deployments</p>
                      </div>
                    </div>
                    <div class="flex-shrink-0">
                      <button
                        @click="settings.autoRestart = !settings.autoRestart"
                        :class="[
                          'relative inline-flex h-6 w-11 flex-shrink-0 cursor-pointer rounded-full border-2 border-transparent transition-colors duration-200 ease-in-out focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2',
                          settings.autoRestart ? 'bg-indigo-600' : 'bg-gray-200'
                        ]"
                      >
                        <span
                          :class="[
                            'pointer-events-none inline-block h-5 w-5 transform rounded-full bg-white shadow ring-0 transition duration-200 ease-in-out',
                            settings.autoRestart ? 'translate-x-5' : 'translate-x-0'
                          ]"
                        />
                      </button>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Repository Settings -->
          <div v-show="activeSection === 'repository'" class="space-y-6">
            <div class="bg-white rounded-xl shadow-sm border border-gray-200 overflow-hidden">
              <div class="px-6 py-4 bg-gradient-to-r from-orange-50 to-red-50 border-b border-gray-200">
                <div class="flex items-center">
                  <CodeBracketIcon class="h-6 w-6 text-orange-600 mr-3" />
                  <div>
                    <h3 class="text-lg font-semibold text-gray-900">Repository Configuration</h3>
                    <p class="text-sm text-gray-600">Configure Git repository integration and settings</p>
                  </div>
                </div>
              </div>
              <div class="p-6 space-y-6">
                <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                  <div>
                    <label class="block text-sm font-semibold text-gray-700 mb-2">Default Branch</label>
                    <input 
                      v-model="settings.defaultBranch" 
                      type="text"
                      placeholder="main"
                      class="block w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition-colors duration-200" 
                    />
                    <p class="mt-1 text-xs text-gray-500">Default branch for new repositories</p>
                  </div>
                  <div>
                    <label class="block text-sm font-semibold text-gray-700 mb-2">Git Provider</label>
                    <select 
                      v-model="settings.gitProvider" 
                      class="block w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition-colors duration-200"
                    >
                      <option value="github">GitHub</option>
                      <option value="gitlab">GitLab</option>
                      <option value="gitea">Gitea</option>
                      <option value="bitbucket">Bitbucket</option>
                    </select>
                    <p class="mt-1 text-xs text-gray-500">Your preferred Git provider</p>
                  </div>
                </div>
                <div>
                  <label class="block text-sm font-semibold text-gray-700 mb-2">Personal Access Token</label>
                  <div class="relative">
                    <input 
                      v-model="settings.gitToken" 
                      :type="showGitToken ? 'text' : 'password'"
                      placeholder="Enter your Git personal access token"
                      class="block w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition-colors duration-200" 
                    />
                    <button
                      @click="showGitToken = !showGitToken"
                      class="absolute inset-y-0 right-0 pr-3 flex items-center text-gray-400 hover:text-gray-600"
                    >
                      <EyeIcon v-if="!showGitToken" class="h-5 w-5" />
                      <EyeSlashIcon v-else class="h-5 w-5" />
                    </button>
                  </div>
                  <p class="mt-1 text-xs text-gray-500">Required for private repositories and advanced Git operations</p>
                </div>
              </div>
            </div>
          </div>

          <!-- Notifications -->
          <div v-show="activeSection === 'notifications'" class="space-y-6">
            <div class="bg-white rounded-xl shadow-sm border border-gray-200 overflow-hidden">
              <div class="px-6 py-4 bg-gradient-to-r from-yellow-50 to-orange-50 border-b border-gray-200">
                <div class="flex items-center">
                  <BellIcon class="h-6 w-6 text-yellow-600 mr-3" />
                  <div>
                    <h3 class="text-lg font-semibold text-gray-900">Notification Settings</h3>
                    <p class="text-sm text-gray-600">Configure how and when you receive notifications</p>
                  </div>
                </div>
              </div>
              <div class="p-6 space-y-6">
                <div class="space-y-4">
                  <div class="flex items-center justify-between p-4 bg-gray-50 rounded-lg">
                    <div class="flex items-center">
                      <div class="flex-shrink-0">
                        <EnvelopeIcon class="h-6 w-6 text-blue-600" />
                      </div>
                      <div class="ml-3">
                        <h4 class="text-sm font-medium text-gray-900">Email Notifications</h4>
                        <p class="text-sm text-gray-500">Receive email alerts for deployment events</p>
                      </div>
                    </div>
                    <div class="flex-shrink-0">
                      <button
                        @click="settings.emailNotifications = !settings.emailNotifications"
                        :class="[
                          'relative inline-flex h-6 w-11 flex-shrink-0 cursor-pointer rounded-full border-2 border-transparent transition-colors duration-200 ease-in-out focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2',
                          settings.emailNotifications ? 'bg-indigo-600' : 'bg-gray-200'
                        ]"
                      >
                        <span
                          :class="[
                            'pointer-events-none inline-block h-5 w-5 transform rounded-full bg-white shadow ring-0 transition duration-200 ease-in-out',
                            settings.emailNotifications ? 'translate-x-5' : 'translate-x-0'
                          ]"
                        />
                      </button>
                    </div>
                  </div>

                  <div class="flex items-center justify-between p-4 bg-gray-50 rounded-lg">
                    <div class="flex items-center">
                      <div class="flex-shrink-0">
                        <ChatBubbleLeftRightIcon class="h-6 w-6 text-green-600" />
                      </div>
                      <div class="ml-3">
                        <h4 class="text-sm font-medium text-gray-900">Slack Notifications</h4>
                        <p class="text-sm text-gray-500">Send alerts to Slack channels</p>
                      </div>
                    </div>
                    <div class="flex-shrink-0">
                      <button
                        @click="settings.slackNotifications = !settings.slackNotifications"
                        :class="[
                          'relative inline-flex h-6 w-11 flex-shrink-0 cursor-pointer rounded-full border-2 border-transparent transition-colors duration-200 ease-in-out focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2',
                          settings.slackNotifications ? 'bg-indigo-600' : 'bg-gray-200'
                        ]"
                      >
                        <span
                          :class="[
                            'pointer-events-none inline-block h-5 w-5 transform rounded-full bg-white shadow ring-0 transition duration-200 ease-in-out',
                            settings.slackNotifications ? 'translate-x-5' : 'translate-x-0'
                          ]"
                        />
                      </button>
                    </div>
                  </div>

                  <div class="flex items-center justify-between p-4 bg-gray-50 rounded-lg">
                    <div class="flex items-center">
                      <div class="flex-shrink-0">
                        <ExclamationTriangleIcon class="h-6 w-6 text-red-600" />
                      </div>
                      <div class="ml-3">
                        <h4 class="text-sm font-medium text-gray-900">Failure Alerts</h4>
                        <p class="text-sm text-gray-500">Get immediate alerts for deployment failures</p>
                      </div>
                    </div>
                    <div class="flex-shrink-0">
                      <button
                        @click="settings.failureAlerts = !settings.failureAlerts"
                        :class="[
                          'relative inline-flex h-6 w-11 flex-shrink-0 cursor-pointer rounded-full border-2 border-transparent transition-colors duration-200 ease-in-out focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2',
                          settings.failureAlerts ? 'bg-indigo-600' : 'bg-gray-200'
                        ]"
                      >
                        <span
                          :class="[
                            'pointer-events-none inline-block h-5 w-5 transform rounded-full bg-white shadow ring-0 transition duration-200 ease-in-out',
                            settings.failureAlerts ? 'translate-x-5' : 'translate-x-0'
                          ]"
                        />
                      </button>
                    </div>
                  </div>
                </div>

                <div v-if="settings.emailNotifications" class="mt-6">
                  <label class="block text-sm font-semibold text-gray-700 mb-2">Email Address</label>
                  <input 
                    v-model="settings.emailAddress" 
                    type="email"
                    placeholder="your@email.com"
                    class="block w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition-colors duration-200" 
                  />
                </div>

                <div v-if="settings.slackNotifications" class="mt-6">
                  <label class="block text-sm font-semibold text-gray-700 mb-2">Slack Webhook URL</label>
                  <input 
                    v-model="settings.slackWebhook" 
                    type="url"
                    placeholder="https://hooks.slack.com/services/..."
                    class="block w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition-colors duration-200" 
                  />
                </div>
              </div>
            </div>
          </div>

          <!-- Security -->
          <div v-show="activeSection === 'security'" class="space-y-6">
            <div class="bg-white rounded-xl shadow-sm border border-gray-200 overflow-hidden">
              <div class="px-6 py-4 bg-gradient-to-r from-red-50 to-pink-50 border-b border-gray-200">
                <div class="flex items-center">
                  <ShieldCheckIcon class="h-6 w-6 text-red-600 mr-3" />
                  <div>
                    <h3 class="text-lg font-semibold text-gray-900">Security Settings</h3>
                    <p class="text-sm text-gray-600">Configure security and access control settings</p>
                  </div>
                </div>
              </div>
              <div class="p-6 space-y-6">
                <div class="space-y-4">
                  <div class="flex items-center justify-between p-4 bg-gray-50 rounded-lg">
                    <div class="flex items-center">
                      <div class="flex-shrink-0">
                        <LockClosedIcon class="h-6 w-6 text-indigo-600" />
                      </div>
                      <div class="ml-3">
                        <h4 class="text-sm font-medium text-gray-900">Two-Factor Authentication</h4>
                        <p class="text-sm text-gray-500">Add an extra layer of security to your account</p>
                      </div>
                    </div>
                    <div class="flex-shrink-0">
                      <button
                        @click="settings.twoFactor = !settings.twoFactor"
                        :class="[
                          'relative inline-flex h-6 w-11 flex-shrink-0 cursor-pointer rounded-full border-2 border-transparent transition-colors duration-200 ease-in-out focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2',
                          settings.twoFactor ? 'bg-indigo-600' : 'bg-gray-200'
                        ]"
                      >
                        <span
                          :class="[
                            'pointer-events-none inline-block h-5 w-5 transform rounded-full bg-white shadow ring-0 transition duration-200 ease-in-out',
                            settings.twoFactor ? 'translate-x-5' : 'translate-x-0'
                          ]"
                        />
                      </button>
                    </div>
                  </div>

                  <div class="flex items-center justify-between p-4 bg-gray-50 rounded-lg">
                    <div class="flex items-center">
                      <div class="flex-shrink-0">
                        <KeyIcon class="h-6 w-6 text-yellow-600" />
                      </div>
                      <div class="ml-3">
                        <h4 class="text-sm font-medium text-gray-900">API Key Rotation</h4>
                        <p class="text-sm text-gray-500">Automatically rotate API keys monthly</p>
                      </div>
                    </div>
                    <div class="flex-shrink-0">
                      <button
                        @click="settings.keyRotation = !settings.keyRotation"
                        :class="[
                          'relative inline-flex h-6 w-11 flex-shrink-0 cursor-pointer rounded-full border-2 border-transparent transition-colors duration-200 ease-in-out focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2',
                          settings.keyRotation ? 'bg-indigo-600' : 'bg-gray-200'
                        ]"
                      >
                        <span
                          :class="[
                            'pointer-events-none inline-block h-5 w-5 transform rounded-full bg-white shadow ring-0 transition duration-200 ease-in-out',
                            settings.keyRotation ? 'translate-x-5' : 'translate-x-0'
                          ]"
                        />
                      </button>
                    </div>
                  </div>
                </div>

                <div class="bg-red-50 rounded-lg p-4">
                  <div class="flex">
                    <ExclamationTriangleIcon class="h-5 w-5 text-red-400 mt-0.5" />
                    <div class="ml-3">
                      <h4 class="text-sm font-medium text-red-800">Danger Zone</h4>
                      <div class="mt-2">
                        <button 
                          @click="showDeleteConfirm = true"
                          class="text-sm bg-red-600 text-white px-4 py-2 rounded-md hover:bg-red-700 transition-colors duration-200"
                        >
                          Reset All Settings
                        </button>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>

        </div>
      </div>
    </div>

    <!-- Success Toast -->
    <div
      v-if="showSuccessToast"
      class="fixed bottom-4 right-4 bg-green-500 text-white px-6 py-3 rounded-lg shadow-lg transform transition-transform duration-300 ease-in-out"
      :class="showSuccessToast ? 'translate-y-0' : 'translate-y-full'"
    >
      <div class="flex items-center">
        <CheckCircleIcon class="h-5 w-5 mr-2" />
        Settings saved successfully!
      </div>
    </div>

    <!-- Delete Confirmation Modal -->
    <div
      v-if="showDeleteConfirm"
      class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50"
      @click="showDeleteConfirm = false"
    >
      <div class="relative top-20 mx-auto p-5 border w-96 shadow-lg rounded-md bg-white" @click.stop>
        <div class="mt-3 text-center">
          <ExclamationTriangleIcon class="mx-auto h-12 w-12 text-red-600" />
          <h3 class="text-lg font-medium text-gray-900 mt-2">Reset All Settings</h3>
          <div class="mt-2 px-7 py-3">
            <p class="text-sm text-gray-500">
              This action will reset all your settings to their default values. This action cannot be undone.
            </p>
          </div>
          <div class="flex justify-center space-x-4 mt-4">
            <button
              @click="showDeleteConfirm = false"
              class="px-4 py-2 bg-gray-300 text-gray-800 text-base font-medium rounded-md shadow-sm hover:bg-gray-400 focus:outline-none focus:ring-2 focus:ring-gray-300"
            >
              Cancel
            </button>
            <button
              @click="confirmReset"
              class="px-4 py-2 bg-red-600 text-white text-base font-medium rounded-md shadow-sm hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500"
            >
              Reset Settings
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import {
  GlobeAltIcon,
  CloudIcon,
  CogIcon,
  CodeBracketIcon,
  BellIcon,
  ShieldCheckIcon,
  CheckIcon,
  ArrowPathIcon,
  RocketLaunchIcon,
  CheckCircleIcon,
  ExclamationCircleIcon,
  InformationCircleIcon,
  EyeIcon,
  EyeSlashIcon,
  EnvelopeIcon,
  ChatBubbleLeftRightIcon,
  ExclamationTriangleIcon,
  LockClosedIcon,
  KeyIcon
} from '@heroicons/vue/24/outline'

// Navigation sections
const sections = [
  { id: 'api', title: 'API Configuration', icon: GlobeAltIcon },
  { id: 'kubernetes', title: 'Kubernetes', icon: CloudIcon },
  { id: 'build', title: 'Build & Deploy', icon: CogIcon },
  { id: 'repository', title: 'Repository', icon: CodeBracketIcon },
  { id: 'notifications', title: 'Notifications', icon: BellIcon },
  { id: 'security', title: 'Security', icon: ShieldCheckIcon }
]

// Active section
const activeSection = ref('api')

// UI state
const saving = ref(false)
const showSuccessToast = ref(false)
const showDeleteConfirm = ref(false)
const showToken = ref(false)
const showGitToken = ref(false)

// Settings data
const settings = reactive({
  // API Configuration
  apiUrl: 'http://localhost:8080',
  authToken: '',
  
  // Kubernetes Configuration
  defaultEnvironment: 'dev',
  defaultNamespace: 'default',
  clusterContext: 'docker-desktop',
  defaultReplicas: 2,
  
  // Build & Deploy Configuration
  buildTimeout: 10,
  deployTimeout: 5,
  autoDeploy: true,
  autoRestart: false,
  
  // Repository Configuration
  defaultBranch: 'main',
  gitProvider: 'github',
  gitToken: '',
  
  // Notification Settings
  emailNotifications: true,
  slackNotifications: false,
  failureAlerts: true,
  emailAddress: '',
  slackWebhook: '',
  
  // Security Settings
  twoFactor: false,
  keyRotation: true
})

// Default settings for reset functionality
const defaultSettings = {
  apiUrl: 'http://localhost:8080',
  authToken: '',
  defaultEnvironment: 'dev',
  defaultNamespace: 'default',
  clusterContext: 'docker-desktop',
  defaultReplicas: 2,
  buildTimeout: 10,
  deployTimeout: 5,
  autoDeploy: true,
  autoRestart: false,
  defaultBranch: 'main',
  gitProvider: 'github',
  gitToken: '',
  emailNotifications: true,
  slackNotifications: false,
  failureAlerts: true,
  emailAddress: '',
  slackWebhook: '',
  twoFactor: false,
  keyRotation: true
}

const saveSettings = async () => {
  try {
    saving.value = true
    
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 1500))
    
    // Here you would make an actual API call to save settings
    // const response = await fetch('/api/v1/settings', {
    //   method: 'POST',
    //   headers: { 'Content-Type': 'application/json' },
    //   body: JSON.stringify(settings)
    // })
    
    // Show success toast
    showSuccessToast.value = true
    setTimeout(() => {
      showSuccessToast.value = false
    }, 3000)
    
  } catch (error) {
    console.error('Failed to save settings:', error)
    alert('Failed to save settings. Please try again.')
  } finally {
    saving.value = false
  }
}

const resetToDefaults = () => {
  showDeleteConfirm.value = true
}

const confirmReset = () => {
  Object.assign(settings, defaultSettings)
  showDeleteConfirm.value = false
  
  // Show success toast
  showSuccessToast.value = true
  setTimeout(() => {
    showSuccessToast.value = false
  }, 3000)
}

// Load settings on component mount (you would fetch from API)
// onMounted(async () => {
//   try {
//     const response = await fetch('/api/v1/settings')
//     const savedSettings = await response.json()
//     Object.assign(settings, savedSettings)
//   } catch (error) {
//     console.error('Failed to load settings:', error)
//   }
// })
</script>
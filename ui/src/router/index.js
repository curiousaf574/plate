import { createRouter, createWebHistory } from 'vue-router'
import Dashboard from '../views/Dashboard.vue'
import Projects from '../views/Projects.vue'
import Deployments from '../views/Deployments.vue'
import Settings from '../views/Settings.vue'
import AppDetail from '../views/AppDetail.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'dashboard',
      component: Dashboard
    },
    {
      path: '/projects',
      name: 'projects',
      component: Projects
    },
    {
      path: '/projects/:name',
      name: 'app-detail',
      component: AppDetail,
      props: true
    },
    {
      path: '/deployments',
      name: 'deployments',
      component: Deployments
    },
    {
      path: '/settings',
      name: 'settings',
      component: Settings
    }
  ]
})

export default router
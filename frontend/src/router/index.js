import { createRouter, createWebHashHistory } from 'vue-router'
import Monitor from '@/views/Monitor'
import Tracking from '@/views/Tracking'
import Table from '@/views/Table'

const routes = [{
  path: '/',
  name: 'Monitor',
  component: Monitor
}, {
  path: '/tracking',
  name: 'Tracking',
  component: Tracking
}, {
  path: '/table',
  name: 'Table',
  component: Table
}]

// noinspection JSCheckFunctionSignatures
const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router

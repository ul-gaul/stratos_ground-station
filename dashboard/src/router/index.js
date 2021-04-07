import { createRouter, createWebHashHistory } from 'vue-router'
import Monitor from '@/views/Monitor.vue'

const routes = [{
  path: '/',
  name: 'Monitor',
  component: Monitor
}, {
  path: '/tracking',
  name: 'Tracking',
  // route level code-splitting
  // this generates a separate chunk (about.[hash].js) for this route
  // which is lazy-loaded when the route is visited.
  component: () => import(/* webpackChunkName: "tracking" */ '@/views/Tracking.vue')
}, {
  path: '/table',
  name: 'Table',
  component: () => import(/* webpackChunkName: "table" */ '@/views/Table.vue')
}]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router

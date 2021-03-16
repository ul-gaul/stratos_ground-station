import Vue from 'vue'
import VueRouter from 'vue-router'
import Monitor from '../views/Monitor.vue'

Vue.use(VueRouter)

const routes = [
  { path: '/', component: Monitor },
  {
    path: '/tracking',
    name: 'Tracking',
    component: () => import('../views/Tracking.vue')
  }
]

const router = new VueRouter({
  routes
})

export default router

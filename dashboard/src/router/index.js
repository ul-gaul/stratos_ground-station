import Vue from 'vue'
import VueRouter from 'vue-router'
import Monitor from '../views/Monitor.vue'

Vue.use(VueRouter)

const routes = [
  { path: '/', component: Monitor },
  {
    path: '/tracking',
    name: 'Tracking',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/Tracking.vue')
  }
]

const router = new VueRouter({
  routes
})

export default router

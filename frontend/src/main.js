import { createApp, readonly } from 'vue'
import App from './App.vue'
import router from './router'
import {onDataReceived, state} from '@/modules/state'
import Wails from '@wailsapp/runtime'
import PrimeVue from 'primevue/config'
import ToastService from "primevue/toastservice";

import 'primevue/resources/themes/bootstrap4-dark-blue/theme.css'
import 'primevue/resources/primevue.min.css'
import 'primeicons/primeicons.css'
import 'primeflex/primeflex.scss'

Wails.Init(() => {
  Wails.Events.On("data", onDataReceived)
  const app = createApp(App)

  app.use(PrimeVue, { ripple: true })
  app.use(ToastService)
  app.use(router)
  
  app.config.globalProperties.$logger = Wails.Log
  app.config.globalProperties.$events = Wails.Events
  app.config.globalProperties.$controller = window.backend.controller
  app.config.globalProperties.$state = readonly(state)
  
  app.mount('#app')
})

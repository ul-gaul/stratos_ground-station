import { createApp, readonly } from 'vue'
import App from './App.vue'
import router from './router'
import { onDataReceived, state } from '@/modules/state'

import PrimeVue from 'primevue/config'
import 'primevue/resources/themes/bootstrap4-dark-blue/theme.css'
import 'primevue/resources/primevue.min.css'
import 'primeicons/primeicons.css'
import 'primeflex/primeflex.scss'

window.go_ondata = onDataReceived;

const app = createApp(App)

// noinspection JSCheckFunctionSignatures
app.use(PrimeVue, { ripple: true })
app.use(router)

app.config.globalProperties.$state = readonly(state)
app.mount('#app')
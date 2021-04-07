<template>
<Menubar :model="items" class="navbar">
  <template #start>
    <img alt="GAUL - Stratos"
         src="@/assets/gaul.webp"
         class="p-mr-2"
         :height="40"/>
  </template>
</Menubar>
</template>

<script>
import Menubar from 'primevue/menubar'
import '@/assets/gaul.webp'

export default {
  name: "Navbar",
  components: { Menubar },
  data() {
    return {
      items: [{
        label: 'Monitor',
        icon: 'pi pi-chart-line',
        to: '/'
      }, {
        label: 'Tracking',
        icon: 'pi pi-map',
        to: '/tracking'
      }, {
        label: 'Table',
        icon: 'pi pi-table',
        to: '/table'
      }]
    }
  }
}
</script>

<style lang="scss" scoped>
.navbar {
  border-radius: 0;
}
</style>
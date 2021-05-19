<template>
  <div>
    <Navbar/>
    <router-view/>
  </div>
</template>
<script>
import Navbar from '@/components/Navbar'
import { onDataReceived } from '@/modules/state'

export default {
  components: { Navbar },
  async created() {
    let data = await this.$controller.Connect()
    if (data != null) {
      onDataReceived(data)
    }
  }
}
</script>

<style lang="scss">
body {
  margin: 0;
  padding: 0;
  background: var(--surface-0);
  color: var(--text-color);
}

* {
  ::-webkit-scrollbar {
    width: 5px;
    background-color: var(--surface-d)
  }

  ::-webkit-scrollbar-thumb {
    background-color: var(--primary-color)
  }
}
</style>
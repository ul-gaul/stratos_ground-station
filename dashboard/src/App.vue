<template>
  <Navbar/>
  <router-view/>
</template>
<script>
import Navbar from '@/components/Navbar'

export default {
  components: { Navbar },
  async created() {
    // noinspection JSUnresolvedVariable
    const goConnect = window.go_connect;
    if (typeof goConnect === 'function') {
      window.go_ondata(await goConnect());
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
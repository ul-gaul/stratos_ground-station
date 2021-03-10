<template>
  <div id="app">
    <b-container>
      <header class="mb-5 text-center">
        <img alt="Gaul logo"
             src="./assets/gaul.webp"
             class="align-self-center"
             style="max-width: 250px">
        <h1 class="display-4">Ground Station Dashboard</h1>

        <b-row class="d-flex justify-content-center text-center">
          <b-col v-for="(link, idx) in links" :key="idx"
                 cols="1">
            <router-link :to="link.path">{{ link.name }}</router-link>
          </b-col>
        </b-row>
      </header>

      <router-view />
    </b-container>
  </div>
</template>

<script>

export default {
  name: 'App',
  components: {},
  data() {
    return {
      links: [
        { path: '/', name: 'Monitor' },
        { path: '/tracking', name: 'Tracking' }
      ]
    }
  }
}
</script>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}

#nav {
  padding: 30px;

  a {
    font-weight: bold;
    color: #2c3e50;

    &.router-link-exact-active {
      color: #42b983;
    }
  }
}
</style>

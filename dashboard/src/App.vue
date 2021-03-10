<template>
  <div id="app">
    <b-container>
      <header class="mb-5 text-center">
        <img alt="Gaul logo"
             src="./assets/gaul.webp"
             class="align-self-center"
             style="max-width: 250px">
        <h1 class="display-4">Ground Station Dashboard</h1>
      </header>

      <b-row class="d-flex flex-wrap justify-content-around"
             cols-lg="4" cols-md="2" cols="1">
        <b-col v-for="i in 10" :key="i"><PlotChart/></b-col>
      </b-row>
    </b-container>
  </div>
</template>

<script>
import PlotChart from './components/PlotChart.vue'

export default {
  name: 'App',
  components: {
    PlotChart,
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
</style>

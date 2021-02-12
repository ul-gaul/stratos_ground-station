<template>
  <div id="app">
    <img alt="Gaul logo" src="./assets/gaul.webp" style="width: 250px;">
    <h1>Ground Station Dashboard</h1>
    <div class="grid-container">
      <div class="grid-item"><PlotChart id="plot-chart"></PlotChart></div>
      <div class="grid-item"><PlotChart id="plot-chart"></PlotChart></div>
    </div>
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

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
#plot-chart {
  width: 20%;
  height: 20%;
}
.grid-container {
  display: grid;
  grid-gap: 50px 50px;
}
</style>

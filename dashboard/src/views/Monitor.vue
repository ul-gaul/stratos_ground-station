<template>
  <b-row class="d-flex flex-wrap justify-content-around"
         cols-lg="4" cols-md="2" cols="1">
    <b-col v-for="i in 10" :key="i">
      <PlotChart />
    </b-col>
  </b-row>
</template>

<script>
import PlotChart from '../components/PlotChart.vue'

export default {
  name: "Monitor",
  components: { PlotChart }
}
</script>

<style lang="scss" scoped>

</style>
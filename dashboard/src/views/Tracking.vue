<template>
  <b-row class="d-flex flex-wrap justify-content-around align-content-stretch"
         cols-lg="4" cols-md="2" cols="1">
    <b-col v-for="i in 10" :key="i">
      <PlotChart :chart-data="chartData" />
    </b-col>
  </b-row>
</template>

<script>
import PlotChart from '../components/PlotChart.vue';
import rdm from '../modules/utils/random';

export default {
  name: "Tracking",
  components: { PlotChart },
  data() {
    return {
      chartData: {
        labels: [
          'Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun',
          'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec' ],
        datasets: [ {
          label: 'Dataset 1',
          backgroundColor: rdm.rgba(0.1),
          data: rdm.coords({ qty: 12 })
        } ]
      }
    }
  }
}
</script>

<style lang="scss" scoped>

</style>
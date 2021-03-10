<template>
  <b-row class="d-flex flex-wrap justify-content-around"
         cols-lg="4" cols-md="2" cols="1">
    <b-col v-for="i in 10" :key="i">
      <PlotChart :chartData="chartData()" :width="200" :height="200" />
    </b-col>
  </b-row>
</template>

<script>
import PlotChart from '../components/PlotChart.vue';
import rdm from '../modules/utils/random';

export default {
  name: "Monitor",
  components: { PlotChart },
  methods: {
    chartData() {
      return {
        labels: [ 'Jan', 'Feb', 'Mar', 'Apr', 'May' ],
        datasets: [ {
          label: 'Dataset 1',
          backgroundColor: rdm.rgba(0.1),
          data: rdm.coords()
        } ]
      };
    }
  }
}
</script>

<style lang="scss" scoped>

</style>
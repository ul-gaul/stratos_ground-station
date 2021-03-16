<template>
  <b-row fluid>
    <b-col lg="8" md="12">
      <b-row class="d-flex flex-wrap justify-content-around align-items-stretch"
             cols-xl="3" cols-sm="2">
        <b-col v-for="i in 5" :key="i">
          <Temperature :chart-data="chartData()" :width="200" :height="200" />
        </b-col>
      </b-row>
    </b-col>

    <b-col lg="4" md="12" class="h-100">
      <Temperature :chart-data="chartData()" class="h-100" />
    </b-col>
  </b-row>
</template>

<script>
import Temperature from '../components/Temperature.vue';
import rdm from '../modules/utils/random';

export default {
  name: "Monitor",
  components: { Temperature },
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
<template>
  <b-row class="d-flex flex-wrap justify-content-around align-content-stretch"
         cols-lg="3" cols-md="2" cols="1">
    <b-col lg="8">
      <Map />
    </b-col>
    <b-col>
      <Altitude :chart-data="chartData()" :width="200" :height="200" />
      <Vitesse :chart-data="chartData()" :width="200" :height="200" />
    </b-col>
  </b-row>
</template>

<script>
import Map from '../components/Map.vue';
import Altitude from '../components/Altitude.vue';
import Vitesse from '../components/Vitesse.vue';
import rdm from '../modules/utils/random';

export default {
  name: "Tracking",
  components: { 
    Map, 
    Altitude, 
    Vitesse 
  },
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
.d-flex {
  margin-top: 20px;
}
</style>
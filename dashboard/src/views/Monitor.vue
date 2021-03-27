<template>
  <b-row fluid>
    <b-col lg="8" md="12">
      <b-row class="d-flex flex-wrap justify-content-around align-items-stretch"
             cols-xl="3" cols-sm="2">
        <b-col key="1"><TemperatureInt :chart-data="chartData()" :width="200" :height="200" /></b-col>        
        <b-col key="2"><TemperatureExt :chart-data="chartData()" :width="200" :height="200" /></b-col>        
        <b-col key="3"><Pression :chart-data="chartData()" :width="200" :height="200" /></b-col>        
        <b-col key="4"><Altitude :chart-data="chartData()" :width="200" :height="200" /></b-col>        
        <b-col key="5"><Vitesse :chart-data="chartData()" :width="200" :height="200" /></b-col>
      </b-row>
    </b-col>
    <b-col lg="4" md="12" class="h-100">
      <Payload :chart-data="chartData()" class="h-100" />
    </b-col>
  </b-row>
</template>

<script>
import Altitude from '../components/Altitude.vue';
import Payload from '../components/Payload.vue';
import Pression from '../components/Pression.vue';
import TemperatureInt from '../components/TemperatureInt.vue';
import TemperatureExt from '../components/TemperatureExt.vue';
import Vitesse from '../components/Vitesse.vue';
import rdm from '../modules/utils/random';

export default {
  name: "Monitor",
  components: {
    Altitude,
    Payload,
    Pression, 
    TemperatureInt,
    TemperatureExt,
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
.h-100 {
  margin-top: 20px;
}
</style>
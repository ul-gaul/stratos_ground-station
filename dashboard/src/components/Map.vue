<template>
<mapbox
      access-token="pk.eyJ1IjoiZ3VpbGxhdW1lbGFuZHJ5IiwiYSI6ImNrOGpsaDQwbDBpNXAzZm12dWF5MTlpMzEifQ.8SaHyf8Dz2NG5G1czLGytA"
      :map-options="{
        style: 'mapbox://styles/mapbox/light-v10',
        center: [-71.300000, 46.800000],
        zoom: 10,
      }"
    />  
</template>

<script>
import Mapbox from 'mapbox-gl-vue'

export default {
  name: 'Map',
  components: { Mapbox },
};
</script>

<style>
#map {
  width: 100%;
  height: 100%;
  border: solid 1px;
  }
</style>
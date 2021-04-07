<template>
  <div class="p-grid p-nogutter">
    <div class="p-col-12 p-lg-8">
      <div class="p-grid p-nogutter">
        <div v-for="(val, i) in charts" :key="i"
             class="p-col-12 p-lg-6 p-p-1">
          <StateTimeDataChart :states="val" :height="400" />
        </div>
      </div>
    </div>
    <div class="p-col-12 p-lg-4 p-p-1">
      <StateTimeDataChart states="payloads" :height="800" />
    </div>
  </div>
</template>
<script>
import StateTimeDataChart from '@/components/StateTimeDataChart'

export default {
  name: 'Monitor',
  components: { StateTimeDataChart },
  data() {
    return {
      charts: [
        [ 'internalTemps:Internal Temp. (°C)', 'externalTemps:External Temp. (°C)' ],
        'pressures:Pressure (Pa)',
        'altitudes:Altitude (m)',
        'speeds:Speed (m/s)'
      ]
    }
  }
}
</script>

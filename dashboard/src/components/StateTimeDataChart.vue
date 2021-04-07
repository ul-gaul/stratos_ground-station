<template>
  <canvas ref="canvas" :width="width" :height="height" />
</template>

<script>
import Chart from 'chart.js/auto'
import rdm from '@/modules/utils/random'

export default {
  name: "StateTimeDataChart",
  props: {
    // States to display on this chart
    states: {
      type: [Array, String],
      required: true
    },
    width: Number,
    height: Number
  },

  data() {
    return {
      options: {
        responsive: true,
        hoverMode: 'index',
        maintainAspectRatio: false,
        stacked: false,
      },
      dataOpts: null,
      chart: null,
    }
  }, // END data()

  mounted() {
    let datasets = []
    for (let stateKey of this.stateKeys) {
      let [key, label] = stateKey.split(':');

      if (!Array.isArray(this.$state[key])) {
        console.warn(`Unknown or invalid state "${key}"`);
      } else {
        datasets.push({
          label: label || key,
          fill: false,
          borderColor: rdm.rgba(0.6),
          data: this.$state[key]
        });
      }
    }

    this.dataOpts = {
      labels: this.times,
      datasets
    };

    window.Chart = Chart
    window.canvas = this.$refs.canvas
    this.chart = new Chart(this.$refs.canvas, {
      type: 'line',
      options: this.options,
      data: this.dataOpts
    });
  }, // END mounted()

  computed: {
    times() { return this.$state.times },
    stateKeys() {
      return typeof this.states === 'string'
          ? this.states.split(',')
          : this.states
    }
  },

  watch: {
    times: {
      deep: true,
      handler() {
        console.log('data changed!')
      }
    }
  }
}
</script>

<style lang="scss" scoped>
canvas {
  background-color: var(--surface-a);
  color: var(--primary-color-text)
}
</style>
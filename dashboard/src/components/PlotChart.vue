<template>
	<chart :chart-data="datacollection" :width="200" :height="200"></chart>
</template>

<script>
import Chart from "./../chart-model.js";
export default {
  components: {
    Chart
  },
  data() {
    return {
      datacollection: null
    };
  },
  mounted() {
    this.fillData();
  },
  methods: {
    fillData() {
      this.datacollection = {
        labels: ["January","February","March","April","May","June","July","August","September"],
        // labels: [this.rand(), this.rand()],
        datasets: [
          {
            label: "Data 1",
            backgroundColor: "red",
            data: [this.rand(), this.rand(), this.rand(), this.rand(), this.rand(), this.rand(), this.rand(), this.rand(), this.rand()]
          },
          {
            label: "Data 2",
            backgroundColor: "gold",
            data: [this.rand(), this.rand(), this.rand(), this.rand(), this.rand(), this.rand(), this.rand(), this.rand(), this.rand()]
          }
        ]
      };
    },
    rand() {
      return Math.floor(Math.random() * (50 - 5 + 1)) + 5;
    }
  }
};
</script>

<style>
</style>
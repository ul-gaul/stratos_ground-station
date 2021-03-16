<script>
import { Line, mixins } from 'vue-chartjs'

export default {
  extends: Line,
  mixins: [mixins.reactiveProp],
  // props: ['chartData'],
  components: {},
  data() {
    return {
      options: {
        scales: {
          yAxes: [{
            ticks: { beginAtZero: true },
            gridLines: { display: true },
          }],
          xAxes: [{
            gridLines: { display: true },
          }]
        }, // END options.scales
        legend: { display: true },
        responsive: true,
        maintainAspectRatio: false,
      }, // END options
    };
  },
  mounted() {
    this.renderChart(this.chartData, this.options);
  },
};
</script>

<style>
</style>
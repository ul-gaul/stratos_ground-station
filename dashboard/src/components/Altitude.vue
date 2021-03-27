<script>
import { Bar, mixins } from 'vue-chartjs'

export default {
  extends: Bar,
  mixins: [mixins.reactiveProp],
  // props: ['chartData'],
  components: {},
  data() {
    return {
      options: {
        title: {
            display: true,
            text: 'Altitude'
        },
        scales: {
          yAxes: [{
            ticks: { beginAtZero: true },
            gridLines: { display: true },
          }],
          xAxes: [{
            gridLines: { display: true },
          }]
        }, // END options.scales
        legend: { display: false },
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
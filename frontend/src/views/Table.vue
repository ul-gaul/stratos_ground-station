<template>
  <div style="height: calc(100vh - 64px)">
    <DataTable :value="tableData" class="p-datatable-sm" :scrollable="true" scrollHeight="flex">
      <Column v-for="(header, field) in columns"
              :field="field"
              :header="header"
              :key="field"
              :sortable="true" />
    </DataTable>
  </div>
</template>

<script>
import DataTable from 'primevue/datatable'
import Column from 'primevue/column'

export default {
  name: "Table",
  components: { DataTable, Column },
  data() {
    return {
      columns: {
        time: 'Time (ms)',
        altitude: 'Altitude',
        pressure: 'Pressure',
        longitude: 'Longitude',
        latitude: 'Latitude',
        internalTemp: 'Internal Temp. (°C)',
        externalTemp: 'External Temp. (°C)',
      }
    }
  },
  computed: {
    tableData() { return this.$state.raw }
  },
}
</script>

<style lang="scss" scoped>

</style>
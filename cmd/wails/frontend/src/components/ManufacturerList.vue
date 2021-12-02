<template>
  <div class="page-body">
    <div class="container-fluid">
      <div class="row">
        <div class="col-12">
          <DataTable :columns="columns"
                     :rows="manufacturers"
                     :total="$store.state.count.manufacturers"
                     @update="load">
          </DataTable>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {Manufacturer} from "@/models";
import DataTable from "@/components/DataTable";


export default {
  name: "ManufacturerList",

  components: {
    DataTable,
  },

  data() {
    return {
      columns: [
        {
          property: 'ID',
          sortable: true,
          label: '#',
          noSearch: true,
          forEach: (d) => `#${d}`
        },
        {
          property: 'Name',
          label: 'Name',
          sortable: true,
          defaultSort: true,
          defaultSearch: true,
          class: 'w-100'
        },
        {
          property: 'CreatedAt',
          label: 'Created',
          noSearch: true,
          forEach: (d) => this.formatDate(d, 'DD/MM/YYYY')
        },
      ],
      manufacturers: [],
      editingAsset: null,
    }
  },

  methods: {
    async load(opts) {
      this.manufacturers = (
          await window.go.sqlite.manufRepository.Paginate({
            Query: _.has(opts, 'query') ? opts.query : '',
            QueryField: _.has(opts, 'queryField') ? opts.queryField : 'name',
            Limit: _.has(opts, 'limit') ? opts.limit : 10,
            Offset: _.has(opts, 'offset') ? opts.offset : 0,
            Order: _.has(opts, 'sort') ? opts.sort : "id desc",
          })
      ).map(m => new Manufacturer(m))
    },

    formatDate: function (val, format) {
      return this.$dayjs(val).format(format)
    },
  },

  mounted() {
    this.load()
  }
}
</script>

<style scoped>

</style>
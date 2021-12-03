<template>
  <div class="page-body">
    <div class="container-fluid">
      <div class="row mb-4">
        <h2 class="page-title">
          Manufacturers
        </h2>
      </div>
      <div class="row">
        <div class="col-12">
          <DataTable :columns="columns"
                     :rows="manufacturers"
                     :total="$store.state.count.manufacturers"
                     @update="load">
            <template v-slot:default="props">
              <td data-label="Delete">
                <a class="btn btn-danger btn-icon" @click="deleteManufacturer(props.row)"><i class="ti ti-trash"></i></a>
              </td>
            </template>
          </DataTable>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {Manufacturer} from "@/models";
import DataTable from "@/components/DataTable";
import {ModelDto} from "@/models.dto";


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
        {
          label: '',
          class: 'w-1'
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

    async deleteManufacturer(manufacturer) {
      this.$confirm("Deleting Asset",
        "Are you sure that you want to delete the Manufacturer \"" + manufacturer.Name + "\"", async (status) => {
          if (!status) {
            return
          }

          try {
            await window.go.sqlite.manufRepository.Delete(ModelDto.fromObject(manufacturer))
          } catch (error) {
            this.$showDialog("Error deleting asset", error)
          }

          await this.load()
      })
    },

    formatDate: function (val, format) {
      return this.$dayjs(val).format(format)
    }
  },

  mounted() {
    this.load()
  }
}
</script>

<style scoped>

</style>
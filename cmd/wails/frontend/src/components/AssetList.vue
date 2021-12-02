<template>
  <div class="page-body">
    <div class="container-fluid">
      <div class="row">
        <div class="col-12 d-flex justify-content-end mb-3">
          <button class="btn btn-primary w-auto" @click="editAsset(0)">Add asset</button>
        </div>
      </div>
      <div class="row">
        <div class="col-12">
          <DataTable :columns="columns"
                     :rows="assets"
                     :total="$store.state.count[$route.params.type]"
                     @update="load">
            <template v-slot:default="props">
              <td data-label="Edit">
                <a href="#" @click="editAsset(props.row)">Edit</a>
              </td>
              <td data-label="Delete">
                <a href="#" @click="deleteAsset(props.row)">Delete</a>
              </td>
            </template>
          </DataTable>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import DataTable from "@/components/DataTable";
import {Asset} from "@/models";
import {AssetDto} from "@/models.dto";

export default {
  name: "AssetList",

  components: {DataTable},

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

        },
        {
          property: 'Manufacturer.Name',
          label: 'Manufacturer',
        },
        {
          property: 'Description',
          label: 'Description',
        },
        {
          property: 'UpdatedAt',
          label: 'Updated',
          sortable: true,
          noSearch: true,
          forEach: (d) => this.formatDate(d, 'DD/MM/YYYY')
        },
        {
          label: '',
          class: 'w-1'
        },
        {
          label: '',
          class: 'w-1'
        },
      ],
      assets: [],
      editingAsset: null,
    }
  },

  methods: {
    async load(opts) {
      try {
        this.assets = (
            await window.go.sqlite.assetRepository.Paginate(
                this.$route.params.type, {
                  Query: _.has(opts, 'query') ? opts.query : '',
                  //QueryField: _.has(opts, 'queryField') ? opts.queryField : 'name',
                  Limit: _.has(opts, 'limit') ? opts.limit : 10,
                  Offset: _.has(opts, 'offset') ? opts.offset : 0,
                  Order: _.has(opts, 'sort') ? opts.sort : "id desc",
                }
            )
        ).map(a => new Asset(a))
      } catch (error) {
        this.$showError("An error occurred", error)
      }
    },

    editAsset(asset) {
      this.editingAsset = asset

      this.$store.dispatch("showModal", {
        classes: ['modal-lg'],

        component: "AssetEditModal",

        props: {
          id: this.editingAssetId
        },

        on: {
          save: this.onSave
        }
      })
    },

    deleteAsset(asset) {
      this.$confirm("Deleting Asset",
          "Are you sure that you want to delete the asset \"" + asset.Name + "\"", async (status) => {
            if (! status) {
              return
            }

            try {
              await window.go.sqlite.assetRepository.Delete(AssetDto.fromObject(asset))
            } catch (error) {
              this.$showError("Error deleting asset", error)
            }

            await this.load()
          })
    },

    onSave(error) {
      if (error !== null) {
        this.$showError("An error occurred", error)
      }

      this.load()
    },

    formatDate: function (val, format) {
      return this.$dayjs(val).format(format)
    },
  },

  watch: {
    "$route.params.type": {
      handler() {
        if (this.$route.name === "assets") {
          this.load()
          this.$store.dispatch("syncAssetCounts")
        }
      },

      immediate: true
    }
  },

  computed: {
    editingAssetId: function () {
      return this.editingAsset !== null ? this.editingAsset.ID : 0;
    }
  },
}
</script>

<style scoped>
</style>
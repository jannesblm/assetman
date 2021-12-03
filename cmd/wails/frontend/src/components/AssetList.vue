<template>
  <div class="page-body">
    <div class="container-fluid">
      <div class="row mb-4">
        <div class="col">
          <h2 class="page-title">
            {{ title }}
          </h2>
        </div>
        <div class="col">
          <div class="col-12 d-flex justify-content-end">
            <button :disabled="!$store.getters.isAdmin" class="btn btn-primary w-auto" @click="editAsset(0)">Add asset
            </button>
          </div>
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
                <a class="btn btn-success btn-icon" @click="editAsset(props.row)"><i class="ti ti-eye"></i></a>
              </td>
              <td data-label="Actions">
                <div class="dropdown">
                  <button aria-expanded="false" class="btn dropdown-toggle align-text-top"
                          data-bs-boundary="viewport" data-bs-toggle="dropdown">
                    Actions
                  </button>
                  <div class="dropdown-menu dropdown-menu-end">
                    <a v-if="$store.getters.isAdmin" class="dropdown-item cursor-pointer text-danger"
                       type="button" @click="deleteAsset(props.row)">
                      <i class="ti ti-trash mr-3"></i>Delete
                    </a>
                    <a class="dropdown-item cursor-pointer" @click="searchNvd(props.row.Description, props.row.ID)">
                      <i class="ti ti-search mr-3"></i>Search NVD
                    </a>
                  </div>
                </div>
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
import {ModelDto} from "@/models.dto";

export default {
  name: "AssetList",

  components: {DataTable},

  data() {
    return {
      columnDefs: {
        hardware: [
          {
            property: 'ID',
            label: '#',
            sortable: true,
            prefix: "assets",
            noSearch: true,
            forEach: (d) => `#${d}`
          },
          {
            property: 'Name',
            label: 'Name',
            sortable: true,
            prefix: "assets",
            defaultSort: true,
            defaultSearch: true,
          },
          {
            property: 'TypeName',
            label: 'Type',
          },
          {
            property: 'HardwareAsset.ModelName',
            label: 'Model',
          },
          {
            property: 'Manufacturer.Name',
            label: 'Manufacturer',
          },
          {
            property: 'Description',
            label: 'Description',
            forEach: (d) => d.length > 20 ? (d.substr(0, 20) + "...") : d
          },
          {
            property: 'UpdatedAt',
            label: 'Updated',
            sortable: true,
            prefix: "assets",
            noSearch: true,
            forEach: (d) => this.formatDate(d, 'DD/MM/YYYY')
          },
          {
            property: 'Report',
            label: 'Vuln.',
            trusted: true,
            noSearch: true,
            forEach: r => {
              if (_.has(r, "Cves") && r.Cves !== null) {
                return `<span class="text-danger font-weight-bold"><i class="ti ti-alert-triangle"></i>&nbsp;${r.Cves.length}</span>`
              }

              return `<span>-</span>`
            }
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

        software: [
          {
            property: 'ID',
            label: '#',
            sortable: true,
            prefix: "assets",
            noSearch: true,
            forEach: (d) => `#${d}`
          },
          {
            property: 'Name',
            label: 'Name',
            sortable: true,
            prefix: "assets",
            defaultSort: true,
            defaultSearch: true,
          },
          {
            property: 'SoftwareAsset.Version',
            label: 'Version',
          },
          {
            property: 'TypeName',
            label: 'Type',
          },
          {
            property: 'Manufacturer.Name',
            label: 'Manufacturer',
          },
          {
            property: 'Description',
            label: 'Description',
            forEach: (d) => d.length > 20 ? (d.substr(0, 20) + "...") : d
          },
          {
            property: 'PurchasedAt',
            label: 'Purchased',
            sortable: true,
            prefix: "assets",
            noSearch: true,
            forEach: (d) => this.formatDate(d, 'DD/MM/YYYY')
          },
          {
            property: 'Report',
            label: 'Vuln.',
            trusted: true,
            noSearch: true,
            forEach: r => {
              if (_.has(r, "Cves") && r.Cves !== null) {
                return `<span class="text-danger font-weight-bold"><i class="ti ti-alert-triangle"></i>&nbsp;${r.Cves.length}</span>`
              }

              return `<span>-</span>`
            }
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
      },
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
                  QueryField: _.has(opts, 'queryField') ? opts.queryField : 'name',
                  Limit: _.has(opts, 'limit') ? opts.limit : 10,
                  Offset: _.has(opts, 'offset') ? opts.offset : 0,
                  Order: _.has(opts, 'sort') ? opts.sort : "assets.id desc",
                }
            )
        ).map(a => new Asset(a))
      } catch (error) {
        this.$showDialog("An error occurred", error)
      }
    },

    editAsset(asset) {
      this.editingAsset = asset

      this.$store.dispatch("showModal", {
        classes: ['modal-full-width'],

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
            if (!status) {
              return
            }

            try {
              await window.go.sqlite.assetRepository.Delete(ModelDto.fromObject(asset))
            } catch (error) {
              this.$showDialog("Error deleting asset", error)
            }

            await this.load()
          })
    },

    onSave(error, asset, searchNvd) {
      if (error !== null) {
        this.$showDialog("An error occurred", error)
      }

      if (error === null && searchNvd) {
        this.searchNvd(asset.Description, asset.ID, this.load)
      } else {
        this.load()
      }
    },

    formatDate: function (val, format) {
      return this.$dayjs(val).format(format)
    },

    searchNvd(keyword, id, callback) {
      this.$store.dispatch("showModal", {
        classes: ['modal-sm', 'modal-dialog-centered'],

        component: "VulnerabilityModal",

        props: {
          keyword: keyword,
          assetId: id
        },

        on: {
          searchDone: typeof callback === 'function' ? callback() : this.load
        }
      })
    }
  },

  watch: {
    "$route.params.type": {
      handler() {
        if (this.$route.name === "assets") {
          this.load()

          this.$store.dispatch("syncAssetCounts")
          this.$store.dispatch("syncManufacturers")
        }
      },

      immediate: true
    }
  },

  computed: {
    editingAssetId: function () {
      return this.editingAsset !== null ? this.editingAsset.ID : 0;
    },

    columns: function () {
      return this.columnDefs[this.$route.params.type]
    },

    title: function () {
      return this.$route.params.type.charAt(0).toUpperCase() + this.$route.params.type.slice(1);
    }
  },
}
</script>

<style scoped>

</style>
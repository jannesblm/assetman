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
          <div class="card">
            <div class="card-body border-bottom py-3">
              <div class="d-flex">
                <div class="text-muted">
                  Show
                  <div class="mx-2 d-inline-block">
                    <input v-model="limit" aria-label="Invoices count" class="form-control form-control-sm" size="3"
                           type="text">
                  </div>
                  entries
                </div>
                <div class="ms-auto text-muted">
                  Search:
                  <div class="ms-2 d-inline-block">
                    <input v-model="query" class="form-control form-control-sm" type="text">
                  </div>
                </div>
              </div>
            </div>
            <div class="table-responsive">
              <table class="table table-vcenter card-table">
                <thead>
                <tr>
                  <th>#</th>
                  <th>Name</th>
                  <th>Manufacturer</th>
                  <th>Description</th>
                  <th>Updated</th>
                  <th class="w-1"></th>
                </tr>
                </thead>
                <tbody>
                <tr v-for="asset in assets" :key="asset.ID">
                  <td data-label="ID">{{ asset.ID }}</td>
                  <td data-label="Name">{{ asset.Name }}</td>
                  <td data-label="Manufacturer">{{ asset.Manufacturer.Name }}</td>
                  <td data-label="Description">{{ asset.Description }}</td>
                  <td data-label="UpdatedAt">{{ formatDate(asset.UpdatedAt, "DD/MM/YYYY") }}</td>
                  <td data-label="Edit">
                    <a href="#" @click="editAsset(asset)">Edit</a>
                  </td>
                </tr>
                </tbody>
              </table>
            </div>
            <div class="card-footer d-flex align-items-center">
              <p class="m-0 text-muted">Showing {{ assets.length }} to {{ limit }} of
                {{ $store.state.assetCount[$route.params.type] }}
                entries</p>
              <ul class="pagination m-0 ms-auto">
                <li :class="[{disabled: page <= availablePages}, 'page-item']">
                  <a class="page-link" href="#" tabindex="-1">
                    <i class="ti ti-chevron-left"></i>
                    prev
                  </a>
                </li>
                <li v-for="i in availablePages" :key="i" class="page-item">
                  <a :class="[{active: page === i}, 'page-link']" @click="page = i">{{ i }}</a>
                </li>
                <li :class="[{disabled: page >= availablePages}, 'page-item']">
                  <a class="page-link" href="#">
                    next
                    <i class="ti ti-chevron-right"></i>
                  </a>
                </li>
              </ul>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <AssetEdit :id="editingAssetId" ref="editModal"
             @saved="onSave"/>
</template>

<script>
import AssetEdit from "@/components/AssetEdit";
import {Asset} from "@/models";
import dayjs from "dayjs";
import Swal from 'sweetalert2'


export default {
  name: "AssetList",

  components: {AssetEdit},

  data() {
    return {
      limit: 10,
      page: 1,
      availablePages: 1,
      offset: 0,
      query: '',
      assets: [],
      editingAsset: null,
    }
  },

  methods: {
    async loadAssets() {
      this.assets = (
          await window.go.sqlite.repository.PaginateByTypeAndName(this.$route.params.type, this.query, {
            Limit: this.limit,
            Offset: this.offset,
            Order: "id desc",
          })
      ).map(a => new Asset(a))

      this.page = 1
      this.availablePages = Math.ceil(this.assets.length / this.limit)

      if (this.availablePages === 0) {
        this.availablePages = 1
      }
    },

    editAsset(asset) {
      this.editingAsset = asset
      $(this.$refs.editModal.$el).modal("show")
    },

    onSave(error) {
      if (error !== null) {
        Swal.fire({
          icon: 'error',
          title: 'Oops...',
          text: error,
        })
      }

      this.loadAssets()
    },

    formatDate: function (val, format) {
      return dayjs(val).format(format)
    },
  },

  watch: {
    query: function () {
      this.loadAssets()
    },
    "$route.params.type": {
      handler(cname) {
        console.log('ROUTE!!', cname);
        this.loadAssets()

        this.$store.commit("syncCount")
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
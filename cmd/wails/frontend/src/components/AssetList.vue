<template>
  <div class="page-body">
    <div class="container-fluid">
      <div class="col-12">
        <div class="card">
          <div class="card-body border-bottom py-3">
            <div class="d-flex">
              <div class="text-muted">
                Show
                <div class="mx-2 d-inline-block">
                  <input aria-label="Invoices count" class="form-control form-control-sm" size="3" type="text"
                         v-model="limit">
                </div>
                entries
              </div>
              <div class="ms-auto text-muted">
                Search:
                <div class="ms-2 d-inline-block">
                  <input class="form-control form-control-sm" type="text" v-model="query">
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
                <td data-label="UpdatedAt">{{ asset.UpdatedAt }}</td>
                <td data-label="Edit">
                  <a href="#" @click="editAsset(asset)">Edit</a>
                </td>
              </tr>
              </tbody>
            </table>
          </div>
          <div class="card-footer d-flex align-items-center">
            <p class="m-0 text-muted">Showing {{ assets.length }} to {{ limit }} of {{ $store.state.assetCount }} entries</p>
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
  <AssetEdit ref="editModal" :id="editingAssetId"></AssetEdit>
</template>

<script>
import AssetEdit from "@/components/AssetEdit";

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
      editingAssetId: null,
    }
  },

  methods: {
    async loadAssets() {
      this.assets = await window.go.sqlite.repository.PaginateByName(this.query, {
        Limit: this.limit,
        Offset: this.offset,
        Order: "id desc",
      })

      this.page = 1
      this.availablePages = Math.ceil(this.assets.length / this.limit)

      if (this.availablePages === 0) {
        this.availablePages = 1
      }
    },

    editAsset(asset) {
      this.editingAssetId = asset.ID
      $(this.$refs.editModal.$el).modal('show')
    },
  },

  watch: {
    query: function() {
      this.loadAssets()
    },
  },

  mounted() {
    this.loadAssets()
    this.$store.commit('updateAssetCount')
    this.$store.commit('updateManufacturers')
  }
}
</script>

<style scoped>

</style>
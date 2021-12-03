<template>
  <div class="row">
    <div class="col-lg-6">
      <div class="mb-3">
        <label class="form-label">License type</label>
        <input v-model="asset.LicenseType" :readonly="!$store.getters.isAdmin" class="form-control" list="license-types"
               type="text">
        <datalist id="license-types">
          <option value="Public Domain"/>
          <option value="Permissive"/>
          <option value="Proprietary"/>
          <option value="Subscription"/>
        </datalist>
      </div>
    </div>
    <div class="col-lg-6">
      <div class="mb-3">
        <label class="form-label">License key</label>
        <input v-model="asset.LicenseType" :readonly="!$store.getters.isAdmin" class="form-control" type="text">
      </div>
    </div>
  </div>
  <div class="row">
    <div class="col-lg-6">
      <div class="mb-3">
        <label class="form-label">Version</label>
        <input v-model="asset.Version" :readonly="!$store.getters.isAdmin" class="form-control" type="text">
      </div>
    </div>
  </div>
</template>

<script>
import {SoftwareAsset} from "@/models";

export default {
  name: "SoftwareOptions",

  emits: ["changed"],

  props: {
    initialAsset: {
      type: SoftwareAsset,
      required: true,
    }
  },

  data() {
    return {
      asset: null,
    }
  },

  created() {
    this.asset = this.initialAsset
  },

  watch: {
    initialAsset: {
      handler: function (newVal) {
        this.asset = newVal
      },
      deep: true,
    },
    asset: {
      handler: function () {
        this.$emit("changed", this.asset)
      },
      deep: true,
    },
  },
}
</script>

<style scoped>

</style>
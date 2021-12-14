<template>
  <fieldset class="form-fieldset pt-0">
    <div class="row text-muted mt-2 mb-2">
      <div class="col-lg-12 text-uppercase lh-base" style="font-size: 0.625rem">
        General
      </div>
    </div>
    <div class="row">
      <div class="col-lg-6">
        <div class="mb-3">
          <label class="form-label">Model</label>
          <input v-model="asset.ModelName" class="form-control" type="text">
        </div>
      </div>
      <div class="col-lg-6">
        <div class="mb-3">
          <label class="form-label">Internal ID</label>
          <input v-model="asset.InternalID" class="form-control" type="text">
        </div>
      </div>
    </div>
    <div class="row">
      <div class="col-lg-12">
        <div class="mb-3">
          <label class="form-label">Location</label>
          <input v-model="asset.Location" class="form-control" type="text">
        </div>
      </div>
    </div>
  </fieldset>
  <fieldset class="form-fieldset pt-0">
    <div class="row text-muted mt-2 mb-2">
      <div class="col-lg-12 text-uppercase lh-base" style="font-size: 0.625rem">
        Networking
      </div>
    </div>
    <div class="row">
      <div class="col-lg-6">
        <div class="mb-3">
          <label class="form-label">MAC</label>
          <Field v-slot="{field, errors}" v-model="asset.MAC" name="mac-address"
                 type="text">
            <input :class="[{'is-invalid': errors.length > 0},'form-control']" :readonly="!$store.getters.isAdmin"
                   v-bind="field">
          </Field>
          <ErrorMessage class="invalid-feedback" name="mac-address"/>
        </div>
      </div>
      <div class="col-lg-6">
        <div class="mb-3">
          <label class="form-label">IP</label>
          <Field v-slot="{field, errors}" v-model="asset.IP"
                 name="ip" type="text">
            <input :class="[{'is-invalid': errors.length > 0},'form-control']" :readonly="!$store.getters.isAdmin"
                   v-bind="field">
          </Field>
          <ErrorMessage class="invalid-feedback" name="ip"/>
        </div>
      </div>
    </div>
  </fieldset>
  <div class="row">
    <div class="mb-3">
      <div class="form-label">Installed Software</div>
      <select v-model="installedSoftware" :readonly="!$store.getters.isAdmin" class="form-control" multiple>
        <option v-for="(software, index) in software" :key="index" :value="software.ID">{{ software.Name }}
          {{ software.SoftwareAsset.Version }}
        </option>
      </select>
    </div>
  </div>
  <div class="row">
    <div class="col-lg-12">
      <div>
        <label class="form-label">Warranty info</label>
        <textarea v-model="asset.WarrantyInfo" :readonly="!$store.getters.isAdmin" class="form-control"
                  rows="3"></textarea>
      </div>
    </div>
  </div>
</template>

<script>
import {HardwareAsset} from "@/models";
import {ErrorMessage, Field} from "vee-validate";

export default {
  name: "HardwareOptions",

  emits: ["changed"],

  components: {
    Field,
    ErrorMessage,
  },

  props: {
    initialAsset: {
      type: HardwareAsset,
      required: true,
    }
  },

  data() {
    return {
      asset: new HardwareAsset({
        InstalledSoftware: new Array(),
      }),

      software: [],
    }
  },

  created() {
    this.asset = this.initialAsset
  },

  async mounted() {
    this.software = await window.go.sqlite.assetRepository.GetAllSoftware()
  },

  watch: {
    initialAsset: {
      handler: async function (newVal) {
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

  computed: {
    installedSoftware: {
      get: function () {
        if (!_.has(this.asset, "InstalledSoftware") || this.asset.InstalledSoftware === null) {
          return []
        }

        return this.asset.InstalledSoftware.map(s => s.ID)
      },

      set: function (ids) {
        this.asset.InstalledSoftware = this.software.filter(s => ids.indexOf(s.ID) !== -1)
      }
    },
  }
}
</script>

<style scoped>

</style>
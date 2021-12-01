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
          <Field v-slot="{field, errors}" v-model="asset.MAC" :rules="rules.macAddress" name="mac-address"
                 type="text">
            <input :class="[{'is-invalid': errors.length > 0},'form-control']" v-bind="field">
          </Field>
          <ErrorMessage class="invalid-feedback" name="mac-address"/>
        </div>
      </div>
      <div class="col-lg-6">
        <div class="mb-3">
          <label class="form-label">IP</label>
          <input v-model="asset.IP" class="form-control" type="text">
        </div>
      </div>
    </div>
  </fieldset>
  <div class="row">
    <div class="col-lg-12">
      <div>
        <label class="form-label">Warranty info</label>
        <textarea v-model="asset.WarrantyInfo" class="form-control" rows="3"></textarea>
      </div>
    </div>
  </div>
</template>

<script>
import {HardwareAsset} from "@/models";

import * as yup from 'yup';
import {Field, ErrorMessage} from "vee-validate";

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
      asset: null,
      rules: {
        macAddress: yup.string()
            .required()
            .matches(/^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$/,
                "Not a valid MAC address (AA:BB:CC:DD:EE:FF)"),
      }
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
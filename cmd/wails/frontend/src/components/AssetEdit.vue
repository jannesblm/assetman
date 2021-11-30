<template>
  <div class="modal modal-blur fade">
    <div class="modal-dialog modal-lg" role="document">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title">Asset: {{ $get(asset, 'Name', 'New asset') }}</h5>
          <button aria-label="Close" class="btn-close" data-bs-dismiss="modal" type="button"></button>
        </div>
        <div class="modal-body">
          <div class="mb-3">
            <label class="form-label">Name</label>
            <input v-model="asset.Name" class="form-control" placeholder="Asset name" type="text">
          </div>
          <label class="form-label">Asset type</label>
          <div class="form-selectgroup-boxes row mb-3">
            <div class="col-lg-6  mb-3 mb-lg-0">
              <label class="form-selectgroup-item">
                <input v-model="asset.AssetType" checked class="form-selectgroup-input" name="asset-type" type="radio"
                       value="hardware" :disabled="asset.ID > 0">
                <span class="form-selectgroup-label d-flex align-items-center p-3">
              <span class="me-3">
                <span class="form-selectgroup-check"></span>
              </span>
              <span class="form-selectgroup-label-content">
                <span class="form-selectgroup-title strong mb-1">Hardware</span>
              </span>
            </span>
              </label>
            </div>
            <div class="col-lg-6">
              <label class="form-selectgroup-item">
                <input v-model="asset.AssetType" class="form-selectgroup-input" name="asset-type" type="radio"
                       value="software" :disabled="asset.ID > 0">
                <span class="form-selectgroup-label d-flex align-items-center p-3">
              <span class="me-3">
                <span class="form-selectgroup-check"></span>
              </span>
              <span class="form-selectgroup-label-content">
                <span class="form-selectgroup-title strong mb-1">Software</span>
              </span>
            </span>
              </label>
            </div>
          </div>
          <div class="row">
            <div class="col-lg-6">
              <div class="mb-3">
                <label class="form-label">Purchased on</label>
                <div class="input-icon">
                  <span class="input-icon-addon">
                    <i class="ti ti-calendar"></i>
                  </span>
                  <input id="asset-purchase-date" :value="asset?.PurchaseDate" class="form-control"
                         placeholder="Select a date">
                </div>
              </div>
            </div>
            <div class="col-lg-6">
              <div class="mb-3">
                <label class="form-label">Manufacturer</label>
                <input id="asset-manufacturer" class="form-control" list="manufacturer-list" type="text" v-model="asset.Manufacturer.Name">
                <datalist v-for="(manufacturer, key) in $store.state.manufacturers" id="manufacturer-list" :key="key">
                  <option :data-value="manufacturer.ID" :value="manufacturer.Name"/>
                </datalist>
              </div>
            </div>
          </div>
        </div>
        <div class="modal-body">
          <SoftwareOptions v-if="asset.AssetType === 'software'"/>
          <HardwareOptions v-else/>
        </div>
        <div class="modal-body">
          <div class="col-lg-12">
            <div>
              <label class="form-label">Notes</label>
              <textarea class="form-control" rows="3"></textarea>
            </div>
          </div>
        </div>
        <div class="modal-footer">
          <a class="btn btn-link link-secondary" data-bs-dismiss="modal" href="#">
            Cancel
          </a>
          <a class="btn btn-primary ms-auto" data-bs-dismiss="modal" href="#">
            Create new report
          </a>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {Litepicker} from 'litepicker'
import {Asset, HardwareAsset, SoftwareAsset} from "@/models";
import {AssetDto} from "@/models.dto";
import {Form} from "vee-validate";

import SoftwareOptions from "@/components/SoftwareOptions";
import HardwareOptions from "@/components/HardwareOptions";
import dayjs from "dayjs";

export default {
  name: "AssetEdit",
  components: {HardwareOptions, SoftwareOptions},
  props: {
    id: {
      type: Number,
      required: true,
      default: 0,
    },
    manufacturers: {
      type: Array,
      required: true,
    }
  },

  data() {
    return {
      asset: new Asset(),
    }
  },

  methods: {
    async loadAsset() {
      this.asset = await window.go.sqlite.repository.GetById(this.id)
      console.log(this.asset)
    }
  },

  watch: {
    id: function() {
      this.loadAsset()
    }
  },

  mounted() {
    if (this.id > 0) {
      this.loadAsset()
    }

    new Litepicker({
      element: $('#asset-purchase-date').get(0),
      buttonText: {
        previousMonth: `<i class="ti ti-chevron-left">`,
        nextMonth: `<i class="ti ti-chevron-right">`
      },
    });
  },
}
</script>

<style scoped>

</style>
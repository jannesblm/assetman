<template>
  <div class="modal-content">
    <div class="modal-header">
      <h5 class="modal-title">Asset: {{ asset.Name }}</h5>
      <button aria-label="Close" class="btn-close" type="button" @click="$emit('close', false)"></button>
    </div>
    <Form>
      <div class="modal-body">
        <div class="mb-3">
          <label class="form-label">Name</label>
          <input v-model="asset.Name" class="form-control" placeholder="Asset name" type="text">
        </div>
        <label class="form-label">Asset type</label>
        <div class="form-selectgroup-boxes row mb-3">
          <div class="col-lg-6  mb-3 mb-lg-0">
            <label class="form-selectgroup-item">
              <input v-model="asset.AssetType" :disabled="asset.ID > 0" checked class="form-selectgroup-input"
                     name="asset-type"
                     type="radio" value="hardware">
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
              <input v-model="asset.AssetType" :disabled="asset.ID > 0" class="form-selectgroup-input"
                     name="asset-type"
                     type="radio" value="software">
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
        <div class="mb-3">
          <label class="form-label">Description</label>
          <input v-model="asset.Description" class="form-control" type="text">
        </div>
        <div class="row">
          <div class="col-lg-6">
            <div class="mb-3">
              <label class="form-label">Purchased on</label>
              <div class="input-icon">
                  <span class="input-icon-addon">
                    <i class="ti ti-calendar"></i>
                  </span>
                <input id="asset-purchase-date" v-model="purchaseDate" class="form-control"
                       placeholder="Select a date">
              </div>
            </div>
          </div>
          <div class="col-lg-6">
            <div class="mb-3">
              <label class="form-label">Manufacturer</label>
              <input id="asset-manufacturer" v-model="asset.Manufacturer.Name"
                     autocomplete="off"
                     class="form-control"
                     list="manufacturer-list"
                     type="text">
              <datalist id="manufacturer-list">
                <option v-for="(manufacturer, key) in $store.state.manufacturers" :key="key"
                        :data-value="manufacturer.ID" :value="manufacturer.Name"/>
              </datalist>
            </div>
          </div>
        </div>
      </div>
      <div class="modal-body">
        <SoftwareOptions v-if="asset.AssetType === 'software'"
                         :initial-asset="asset.SoftwareAsset"
                         @changed="onChange"/>
        <HardwareOptions v-else
                         :initial-asset="asset.HardwareAsset"
                         @changed="onChange"/>
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
        <a class="btn btn-link link-secondary" @click="$emit('close', false)">
          Cancel
        </a>
        <button type="button" class="btn btn-primary ms-auto" @click="save($event)">
          Create new report
        </button>
      </div>
    </Form>
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

  components: {HardwareOptions, SoftwareOptions, Form},

  emits: ["message", "close"],

  props: {
    id: {
      type: Number,
      required: true,
      default: 0,
    },
  },

  data() {
    return {
      asset: new Asset({
        'AssetType': 'hardware',
        'HardwareAsset': new HardwareAsset(),
        'SoftwareAsset': new SoftwareAsset()
      }),
    }
  },

  methods: {
    async load() {
      if (this.id > 0) {
        this.asset = new Asset(await window.go.sqlite.assetRepository.GetById(this.id))
      } else {
        this.asset = new Asset({
          'AssetType': 'hardware',
          'HardwareAsset': new HardwareAsset(),
          'SoftwareAsset': new SoftwareAsset()
        })
      }
    },

    async save(e) {
      e.preventDefault()

      let error = null;

      try {
         await window.go.sqlite.assetRepository.Save(AssetDto.fromObject(this.asset))
      } catch (err) {
        error = err
      }

      this.$emit("close", true)
      this.$emit("message", "save", error, this.asset)
    },

    onChange(asset) {
      switch (Object.getPrototypeOf(asset)) {
        case SoftwareAsset.prototype:
          this.asset.SoftwareAsset = asset;
          break;
        case HardwareAsset.prototype:
          this.asset.HardwareAsset = asset;
          break;
      }
    }
  },

  watch: {
    id: function () {
      this.load()
    },

    'asset.Manufacturer.Name': function (newVal) {
      let id = $("#manufacturer-list option[value='" + newVal.toString() + "']").data('value')

      if (id) {
        this.asset.ManufacturerID = id
        this.asset.Manufacturer.ID = id
      }
    },
  },

  computed: {
    purchaseDate: {
      get: function () {
        return dayjs(this.asset.PurchaseDate).format("DD/MM/YYYY")
      },

      set: function (newValue) {
        this.asset.PurchaseDate = dayjs(newValue, "DD/MM/YYYY").toDate()
      }
    }
  },

  mounted() {
    this.load()

    new Litepicker({
      element: $('#asset-purchase-date').get(0),
      format: "DD/MM/YYYY",
      buttonText: {
        previousMonth: `<i class="ti ti-chevron-left">`,
        nextMonth: `<i class="ti ti-chevron-right">`
      },
    })
  },
}
</script>

<style scoped>

</style>
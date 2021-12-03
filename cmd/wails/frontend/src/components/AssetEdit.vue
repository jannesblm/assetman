<template>
  <div class="modal-content">
    <div class="modal-header">
      <h5 class="modal-title">Asset: {{ asset.Name }}</h5>
      <button aria-label="Close" class="btn-close" type="button" @click="$emit('close', false)"></button>
    </div>
    <Form>

      <div class="modal-body">
        <div class="row mb-3">
          <div class="col">
            <label class="form-label">Name</label>
            <input v-model="asset.Name" :readonly="!$store.getters.isAdmin" class="form-control" placeholder="Asset name"
                   type="text">
          </div>
          <div class="col">
            <label class="form-label">Type</label>
            <input v-model="asset.TypeName" :readonly="!$store.getters.isAdmin" class="form-control" placeholder="Type"
                   type="text">
          </div>
        </div>
        <label class="form-label">Asset type</label>
        <div class="form-selectgroup-boxes row mb-3">
          <div class="col-lg-6  mb-3 mb-lg-0">
            <label :class="[{disabled: asset.ID > 0 || !$store.getters.isAdmin}, 'form-selectgroup-item']">
              <input v-model="asset.AssetType" :disabled="asset.ID > 0 || !$store.getters.isAdmin" checked
                     class="form-selectgroup-input"
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
            <label :class="[{disabled: asset.ID > 0 || !$store.getters.isAdmin}, 'form-selectgroup-item']">
              <input v-model="asset.AssetType" :disabled="asset.ID > 0 || !$store.getters.isAdmin"
                     class="form-selectgroup-input"
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
          <input v-model="asset.Description" :readonly="!$store.getters.isAdmin" class="form-control" type="text">
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
                       placeholder="Select a date" :readonly="!$store.getters.isAdmin">
              </div>
            </div>
          </div>
          <div class="col-lg-6">
            <div class="mb-3">
              <label class="form-label">
                Manufacturer
                <span v-if="manufacturerNotFound" class="badge bg-blue-lt">Creating New</span>
              </label>
              <input id="asset-manufacturer" v-model="manufacturer"
                     :readonly="!$store.getters.isAdmin"
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
            <textarea v-model="asset.Note" :readonly="!$store.getters.isAdmin" class="form-control" rows="3"></textarea>
          </div>
        </div>
      </div>
      <div class="modal-footer">
        <a class="btn btn-link link-secondary" @click="$emit('close', false)">
          Cancel
        </a>
        <button :disabled="!$store.getters.isAdmin" class="btn btn-warning ms-auto" type="button"
                @click="save($event, true)">
          Save & Search NVD
        </button>
        <button :disabled="!$store.getters.isAdmin" class="btn btn-primary" type="button" @click="save($event)">
          Save
        </button>
      </div>
    </Form>
  </div>
</template>

<script>
import {Litepicker} from 'litepicker'
import {Asset, HardwareAsset, SoftwareAsset} from "@/models";
import {ModelDto} from "@/models.dto";
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

      manufacturerNotFound: false
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

    async save(e, searchNvd = false) {
      e.preventDefault()

      let error = null

      try {
        await window.go.sqlite.assetRepository.Save(ModelDto.fromObject(this.asset))
        await this.load()
      } catch (err) {
        error = err
      }

      this.$emit("close", true)
      this.$emit("message", "save", error, this.asset, searchNvd)
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
  },

  computed: {
    manufacturer: {
      get: function () {
        return this.asset.Manufacturer.Name
      },

      set: function (newVal) {
        if (newVal === "") {
          this.manufacturerNotFound = false
          return
        }

        let el = document.querySelector("#manufacturer-list option[value='"
            + newVal.replace(/["\\]/g, '\\$&') + "']");

        if (el) {
          this.manufacturerNotFound = false
          this.asset.ManufacturerID = parseInt(el.dataset.value)
          this.asset.Manufacturer.Name = el.value
          this.asset.Manufacturer.ID = parseInt(el.dataset.value)
        } else {
          this.manufacturerNotFound = true
          this.asset.ManufacturerID = 0
          this.asset.Manufacturer.ID = 0
          this.asset.Manufacturer.Name = newVal
        }
      }
    },

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

    this.$store.dispatch("syncManufacturers")

    if (this.$store.getters.isAdmin) {
      new Litepicker({
        element: document.getElementById('asset-purchase-date'),
        format: "DD/MM/YYYY",
        buttonText: {
          previousMonth: `<i class="ti ti-chevron-left">`,
          nextMonth: `<i class="ti ti-chevron-right">`
        },
      })
    }
  },
}
</script>

<style scoped>

</style>
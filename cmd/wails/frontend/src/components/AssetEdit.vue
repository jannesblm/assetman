<template>
  <div class="modal-content">
    <div class="modal-header">
      <h5 class="modal-title">Asset: {{ asset.Name }}</h5>
      <button aria-label="Close" class="btn-close" type="button" @click="$emit('close', false)"></button>
    </div>
    <Form ref="form" :validation-schema="schema" @keydown.enter="save($event)">
      <div class="modal-body">
        <div class="row mb-3">
          <div class="col">
            <label class="form-label">Name</label>
            <Field v-slot="{field, errors}" v-model="asset.Name" name="name" type="text">
              <input :class="[{'is-invalid': errors.length > 0}, 'form-control']" :readonly="!$store.getters.isAdmin"
                     placeholder="Asset name" v-bind="field">
            </Field>
            <ErrorMessage class="invalid-feedback" name="name"/>
          </div>
          <div class="col">
            <label class="form-label">Type</label>
            <Field v-slot="{field, errors}" v-model="asset.TypeName" name="type-name" type="text">
              <input :class="[{'is-invalid': errors.length > 0}, 'form-control']" :readonly="!$store.getters.isAdmin"
                     placeholder="Type" v-bind="field">
            </Field>
            <ErrorMessage class="invalid-feedback" name="type-name"/>
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
          <Field v-slot="{field, errors}" v-model="asset.Description" name="description" type="text">
            <input :class="[{'is-invalid': errors.length > 0}, 'form-control']" :readonly="!$store.getters.isAdmin"
                   v-bind="field">
          </Field>
          <ErrorMessage class="invalid-feedback" name="description"/>
        </div>
        <div class="row">
          <div class="col-lg-6">
            <div class="mb-3">
              <label class="form-label">Purchased on</label>
              <div class="input-icon">
                  <span class="input-icon-addon">
                    <i class="ti ti-calendar"></i>
                  </span>
                <input id="asset-purchase-date" v-model="purchaseDate" :readonly="!$store.getters.isAdmin"
                       class="form-control" placeholder="Select a date">
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
import {ErrorMessage, Field, Form} from "vee-validate";
import {object, string} from 'yup';
import {markRaw} from "vue";

import SoftwareOptions from "@/components/SoftwareOptions";
import HardwareOptions from "@/components/HardwareOptions";
import dayjs from "dayjs";

export default {
  name: "AssetEdit",

  components: {
    HardwareOptions,
    SoftwareOptions,
    Form,
    Field,
    ErrorMessage
  },

  emits: ["message", "close"],

  props: {
    id: {
      type: Number,
      required: true,
      default: 0,
    },

    selectType: {
      type: String,
      required: false,
      default() {
        return "hardware"
      }
    }
  },

  data() {
    return {
      asset: new Asset({
        'AssetType': 'hardware',
        'HardwareAsset': new HardwareAsset(),
        'SoftwareAsset': new SoftwareAsset()
      }),

      schema: markRaw(object({
        'name': string().label('Asset Name').required().min(3),
        'type-name': string().label('Type').required(),
        'description': string().label('Description').required(),
        'mac-address': string().matches(/^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$/, {
          message: "Not a valid MAC address (AA:BB:CC:DD:EE:FF)",
          excludeEmptyString: true
        }),
        'ip': string().matches(/((^\s*((([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]))\s*$)|(^\s*((([0-9A-Fa-f]{1,4}:){7}([0-9A-Fa-f]{1,4}|:))|(([0-9A-Fa-f]{1,4}:){6}(:[0-9A-Fa-f]{1,4}|((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){1,2})|:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){1,3})|((:[0-9A-Fa-f]{1,4})?:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){1,4})|((:[0-9A-Fa-f]{1,4}){0,2}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){1,5})|((:[0-9A-Fa-f]{1,4}){0,3}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){1,6})|((:[0-9A-Fa-f]{1,4}){0,4}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(:(((:[0-9A-Fa-f]{1,4}){1,7})|((:[0-9A-Fa-f]{1,4}){0,5}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:)))(%.+)?\s*$))/, {
          message: "Not a valid IPv4/v6 address",
          excludeEmptyString: true
        })
      })),

      manufacturerNotFound: false
    }
  },

  methods: {
    async load() {
      if (this.asset.ID > 0) {
        this.asset = new Asset(await window.go.sqlite.assetRepository.GetById(this.asset.ID))
      } else {
        this.asset = new Asset({
          'AssetType': this.selectType,
          'HardwareAsset': new HardwareAsset(),
          'SoftwareAsset': new SoftwareAsset()
        })
      }
    },

    async save(e, searchNvd = false) {
      const {valid, errors} = await this.$refs.form.validate()

      if (!valid) {
        // Scroll to first error element if validation fails
        document.getElementsByClassName('modal')[0].scrollTo({
          top: document.getElementsByName(Object.keys(errors)[0])[0].offsetTop,
          behavior: 'smooth'
        });

        return
      }

      let error = null

      try {
        this.asset.ID = await window.go.sqlite.assetRepository.Save(ModelDto.fromObject(this.asset))
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
    id: function (newVal) {
      this.asset.ID = newVal
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
    this.asset.ID = this.id

    this.load()
    this.$store.dispatch("syncManufacturers")

    if (this.$store.getters.isAdmin) {
      new Litepicker({
        element: document.getElementById('asset-purchase-date'),
        dropdowns: {
          minYear: 1980,
          months: true,
          years: true,
        },
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
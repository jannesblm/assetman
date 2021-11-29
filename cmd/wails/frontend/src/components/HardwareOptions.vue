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
          <input class="form-control" type="text">
        </div>
      </div>
      <div class="col-lg-6">
        <div class="mb-3">
          <label class="form-label">Internal ID</label>
          <input class="form-control" type="text">
        </div>
      </div>
    </div>
    <div class="row">
      <div class="col-lg-12">
        <div class="mb-3">
          <label class="form-label">Location</label>
          <input class="form-control" type="text">
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
          <vue-mask v-model="MAC" :options="{placeholder: '__:__:__:__:__:__'}" autocomplete="off"
                    class="form-control" mask="00:00:00:00:00:00"/>
        </div>
      </div>
      <div class="col-lg-6">
        <div class="mb-3">
          <label class="form-label">IP</label>
          <input class="form-control" type="text">
        </div>
      </div>
    </div>
  </fieldset>
  <div class="row">
    <div class="col-lg-12">
      <div>
        <label class="form-label">Warranty info</label>
        <textarea class="form-control" rows="3"></textarea>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "HardwareOptions",

  data() {
    return {
      MAC: '',
    }
  }
}
</script>

<style scoped>

</style>
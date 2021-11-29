<template>
  <div class="row">
    <div class="col-lg-6">
      <div class="mb-3">
        <label class="form-label">License type</label>
        <input class="form-control" type="text" list="license-types">
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
        <input class="form-control" type="text">
      </div>
    </div>
  </div>
  <div class="row">
    <div class="col-lg-6">
      <div class="mb-3">
        <label class="form-label">Version</label>
        <input class="form-control" type="text">
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "SoftwareOptions"
}
</script>

<style scoped>

</style>
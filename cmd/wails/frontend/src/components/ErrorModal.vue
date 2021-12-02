<template>
  <div class="modal-content">
    <button aria-label="Close" class="btn-close" type="button" @click="$emit('close', false)"></button>
    <div :class="[{'bg-success': opts.success, 'bg-danger': ! opts.success}, 'modal-status']"></div>
    <div class="modal-body text-center py-4">
      <svg v-if="opts.success" class="icon mb-2 text-green icon-lg" fill="none" height="24" stroke="currentColor"
           stroke-linecap="round" stroke-linejoin="round" stroke-width="2" viewBox="0 0 24 24" width="24"
           xmlns="http://www.w3.org/2000/svg">
        <path d="M0 0h24v24H0z" fill="none" stroke="none"></path>
        <circle cx="12" cy="12" r="9"></circle>
        <path d="M9 12l2 2l4 -4"></path>
      </svg>
      <svg v-else class="icon mb-2 text-danger icon-lg" fill="none" height="24" stroke="currentColor"
           stroke-linecap="round" stroke-linejoin="round" stroke-width="2" viewBox="0 0 24 24" width="24"
           xmlns="http://www.w3.org/2000/svg">
        <path d="M0 0h24v24H0z" fill="none" stroke="none"></path>
        <path d="M12 9v2m0 4v.01"></path>
        <path
            d="M5 19h14a2 2 0 0 0 1.84 -2.75l-7.1 -12.25a2 2 0 0 0 -3.5 0l-7.1 12.25a2 2 0 0 0 1.75 2.75"></path>
      </svg>
      <h3>{{ opts.title }}</h3>
      <div class="text-muted">
        {{ opts.description }}
      </div>
    </div>
    <div class="modal-footer">
      <div class="w-100">
        <div class="row">
          <div v-if="opts.showCancel" class="col">
            <a class="btn btn-white w-100" @click="$emit('close', false)">
              {{ opts.cancelText }}
            </a>
          </div>
          <div v-if="opts.showConfirm" class="col">
            <a :class="[{'btn-success': opts.success, 'btn-danger': !opts.success}, 'btn', 'w-100']"
               @click="$emit('close', true)">
              {{ opts.confirmText }}
            </a></div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "ErrorModal",

  emits: ["close"],

  props: {
    opts: {
      type: Object,
      required: true,
      default() {
        return {
          success: true,
          title: "Title",
          description: "Description",
          cancelText: "Cancel",
          showCancel: true,
          confirmText: "Confirm",
          showConfirm: true,
        }
      }
    },
  }
}
</script>

<style scoped>

</style>
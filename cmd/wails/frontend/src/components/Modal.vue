<template>
  <transition name="fade">
    <div v-if="show" class="modal modal-blur" role="dialog"
         style="display: block">
      <div :class="['modal-dialog', ...classes]" role="document">
        <slot/>
      </div>
    </div>
  </transition>
  <transition name="fade">
    <div v-if="show" class="modal-backdrop fade show"></div>
  </transition>
</template>

<script>
export default {
  name: "Modal",

  props: {
    show: {
      type: Boolean,
      required: true,
      default: false,
    },

    classes: {
      type: Array,
      required: false,
      default() {
        return ['modal-sm', 'modal-dialog-centered']
      }
    }
  },
}
</script>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s ease;
}

.fade-enter-from,
.fade-leave-to {
  transition: opacity 0.1s ease;
  opacity: 0;
}
</style>
<template>
  <div class="wrapper">
    <Navigation v-if="$route.name !== 'login' && $store.state.user.isLoggedIn"/>
    <div class="page-wrapper">
      <router-view/>
    </div>
  </div>

  <Modal ref="modal" :show="$store.state.modal.show" :classes="$store.state.modal.classes">
    <template v-slot:default>
      <component :is="$store.state.modal.component"
                 v-bind="$store.state.modal.props"
                 @close="closeModal"
                 @message="handleMessage"
      ></component>
    </template>
  </Modal>
</template>

<script>

import Navigation from "@/components/Navigation";
import Modal from "@/components/Modal";

export default {
  name: 'App',

  components: {Modal, Navigation},

  methods: {
    closeModal(status) {
      this.$store.dispatch("hideModal", status)
    },

    handleMessage(type, ...args) {
      this.$store.dispatch("handleModalMessage", {
        type: type,
        args: args,
      })
    }
  }
}
</script>

<style>

</style>
<template>
  <div class="wrapper">
    <Navigation/>
    <div class="page-wrapper">
        <router-view/>
    </div>
  </div>
</template>

<script>

import Navigation from "@/components/Navigation";

export default {
  name: 'App',
  components: {Navigation},
}
</script>

<style>

</style>
<template>
  <aside class="navbar navbar-vertical navbar-expand-lg navbar-dark">
    <div class="container-fluid">
      <button class="navbar-toggler" data-bs-target="#navbar-menu" data-bs-toggle="collapse" type="button">
        <span class="navbar-toggler-icon"></span>
      </button>
      <h1 class="navbar-brand navbar-brand-autodark">
        <router-link to="/">
          AssetMan
        </router-link>
      </h1>
      <div id="navbar-menu" class="collapse navbar-collapse">
        <ul class="navbar-nav pt-lg-3">
          <li class="nav-item">
            <router-link class="nav-link" href="./index.html" to="/">
              <span class="nav-link-icon d-md-none d-lg-inline-block">
                <i class="ti ti-layout-2"></i>
              </span>
              <span class="nav-link-title">
                    Home
              </span>
            </router-link>
          </li>
          <li class="nav-item dropdown">
            <a class="nav-link dropdown-toggle" data-bs-auto-close="false"
               data-bs-toggle="dropdown" href="#" role="button">
              <span class="nav-link-icon d-md-none d-lg-inline-block">
                <i class="ti ti-box"></i>
              </span>
              <span class="nav-link-title">
                Assets
              </span>
            </a>
            <div class="dropdown-menu" data-bs-popper="none">
              <router-link to="/asset/list" class="dropdown-item">
                List
              </router-link>
            </div>
          </li>
        </ul>
      </div>
    </div>
  </aside>
</template>

<script>
export default {
  name: "Navigation"
}
</script>

<style scoped>

</style>
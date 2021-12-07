<template>
  <aside :class="[{'navbar-dark': darkMode, 'navbar-transparent': !darkMode}, 'navbar', 'navbar-vertical', 'navbar-expand-lg']">
    <div class="container-fluid">
      <button class="navbar-toggler" data-bs-target="#navbar-menu" data-bs-toggle="collapse" type="button">
        <span class="navbar-toggler-icon"></span>
      </button>
      <h1 class="navbar-brand navbar-brand-autodark">
        <router-link to="/asset/list/hardware">
          <img class="logo" src="scottishglen.png" alt="logo"/>
        </router-link>
      </h1>
      <div id="navbar-menu" class="collapse navbar-collapse">
        <ul class="navbar-nav pt-lg-3">
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
              <router-link
                  :class="[{'active': $route.name === 'assets' && $route.params.type === 'hardware'}, 'dropdown-item']"
                  to="/asset/list/hardware">
                Hardware
              </router-link>
              <router-link
                  :class="[{'active': $route.name === 'assets' && $route.params.type === 'software'}, 'dropdown-item']"
                  to="/asset/list/software">
                Software
              </router-link>
            </div>
          </li>
          <li class="nav-item" v-if="$store.getters.isAdmin">
            <router-link :class="[{'active': $route.name === 'manufacturers'}, 'nav-link']"
                         to="/manufacturers">
              <span class="nav-link-icon d-md-none d-lg-inline-block">
                <i class="ti ti-users"></i>
              </span>
              <span class="nav-link-title">
                  Manufacturers
              </span>
            </router-link>
          </li>
          <li class="nav-item" v-if="$store.getters.isAdmin">
            <router-link :class="[{'active': $route.name === 'backup'}, 'nav-link']"
                         to="/backup">
              <span class="nav-link-icon d-md-none d-lg-inline-block">
                <i class="ti ti-download"></i>
              </span>
              <span class="nav-link-title">
                  Backup
              </span>
            </router-link>
          </li>
          <li class="nav-item">
            <a class="nav-link cursor-pointer" @click="logout">
              <span class="nav-link-icon d-md-none d-lg-inline-block">
                <i class="ti ti-logout"></i>
              </span>
              <span class="nav-link-title">
                  Logout
              </span>
            </a>
          </li>
          <li class="flex-grow-1"></li>
          <li style="padding: 0.5rem 1.5rem">
            Logged in as: {{ $store.state.user.Name }}
          </li>
          <li style="padding: 0.5rem 1.5rem">
            <label class="form-check form-switch">
              <input class="form-check-input cursor-pointer" type="checkbox" v-model="darkMode">
              <span class="form-check-label">Dark Mode</span>
            </label>
          </li>
        </ul>
      </div>
    </div>
  </aside>
</template>

<script>
export default {
  name: "Navigation",

  data() {
    return {
      darkMode: false,
    }
  },

  methods: {
    setDarkMode: function() {
      if (this.darkMode) {
        document.body.classList.add('theme-dark');
      } else {
        document.body.classList.remove('theme-dark');
      }
    },

    logout: function () {
      this.$confirm("Confirm", "Do you want to logout from the system?", (yes) => {
        if (yes) {
          window.go.auth.service.Logout()
          this.$router.replace({name: "login"})
        }
      })
    }
  },

  watch: {
    darkMode: function () {
      this.setDarkMode()
    }
  },

  mounted() {
    this.setDarkMode()
  }
}
</script>

<style scoped>
@media (max-width: 991.98px) {
  .logo {
    width: 100px;
  }
}

</style>
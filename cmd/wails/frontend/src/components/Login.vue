<template>
  <div class="page page-center">
    <div class="container-tight py-4">
      <div class="text-center mb-4">
        <img alt="logo" class="logo" src="scottishglen.png"/>
      </div>
      <form action="." autocomplete="off" class="card card-md" @keydown.enter="login">
        <div class="card-body">
          <h2 class="card-title text-center mb-4">Login to your account</h2>
          <div class="mb-3">
            <label class="form-label">User</label>
            <input v-model="user" class="form-control" placeholder="Username" type="text">
          </div>
          <div class="mb-2">
            <label class="form-label">Password</label>
            <input v-model="password" autocomplete="off" class="form-control" placeholder="Password" type="password">
          </div>
          <div class="form-footer">
            <button class="btn btn-primary w-100" type="button" @click="login">Sign in</button>
          </div>
        </div>
      </form>
    </div>
  </div>

</template>

<script>
export default {
  name: "Login",

  data() {
    return {
      user: '',
      password: '',
    }
  },

  methods: {
    async login() {
      try {
        const user = await window.go.auth.service.Authenticate(this.user, this.password)
        this.$store.dispatch("setUser", user)

        await this.$router.replace({ name: "home" })
      } catch (error) {
        this.$showDialog("Login failed",
            "Please check username and password and try again (" + error + ")")
      }
    }
  }
}
</script>

<style scoped>
.logo {
  width: 200px;
}
</style>
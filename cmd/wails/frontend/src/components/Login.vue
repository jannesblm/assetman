<template>
  <div class="page page-center">
    <div class="container-tight py-4">
      <div class="text-center mb-4">
        AssetMan
      </div>
      <form action="." autocomplete="off" class="card card-md" method="get">
        <div class="card-body">
          <h2 class="card-title text-center mb-4">Login to your account</h2>
          <div class="mb-3">
            <label class="form-label">User</label>
            <input class="form-control" placeholder="Username" type="email">
          </div>
          <div class="mb-2">
            <label class="form-label">Password</label>
            <input autocomplete="off" class="form-control" placeholder="Password" type="password">
          </div>
          <div class="form-footer">
            <button class="btn btn-primary w-100" type="submit">Sign in</button>
          </div>
        </div>
      </form>
    </div>
  </div>

</template>

<script>
export default {
  name: "Login"
}
</script>

<style scoped>

</style>
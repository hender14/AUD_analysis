<template>
  <nav class="navbar navbar-expand-md navbar-light bg-light py-3 border-bottom">
    <div class="container-fluid">
        <router-link to="/" class="nav-link navbar-brand fs-4 text-decoration-underline fw-bold">AUD analysis</router-link>
        <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
          <span class="navbar-toggler-icon"></span>
        </button>
    </div>

    <div class="collapse navbar-collapse" id="navbarNav">
      <ul class="navbar-nav ms-auto">
        <!-- ﾛｸﾞｲﾝ済みなら表示 -->
        <template v-if="auth">
          <li class="nav-item bg-dark px-2 me-4 rounded">
            <router-link to="/logout" class="nav-link nav-link text-light active">Logout</router-link>
          </li>
        </template>
        <!-- 未ﾛｸﾞｲﾝなら非表示 -->
        <template v-if="!auth">
          <li class="nav-item bg-dark px-2 me-4 rounded">
            <router-link to="/login" class="nav-link text-light active">Login</router-link>
          </li>
          <li class="nav-item bg-dark me-5 rounded">
            <router-link to="/register" class="nav-link text-light active">Register</router-link>
          </li>
        </template>
      </ul>
    </div>
  </nav>
</template>

<script>
import { computed } from 'vue'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'
import { userURL } from '@/main'
import axios from 'axios'

export default {
  name: "Nav",
  setup() {
    const store = useStore()
    const router = useRouter()
    const auth = computed(() => store.state.auth)
    const logout = async () => {
      await axios.get( userURL + 'logout', {})
      // ｴﾗｰ処理を追加する必要あり
      store.dispatch('setAuth', false)
      await router.push('/login')
    }
    return {
      auth,
      logout
    }
  }
}
</script>

<style>

</style>
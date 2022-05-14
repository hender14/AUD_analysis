<template>
  <div v-if="notify.cls" :class="`alert alert-${notify.cls}`" role="alert">
    {{ notify.message }}
  </div>
  <h1 class="h3 mb-3 fw-normal">Logout</h1>
  <p>ログアウトしてもよろしいでしょうか。</p>
  <div class="border-bottom d-md-flex flex-col justify-content-md-center">
  <button type="button" class="btn btn-dark btn-sm me-5 my-4">
    <router-link to="/login" class="nav-link text-light" @click="logout">Yes</router-link>
  </button>
    <button type="button" class="btn btn-dark btn-sm ms-5 my-4">
      <router-link to="/" class="nav-link text-light active">No</router-link>
    </button>
  </div>
</template>

<script>
import { reactive, onMounted } from 'vue'
import { useStore } from 'vuex'
import axios from 'axios'
import { useRouter } from 'vue-router'
import { userURL } from '@/main'

export default {
  name: "Logout",
  setup() {
    const router = useRouter()
    const store = useStore()
    const notify = reactive({
      cls: '',
      message: ''
    })

    const logout = async () => {
      try {
        await axios.get( userURL + 'logout', {})
        await store.dispatch('setAuth', false)
        await router.push('/login')
      } catch(e) {
        if (axios.isAxiosError(e) && e.response.status === 400) {
          notify.cls = 'danger'
          notify.message = e.response.data["message"]
        }
      }
    }

    onMounted(async () => {
      try {
        await axios.get( userURL + 'user')
        // actionsに設定したﾊﾟﾗﾒｰﾀ名を設定
        await store.dispatch('setAuth', true)
      } catch(e) {
        await store.dispatch('setAuth', false)
        router.push('/login')
      }
    })

    return {
      logout,
      notify
    }
  }
}
</script>

<style>
.form-login {
  width: 100%;
  max-width: 330px;
  padding: 15px;
  margin: auto;
}

.form-login .form-control {
  position: relative;
  box-sizing: border-box;
  height: auto;
  padding: 10px;
  font-size: 16px;
}
.form-login .form-control:focus {
  z-index: 2;
}
.form-login input[type="email"] {
  margin-bottom: -1px;
  border-bottom-right-radius: 0;
  border-bottom-left-radius: 0;
}
.form-login input[type="password"] {
  margin-bottom: 10px;
  border-top-left-radius: 0;
  border-top-right-radius: 0;
}

</style>
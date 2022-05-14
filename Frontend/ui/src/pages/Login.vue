<template>
  <main class="form-login">
    <form @submit.prevent="login">
      <div v-if="notify.cls" :class="`alert alert-${notify.cls}`" role="alert">
        {{ notify.message }}
      </div>
      <h1 class="h3 mb-3 fw-normal">Login</h1>
      <!-- <p>{{ message }}</p> -->
      <input
        v-model="email"
        type="email"
        class="form-control"
        placeholder="Email"
        required>

      <input
        v-model="password"
        type="password"
        class="form-control"
        placeholder="Password"
        required>

      <div class="mb-2">
        <router-link to="/forgot">Forgot Password?</router-link>
      </div>

      <button
        class="w-100 btn btn-lg btn-primary"
        type="submit">Login</button>
    </form>
  </main>
  <div class="container">
    <div class="d-flex mt-3 border-top text-center">
      <router-link to="/register">アカウントを作成する</router-link>
    </div>
  </div>
</template>

<script>
// import { reactive,ref } from 'vue'
import { reactive, ref } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'
import { userURL } from '@/main'

export default {
  name: "Login",
  setup() {
    const email = ref('')
    const password = ref('')
    const router = useRouter()
    var response = ref('')
    // var message = ref('')
    const notify = reactive({
      cls: '',
      message: ''
    })

    const login = async () => {
      try {
        response = await axios.post( userURL + 'login', {
          email: email.value,
          password: password.value
        })
        await router.push('/input')
        return response
      } catch (e) {
        // if (axios.isAxiosError(e) && e.response && e.response.status === 400) {
        if (axios.isAxiosError(e) && e.response.status === 400) {
          notify.cls = 'danger'
          notify.message = e.response.data["message"]
          // message.value = e.response.data["message"]
        }
      }
    }

    return {
      email,
      password,
      login,
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
<template>
  <main class="form-reset">
    <form @submit.prevent="submit">
      <div v-if="notify.cls" :class="`alert alert-${notify.cls}`" role="alert">
        {{ notify.message }}
      </div>
      <h1 class="h3 mb-3 fw-normal">Please Reset Your Password</h1>
      <input
        v-model="password"
        type="password"
        class="form-control"
        placeholder="Password"
        required>
      <input
        v-model="passwordConfirm"
        type="password"
        class="form-control"
        placeholder="Password Confirm"
        required>

      <button
        class="w-100 btn btn-lg btn-primary"
        type="submit">Submit</button>
    </form>
  </main>
</template>

<script>
import { reactive, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import axios from 'axios'
import { userURL } from '@/main'

export default {
  name: "Reset",
  setup() {
    const password = ref('')
    const passwordConfirm = ref('')
    const route = useRoute()
    const router = useRouter()
    var response = ref('')
    const notify = reactive({
      cls: '',
      message: ''
    })

    const submit = async () => {
      try {
      response = await axios.post( userURL + 'reset', {
        // urlからtokenを取得
        token: route.params.token,
        password: password.value,
        password_confirm: passwordConfirm.value
      })
      // ﾛｸﾞｲﾝ画面へ遷移する
      await router.push('/login')
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
      password,
      passwordConfirm,
      submit,
      notify
    }
  }
}
</script>

<style>
.form-reset {
  width: 100%;
  max-width: 330px;
  padding: 15px;
  margin: auto;
}

.form-reset .form-control {
  position: relative;
  box-sizing: border-box;
  height: auto;
  padding: 10px;
  font-size: 16px;
}
.form-reset .form-control:focus {
  z-index: 2;
}
.form-reset input[type="email"] {
  margin-bottom: -1px;
  border-bottom-right-radius: 0;
  border-bottom-left-radius: 0;
}
.form-reset input[type="password"] {
  margin-bottom: 10px;
  border-top-left-radius: 0;
  border-top-right-radius: 0;
}

</style>
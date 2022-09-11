<template>
  <main class="form-register">
    <form @submit.prevent="submit">
      <div v-if="notify.cls" :class="`alert alert-${notify.cls}`" role="alert">
        {{ notify.message }}
      </div>
      <h1 class="h3 mb-3 fw-normal">Please Register</h1>
      <p>{{ message }}</p>
      <input
        v-model="firstName"
        class="form-control"
        placeholder="First Name"
        required>
      <input
        v-model="lastName"
        class="form-control"
        placeholder="Last Name"
        required>

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
        @change="validate_uploads"
        required>
        <p class="small text-secondary">{{ notify.password }}</p>
      <input
        v-model="passwordConfirm"
        type="password"
        class="form-control"
        placeholder="Password Confirm"
        required>

      <div class="checkbox">
        <section class="row">
          <label class="py-2">
            <input type="checkbox" required>
            <router-link to="/rule/teams" target= "_blank" class="ms-1">利用規約</router-link>
            に同意する
          </label>
          <label class="py-2">   
          <input type="checkbox" required>
            <router-link to="/rule/privacy" target= "_blank" class="ms-1">プライバシーポリシー</router-link>
              に同意する
          </label>
        </section>
      </div>

      <button
        class="w-100 btn btn-lg btn-primary"
        type="submit">登録する</button>
    </form>
  </main>
</template>

<script>
import { reactive, ref } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'
import { userURL } from '@/main'

export default {
  name: "Register",
  setup() {
    // v-modelからﾌｫｰﾑに入力された値を格納
    const firstName = ref('');
    const lastName = ref('');
    const email = ref('');
    const password = ref('');
    const passwordConfirm = ref('');
    const router = useRouter()
    var response = ref('')
    const notify = reactive({
      cls: '',
      message: '',
      password: 'passwordは英数字8文字以上16字以内が使えます'
    })
    const errors = ref('')

    const validate_uploads = async () => {
      errors.value = ''
      if (password.value.length < 8 || password.value.length > 16) {
        errors.value += 'passwordは8文字以上16字以内にしてください\n'
      }
      if (password.value.match(/[^A-Za-z0-9]+/)) {
        errors.value += 'passwordは英数字のみを使用してください\n'
      }
      if (errors.value) {
        notify.password = errors.value
      }
    }

    const submit = async () => {
      try {
        errors.value = ''
        if (password.value.length < 8 || password.value.length > 16) {
          errors.value += 'passwordは8文字以上16字以内にしてください\n'
        }
        if (password.value.match(/[^A-Za-z0-9]+/)) {
          errors.value += 'passwordは英数字のみを使用してください\n'
        }
        if (errors.value) {
        console.log(errors.value)
        //errorsが存在する場合は内容をalert
        alert(errors.value)
        //valueを空にしてﾘｾｯﾄする
        return
        }
        // Register apiへPOSTﾘｸｴｽﾄ
        response = await axios.post( userURL + 'register', {
          first_name: firstName.value,
          last_name: lastName.value,
          email: email.value,
          password: password.value,
          password_confirm: passwordConfirm.value,
        })
        // ﾛｸﾞｲﾝ画面に戻る
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
      firstName,
      lastName,
      email,
      password,
      passwordConfirm,
      submit, 
      notify,
      validate_uploads,
    }
  }
}
</script>

<style>
.form-register {
  width: 100%;
  max-width: 330px;
  padding: 15px;
  margin: auto;
}

.form-register .form-control {
  position: relative;
  box-sizing: border-box;
  height: auto;
  padding: 10px;
  font-size: 16px;
}
.form-register .form-control:focus {
  z-index: 2;
}
.form-register input[type="email"] {
  margin-bottom: -1px;
  border-bottom-right-radius: 0;
  border-bottom-left-radius: 0;
}
.form-register input[type="password"] {
  margin-bottom: 10px;
  border-top-left-radius: 0;
  border-top-right-radius: 0;
}

</style>
<template>
  <div class="container">
    <h1 class="h2 mb-4 border-bottom">お問い合わせ</h1>
  </div>
  <main class="form-login">
    <form @submit.prevent="contact">
      <div v-if="notify.cls" :class="`alert alert-${notify.cls}`" role="alert">
        {{ notify.message }}
      </div>
      <!-- <p>{{ message }}</p> -->
      <input
        v-model="email"
        type="email"
        class="form-control mb-5"
        placeholder="Email"
        required>

      <input
        v-model="title"
        class="form-control"
        placeholder="Title"
        maxlength="20"
        required>
        <p class="mb-3 d-flex flex-row-reverse">20文字以内</p>

      <textarea
        v-model="content"
        class="form-control"
        placeholder="お問い合わせ内容"
        style="height: 300px"
        rows="3"
        maxlength="200"
        required>
      </textarea>
      <p class="mb-5 d-flex flex-row-reverse">200文字以内</p>

      <button
        class="w-100 btn btn-lg btn-primary"
        type="submit">お問い合わせ</button>
    </form>
  </main>
</template>

<script>
import { reactive, ref } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'
import { userURL } from '@/main'

export default {
  name: "Contact",
  setup() {
    let user
    const email = ref('')
    const title = ref('')
    const content = ref('')
    const router = useRouter()
    // var message = ref('')
    const notify = reactive({
      cls: '',
      message: ''
    })

    const contact = async () => {
      try {
        const { data } = await axios.get( userURL + 'user')
        user = data.id
        await axios.post( userURL + 'contact', {
          id: user,
          email: email.value,
          title: title.value,
          content: content.value
        })
        await router.push('/')
        return
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
      user,
      email,
      title,
      content,
      contact,
      notify
    }
  }
}
</script>

<style>

</style>
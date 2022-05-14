<template>
  <!-- <h1>{{ message }}</h1> -->
    <div class="container mt-5">
      <section class="row">
        <section class="col-md-8">
          <header class="border-bottom pb-2 mb-3 d-flex align-items-center">
            <h1 class="fs-5 m-0">ユーザー情報</h1>
            <!-- <button type="button" class="btn btn-primary btn-sm ms-auto">修正</button> -->
          </header>
          <p>現在、登録しているユーザーに関する情報です。</p>
          <table class="table table-striped table-bordered">
            <thead>
              <tr>
                <th scope="col">Item</th>
                <th scope="col">ユーザー情報</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <th scope="row">First name</th>
                <td>{{ items.first_name }}</td>
              </tr>
              <tr>
                <th scope="row">Last name</th>
                <td>{{ items.last_name }}</td>
              </tr>
              <tr>
                <th scope="row">Email</th>
                <td>{{ items.email }}</td>
              </tr>
            </tbody>
          </table>
        </section>
      </section>
  </div>
  <div class="container">
    <div class="d-flex mt-3 border-top text-center">
    <!-- <div class="d-flex mt-3 border-top text-center"> -->
      <router-link to="/forgot">passwordを変更する</router-link>
    </div>
  </div>
</template>

<script>
import { userURL } from '../main.ts'
import { ref, onMounted } from 'vue'
import axios from 'axios'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'

export default {
  name: "User",
  setup() {
    const store = useStore()
    const router = useRouter()
    const items = ref('')

    onMounted(async () => {
      try {
        const { data } = await axios.get( userURL + 'user')
        items.value = data
        // actionsに設定したパラメータ名を設定
        await store.dispatch('setAuth', true)
      } catch(e) {
        await store.dispatch('setAuth', false)
        router.push('/login')
      }
    })

    return {
      items
    }
  }

}
</script>

<style>

</style>
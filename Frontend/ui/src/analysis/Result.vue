<template>
  <!-- <h1>{{ message }}</h1> -->
  <div v-if="notify.cls" :class="`alert alert-${notify.cls}`" role="alert">
  {{ notify.message }}
  </div>
  <header class="border-bottom pb-2 mb-3 d- flex align-items-center">
    <h1 class="fs-5 m-0">解析結果</h1>
  </header>
  <p>現在、保存されている解析結果になります。</p>
  <!-- <button type="button" class="btn btn-primary btn-sm ms-auto">削除</button> -->
  <table class="table table-striped table-bordered">
    <thead>
      <tr>
        <th scope="col">名前</th>
        <th scope="col">作成日</th>
      </tr>
    </thead>
    <tbody v-for="(name, index) in names" v-bind:key='(name)'>
      <tr>
        <div class="form-check">
          <!-- <input class="form-check-input" type="checkbox" value="" id="flexCheckIndeterminate"> -->
          <label class="form-check-label" for="flexCheckIndeterminate">
            <th scope="row">
              <router-link to="/detail" class="nav-link" @click="detail(name)">
              {{ name }}
              </router-link>
            </th>
          </label>
          </div>
            <td>
              {{ generation[index] }}
            </td>
      </tr>
    </tbody>
  </table>
</template>


<script>
import { reactive, ref, onMounted } from 'vue'
import axios from 'axios'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'
import { userURL, analysisURL } from '@/main'

export default {
  name: "Result",
  setup() {
    let user
    // const message = ref('')
    const names = ref('')
    const generation = ref('')
    localStorage.setItem('filename', 'None')
    // let ch = (localStorage.getItem('filename'))
    // console.log(ch)

    const store = useStore()
    const router = useRouter()
    const notify = reactive({
      cls: '',
      message: ''
    })

    onMounted(async () => {
      try {
        const { data } = await axios.get( userURL + 'user')
        user = data.ID
        // actionsに設定したパラメータ名を設定
        await store.dispatch('setAuth', true)
      } catch(e) {
        await store.dispatch('setAuth', false)
        router.push('/login')
      }
      
      try {
        await axios.get( analysisURL + 'list', { params: {
          username: user
        }}).then(res => {
          names.value = res['data']['name']
          let buffer = res['data']['generation']
          // for(var i=0; i <= 16; i++){
          //   buffer[i] = new Date(buffer[i])
          // }
          generation.value = buffer
        })
      } catch(e) {
        if (axios.isAxiosError(e) && e.response.status === 400) {
          notify.cls = 'danger'
          notify.message = e.response.data["message"]
        // message.value = e.response.data["message"]
        }
      }
    })

    const detail = async (filename) => {
      await router.push({name: 'Detail', params: {messag: filename}})
    }

    return {
      // message,
      names,
      generation,
      notify,
      detail
    }
  }

}
</script>
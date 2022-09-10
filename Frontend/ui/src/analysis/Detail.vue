<template>
  <!-- <h1>{{ message }}</h1> -->
  <div v-if="notify.cls" :class="`alert alert-${notify.cls}`" role="alert">
  {{ notify.message }}
  </div>
  <header class="border-bottom pb-2 mb-3 d- flex align-items-center">
    <h1 class="fs-5 m-0">解析結果詳細</h1>
  </header>
  <header class="border-bottom pb-2 mb-3 d- flex align-items-center">
    <h1 class="fs-5 m-0">ファイル名</h1>
    <h1>{{ messag }}</h1>
    <!-- <h1>{{ $route.params.messag }}</h1> -->
    <div class="row">
      <div class="col">
        <!-- <a href="#" class="list-group-item list-group-item-action active" aria-current="true">ｷｰﾜｰﾄﾞ解析</a> -->
        <a class="list-group-item list-group-item-action active" aria-current="true">ｷｰﾜｰﾄﾞ解析</a>
      </div>
      <div class="col">
        <!-- <a href="#" class="list-group-item list-group-item-action">文字解析</a> -->
        <a  class="list-group-item list-group-item-action">文字解析</a>
      </div>
    </div>
  </header>
  <header class="border-bottom pb-2 mb-3 d- flex align-items-center">
    <h1 class="fs-5 m-0">ｷｰﾜｰﾄﾞ名</h1>
    <div class="row">
      <div class="col">
        <a >ｷｰﾜｰﾄﾞ1 <br>{{ keylist.word1 }}</a>
      </div>
      <div class="col">
        <a >ｷｰﾜｰﾄﾞ1 <br>{{ keylist.word2 }}</a>
      </div>
      <div class="col">
        <a >ｷｰﾜｰﾄﾞ1 <br>{{ keylist.word3 }}</a>
      </div>
    </div>
  </header>
  <table class="table table-striped table-bordered">
      <thead>
        <tr>
          <th scope="col">ｷｰﾜｰﾄﾞ詳細</th>
          <th scope="col">タイトル</th>
          <th scope="col">再生数</th>
          <th scope="col">概要</th>
          <th scope="col">動画url</th>
          <th scope="col">サムネイル</th>
        </tr>
    </thead>
    <tbody v-for="(item) in items" v-bind:key='item'>
      <tr>
        <td>{{item.keyword}}</td>
        <td>
            {{item.title}}
        </td>
        <td>{{item.再生数}}</td>
        <td>{{item.概要}}</td>
        <td>{{item.動画url}}</td>
        <td><img :src="item.サムネイル" alt=""></td>
        <!-- <td><img alt="main image" class="img-sample" src= {{item.サムネイル}}></td> -->
      </tr>
    </tbody>
  </table>
</template>


<script>
import { userURL, analysisURL } from '../main.ts'
import { reactive, ref, onMounted } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'
// import { toRef, toRefs } from 'vue'

export default {
  name: "Detail",
  props: ['messag'],
  setup(props) {
    const router = useRouter()
    let user
    let buffer = props
    let filename = buffer.messag
    const notify = reactive({
      cls: '',
      message: ''
    })
    // const filename = toRefs(props)
    // const { messag } = toRef(props)
    // console.log(localStorage.getItem('filename'))
    // if (localStorage.getItem('filename') != 'None') {
    //   filename = localStorage.getItem('filename')
    // }
    // localStorage.setItem('filename', filename)
    const message = ref('You are not logged in!')
    const items = ref('')
    const keylist = reactive({
      word1: '',
      word2: '',
      word3: '',
    })

    const store = useStore()

    onMounted(async () => {
      try {
        const { data } = await axios.get( userURL + 'user' )
        user = data.ID
        // actionsに設定したパラメータ名を設定
        await store.dispatch('setAuth', true)
      } catch(e) {
        await store.dispatch('setAuth', false)
        router.push('/login')
      }
      try {
        await axios.get( analysisURL + 'detail', { params: {
          username: user,
          filename: filename
        }}).then(res => {
          console.log(res['data'])
          message.value = `get completed`
          items.value = res['data']
          keylist.word1 = res['data']['key000']['keyword']
          keylist.word2 = res['data']['key010']['keyword']
          keylist.word3 = res['data']['key020']['keyword']
        })
      } catch(e) {
        if (axios.isAxiosError(e) && e.response.status === 400) {
          notify.cls = 'danger'
          notify.message = e.response.data["message"]
          // message.value = e.response.data["message"]
        }
      }
    })

    return {
      message,
      items,
      notify,
      keylist
    }
  }

}
</script>
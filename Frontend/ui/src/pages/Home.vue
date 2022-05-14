<template>
  <div class="container">
    <div class="card" style="max-width: 640px;">
      <img alt="main image" class="img-responsive" src="../static/assets/top.jpg">
      <div class="card-img-overlay">
        <h1 class="text-white fs-1">AUD Analysis</h1>
        <p class="text-white">AUDは動画を効率的に作成するDXツールです。</p>
        <button type="button" class="btn btn-light btn-sm ms-auto my-4">
          <router-link to="/register" class="nav-link navbar-brand fs-5 text-dark fw-light">Let's Start</router-link>
        </button>
      </div>
    </div>

    <div class="container mt-5">
      <div class="border-bottom d-md-flex flex-col justify-content-md-between align-items-center">
        <div class="list-group fs-5 me-5">
          <p class="py-2 fs-6 fw-light">FUNCTION</p>
          <h1 class="py-2">AUDの基本機能</h1>
          <p class="fs-6 text-secondary">AUD analysisの結果から、動画のｷｰﾜｰﾄﾞの提示、NGﾜｰﾄﾞリスクの確認、またYoutubeにおいて顧客の注目を取得しやすいｷｰﾜｰﾄﾞを確認する事が出来ます。</p>
        </div>
        <div class="">
          <img alt="main image" class="img-sample" src="../static/assets/detail.jpeg" width="360" height="360">
        </div>
      </div>
      <div class="border-bottom d-md-flex flex-col justify-content-md-between align-items-center mt-5 pt-5">
        <div class="me-5">
          <img alt="main image" class="img-sample" src="../static/assets/result.jpeg" width="360" height="360">
        </div>
        <div class="list-group fs-5">
          <p class="py-2 fs-6 fw-light">MERIT</p>
          <h1 class="py-2">AUDで作業の生産性を向上</h1>
          <p class="fs-6 text-secondary">必要な操作は非常に容易で、動画をアップロードするだけ。動画をアップロードして頂ければ、Voicerが内容を解析して動画に関するアドバイスを提供します。</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { userURL } from '../main.ts'
import { onMounted } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'

export default {
  name: "Home",
  setup() {
    const store = useStore()
    const router = useRouter()

    onMounted(async () => {
      try {
        await axios.get( userURL + 'user')
        // actionsに設定したﾊﾟﾗﾒｰﾀ名を設定
        await store.dispatch('setAuth', true)
        router.push('/input')
      } catch(e) {
        await store.dispatch('setAuth', false)
      }
    })

    return {
    }
  }

}
</script>

<style>

/* 文字を中央に揃える */
.card-img-overlay{
  padding: 0;
  top: calc(30% - 0.5rem);
  text-align: center;
}

.img-responsive {
  font-family: Montserrat, sans-serif;
  position: relative;
  overflow: hidden;
  width: 100%;
  height: auto;
}

</style>
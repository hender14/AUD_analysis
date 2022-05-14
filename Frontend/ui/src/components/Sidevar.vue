<template>
  <!-- ﾛｸﾞｲﾝ済みなら表示 -->
  <template v-if="auth">
  </template>
  <!-- 未ﾛｸﾞｲﾝなら非表示 -->
  <template v-if="auth">
    <router-link to="/user" class="list-group-item list-group-item-action">
      ユーザー情報
    </router-link>
    <router-link to="/input" class="list-group-item list-group-item-action">
      音声解析
    </router-link>
    <router-link to="/result" class="list-group-item list-group-item-action">
      解析結果
    </router-link>
  </template>
</template>

<script>
import { computed } from 'vue'
import { useStore } from 'vuex'

export default {
  name: "Sidevar",
  setup() {
    const store = useStore()
    const auth = computed(() => store.state.auth)
    
    return {
      auth,
    }
  }
}
</script>

<style>

</style>
<template>
  <footer class="bg-secondary text-center text-light p-5 mt-5 contents-list-page">
  <!-- <footer class="bg-secondary text-center text-light p-5 mt-5"> -->
    <!-- <div class="footer-wrap flex"> -->
    <div class="footer-wrap d-flex">
      <div class="footer-contests-l">
        <ul class="footer-menu d-flex align-items-start list-unstyled">
        <!-- <ul class="footer-menu flex"> -->
          <li>
            <router-link to="/rule/teams" class="list-group-item-action">
            <!-- <router-link to="/input" class="list-group-item list-group-item-action"> -->
              利用規約
            </router-link>
            <span class="divide px-2">|</span>
          </li>
          <li>
            <router-link to="/rule/privacy" class="list-group-item-action">
            <!-- <router-link to="/input" class="list-group-item list-group-item-action"> -->
              プライバシーポリシー
            </router-link>
            <span class="divide px-2">|</span>
          </li>
          <li>
            <router-link to="/rule/contact" class="list-group-item-action">
            <!-- <router-link to="/input" class="list-group-item list-group-item-action"> -->
              お問い合わせ
            </router-link>
            <!-- <span class="divide px-2">|</span> -->
          </li>
        </ul>
        <ul class="list-unstyled">
          <li>
            <p class="copyright">
              Copyright©2022
              <span class="company">conforto</span>
              All rights reserved.
            </p>
          </li>
        </ul>
      </div>
    </div>
  </footer>
</template>

<script>
import { computed } from 'vue'
import { useStore } from 'vuex'

export default {
  name: "Footer",
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
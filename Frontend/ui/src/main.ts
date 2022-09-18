import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import axios from 'axios'

export const userURL = process.env.VUE_APP_USERURL
export const analysisURL = process.env.VUE_APP_ANALYSISURL

axios.defaults.withCredentials = true

export async function jsondec(data: Promise<ArrayBuffer>) {
  const check = new TextDecoder().decode(await data)
  const json = JSON.parse(check.toString());
  return json
}

createApp(App).use(store).use(router).mount('#app')
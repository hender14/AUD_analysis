import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import Home from '../pages/Home.vue'
import Register from '../pages/Register.vue'
import Login from '../pages/Login.vue'
import Logout from '../pages/Logout.vue'
import Forgot from '../pages/Forgot.vue'
import Reset from '../pages/Reset.vue'
import Teams from '../rule/teams.vue'
import Privacy from '../rule/privacy.vue'
import User from '../analysis/User.vue'
import Input from '../analysis/Input.vue'
import Result from '../analysis/Result.vue'
import Detail from '../analysis/Detail.vue'

const routes: Array<RouteRecordRaw> = [
  { path: '/', component: Home},
  { path: '/register', component: Register},
  { path: '/login', component: Login},
  { path: '/logout', component: Logout},
  { path: '/forgot', component: Forgot},
  { path: '/reset/:token', component: Reset},
  { path: '/rule/teams', component: Teams},
  { path: '/rule/privacy', component: Privacy},
  { path: '/user', component: User},
  { path: '/input', component: Input, name: "Input", props: true},
  { path: '/result', component: Result, name: "Result", props: true},
  { path: '/detail', component: Detail, name: "Detail", props: true},
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router

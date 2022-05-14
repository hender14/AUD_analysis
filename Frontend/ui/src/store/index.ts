import { Commit, createStore } from 'vuex'

export default createStore({
  state: {
    auth: false
  },
  mutations: {
    setAuth: (state: { auth: boolean }, auth: boolean) => { 
      state.auth = auth
    }
  },
  actions: {
    // ログイン済みかどうかの状態を管理する
    setAuth: ({commit}: { commit: Commit }, auth: boolean) => { 
      commit('setAuth', auth )
    }
  },
  // modules: {
  // }
})

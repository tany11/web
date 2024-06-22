import { createStore } from 'vuex'

export default createStore({
    state: {
        auth: {
            isLoggedIn: false
        }
    },
    mutations: {
        setLoggedIn(state, value) {
            state.auth.isLoggedIn = value
        }
    },
    actions: {
        login({ commit }) {
            // ログイン処理
            commit('setLoggedIn', true)
        },
        logout({ commit }) {
            // ログアウト処理
            commit('setLoggedIn', false)
        }
    }
})
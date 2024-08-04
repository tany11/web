import { createStore } from 'vuex'
import axios from 'axios'

export default createStore({
    state: {
        token: null,
        userId: null,
        isLoggedIn: false
    },
    mutations: {
        setToken(state, token) {
            state.token = token
        },
        setUserId(state, userId) {
            state.userId = userId
        },
        setLoggedIn(state, isLoggedIn) {
            state.isLoggedIn = isLoggedIn
        }
    },
    actions: {
        async login({ commit }, credentials) {
            try {
                const response = await axios.post('http://localhost:3000/api/v1/login', credentials)
                const { token, staff_id } = response.data
                commit('setToken', token)
                commit('setUserId', staff_id)
                commit('setLoggedIn', true)
                localStorage.setItem('token', token)
                localStorage.setItem('userId', staff_id)
                return true
            } catch (error) {
                console.error('ログインエラー:', error)
                return false
            }
        },
        async logout({ commit }) {
            try {
                await axios.post('http://localhost:3000/api/v1/logout')
                commit('setToken', null)
                commit('setUserId', null)
                commit('setLoggedIn', false)
                localStorage.removeItem('token')
                localStorage.removeItem('userId')
                return true
            } catch (error) {
                console.error('ログアウトエラー:', error)
                return false
            }
        }
    }
})
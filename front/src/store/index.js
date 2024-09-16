import { createStore } from 'vuex'
import axios from 'axios'

export default createStore({
    state: {
        apiBaseUrl: import.meta.env.VITE_API_BASE_URL || 'http://192.168.1.13:3000/api/v1',
        token: localStorage.getItem('token') || null,
        userId: localStorage.getItem('userId') || null,
        isLoggedIn: !!localStorage.getItem('token')
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
        async login({ commit, state }, credentials) {
            try {
                const response = await axios.post(`${state.apiBaseUrl}/login`, credentials)
                const { token, staff_id, group_id } = response.data
                commit('setToken', token)
                commit('setUserId', staff_id)
                commit('setLoggedIn', true)
                localStorage.setItem('token', token)
                localStorage.setItem('userId', staff_id)
                localStorage.setItem('groupId', group_id)

                // Set default Authorization header for all future requests
                axios.defaults.headers.common['Authorization'] = `Bearer ${token}`

                return true
            } catch (error) {
                console.error('ログインエラー:', error)
                return false
            }
        },
        async logout({ commit, state }) {
            try {
                await axios.post(`${state.apiBaseUrl}/logout`)
                commit('setToken', null)
                commit('setUserId', null)
                commit('setLoggedIn', false)
                localStorage.removeItem('token')
                localStorage.removeItem('userId')
                localStorage.removeItem('groupId')

                // Remove Authorization header
                delete axios.defaults.headers.common['Authorization']

                return true
            } catch (error) {
                console.error('ログアウトエラー:', error)
                return false
            }
        }
    }
})
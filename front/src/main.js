import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import vuetify from './plugin/vuetify'
import axios from 'axios'

// Add a request interceptor
axios.interceptors.request.use(
    (config) => {
        const token = localStorage.getItem('token')
        if (token) {
            config.headers['Authorization'] = `Bearer ${token}`
        }
        return config
    },
    (error) => {
        return Promise.reject(error)
    }
)

// Add a response interceptor
axios.interceptors.response.use(
    (response) => response,
    (error) => {
        if (error.response && error.response.status === 401) {
            // Token has expired or is invalid
            store.dispatch('logout')
            router.push('/login')
        }
        return Promise.reject(error)
    }
)

createApp(App)
    .use(router)
    .use(store)
    .use(vuetify)
    .mount('#app')
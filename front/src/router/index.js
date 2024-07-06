import { createRouter, createWebHistory } from 'vue-router'
import Login from '../views/Login.vue'
import Register from '../views/Register.vue'
import Dashboard from '../views/Dashboard.vue'
import Order from '../views/Order.vue'
import Staff from '../views/Staff.vue'
import Cast from '../views/Cast.vue'

const routes = [
    { path: '/', redirect: '/login' },
    { path: '/login', component: Login },
    { path: '/register', component: Register },
    { path: '/dashboard', component: Dashboard },
    { path: '/order', component: Order },
    { path: '/staff', component: Staff },
    { path: '/cast', component: Cast },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router
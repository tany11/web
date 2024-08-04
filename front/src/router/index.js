import { createRouter, createWebHistory } from 'vue-router'
import Login from '../views/Login.vue'
import Register from '../views/Register.vue'
import Dashboard from '../views/Dashboard.vue'
import Order from '../views/Order.vue'
import CastStaffRegistration from '../views/CastStaffRegistration.vue'
import StoreRegistration from '../views/StoreRegistration.vue'
import CustomerRegistration from '../views/CustomerRegistration.vue'


const routes = [
    { path: '/', redirect: '/login' },
    { path: '/login', name: 'Login', component: Login },
    { path: '/register', component: Register },
    { path: '/dashboard', component: Dashboard, meta: { requiresAuth: true } },
    { path: '/order', component: Order, meta: { requiresAuth: true } },
    { path: '/staff-registration', component: CastStaffRegistration, meta: { requiresAuth: true } },
    { path: '/store-registration', component: StoreRegistration, meta: { requiresAuth: true } },
    { path: '/customer-registration', component: CustomerRegistration, meta: { requiresAuth: true } },
]

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes
})

router.beforeEach((to, from, next) => {
    const isLoggedIn = localStorage.getItem('token') !== null && localStorage.getItem('userId') !== null
    if (to.matched.some(record => record.meta.requiresAuth) && !isLoggedIn) {
        next('/login')
    } else if ((to.path === '/login' || to.path === '/register') && isLoggedIn) {
        next('/dashboard')
    } else {
        next()
    }
})

export default router
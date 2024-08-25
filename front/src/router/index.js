import { createRouter, createWebHistory } from 'vue-router'
import Login from '../views/Login.vue'
import Register from '../views/Register.vue'
import Dashboard from '../views/Dashboard.vue'
import Order from '../views/Order.vue'
import CustomerRegistration from '../components/CustomerRegistration.vue'
import StaffRegistration from '../components/StaffRegistration.vue'
import CastRegistration from '../components/CastRegistration.vue'
import StaffEdit from '../components/StaffEdit.vue'
import CastEdit from '../components/CastEdit.vue'
import StoreRegistration from '../components/StoreRegistration.vue'
import StoreEdit from '../components/StoreEdit.vue'
import MediaRegistration from '../components/MediaForm.vue'
import Timeboard from '../views/TimeBoard.vue'

const routes = [
    { path: '/', redirect: '/login' },
    { path: '/login', name: 'Login', component: Login },
    { path: '/register', component: Register },
    { path: '/dashboard', component: Dashboard, meta: { requiresAuth: true } },
    { path: '/order', component: Order, meta: { requiresAuth: true } },
    { path: '/staff-registration', component: StaffRegistration, meta: { requiresAuth: true } },
    { path: '/staff-edit', component: StaffEdit, meta: { requiresAuth: true } },
    { path: '/cast-registration', component: CastRegistration, meta: { requiresAuth: true } },
    { path: '/cast-edit', component: CastEdit, meta: { requiresAuth: true } },
    { path: '/store-registration', component: StoreRegistration, meta: { requiresAuth: true } },
    { path: '/store-edit', component: StoreEdit, meta: { requiresAuth: true } },
    { path: '/media-registration', component: MediaRegistration, meta: { requiresAuth: true } },
    { path: '/customer-registration', component: CustomerRegistration, meta: { requiresAuth: true } },
    { path: '/timeboard', component: Timeboard, meta: { requiresAuth: true } },
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
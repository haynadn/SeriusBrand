import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import OrderSuccess from '../views/OrderSuccess.vue'
import UmkmTemplate from '../views/UmkmTemplate.vue'
import AdminDashboard from '../views/AdminDashboard.vue'
import LoginView from '../views/LoginView.vue'

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/',
            name: 'home',
            component: HomeView
        },
        {
            path: '/login',
            name: 'login',
            component: LoginView
        },
        {
            path: '/order-success/:id',
            name: 'order-success',
            component: OrderSuccess
        },
        {
            path: '/umkm/:slug',
            name: 'umkm',
            component: UmkmTemplate
        },
        {
            path: '/admin',
            name: 'admin',
            component: AdminDashboard,
            meta: { requiresAuth: true }
        }
    ]
})

// Navigation Guard
router.beforeEach((to, from, next) => {
    if (to.meta.requiresAuth) {
        const token = localStorage.getItem('admin_token')
        if (!token) {
            next('/login')
        } else {
            next()
        }
    } else {
        next()
    }
})

export default router

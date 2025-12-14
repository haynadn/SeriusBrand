import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import OrderSuccess from '../views/OrderSuccess.vue'
import UmkmTemplate from '../views/UmkmTemplate.vue'
import AdminDashboard from '../views/AdminDashboard.vue'

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/',
            name: 'home',
            component: HomeView
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
            component: AdminDashboard
        }
    ]
})

export default router

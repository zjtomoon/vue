import {createRouter,createWebHashHistory} from 'vue-router'
import Home from '@/views/Home'

// 路由的配置数组
const routes = [
    {
        path: '/',
        name: 'Home',
        component: Home
    },
]

// 创建路由对象
const router = createRouter({
    history:createWebHashHistory(process.env.BASE_URL),
    routes
})

export default router
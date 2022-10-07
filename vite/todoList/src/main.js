import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import store from './store'

createApp(App).use(store).mount('#app')

// 链式使用插件 
// createApp(App).use(store).use(router).use(ElementUI).mount('#app')
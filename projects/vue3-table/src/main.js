import { createApp } from 'vue'
import App from './App.vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import Antd from 'ant-design-vue'
// import Print from '@/utils/print'

const app = createApp(App)
app.use(ElementPlus)
app.use(Antd)
// app.use(Print)

app.mount('#app')

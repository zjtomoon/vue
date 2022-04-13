import Vue from 'vue'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import App from './App.vue'

// 引入自定义样式
import './styles/common.css'

Vue.use(ElementUI, { size: 'small', zIndex: 1000 })

Vue.config.productionTip = false

new Vue({
  render: h => h(App),
}).$mount('#app')

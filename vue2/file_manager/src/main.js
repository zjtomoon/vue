// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
// eslint-disable-next-line no-unused-vars
import axios from 'axios'
import BootstrapVue from 'bootstrap-vue'
// eslint-disable-next-line no-unused-vars
import store from './store'

Vue.use(BootstrapVue)

import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'


import {
  library
} from '@fortawesome/fontawesome-svg-core'
import {
  faFilePdf,
  faFileImage,
  faFileExcel,
  faFilePowerpoint,
  faFileWord,
  faFileVideo,
  faFileArchive,
  faFileAlt,
  faFile,
  faTrashAlt,
  faUpload,
  faTasks
} from '@fortawesome/free-solid-svg-icons'
import {
  FontAwesomeIcon
} from '@fortawesome/vue-fontawesome'

library.add(
  faFilePdf,
  faFileImage,
  faFileExcel,
  faFilePowerpoint,
  faFileWord,
  faFileVideo,
  faFileArchive,
  faFileAlt,
  faFile,
  faTrashAlt,
  faUpload,
  faTasks
)

Vue.component('font-awesome-icon', FontAwesomeIcon)

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: {App},
  template: '<App/>'
})

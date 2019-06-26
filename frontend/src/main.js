import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import './registerServiceWorker'
import Notifications from 'vue-notification'
import { apiService } from './services/apiService'
Vue.config.productionTip = false
Vue.prototype.$apiService = apiService
Vue.use(Notifications)

new Vue({
  router,
  store,
  render: function (h) { return h(App) }
}).$mount('#app')

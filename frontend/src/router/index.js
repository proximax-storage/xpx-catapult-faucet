import Vue from 'vue'
import VueRouter from 'vue-router'
import Faucet from '../views/Faucet.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Faucet',
    component: Faucet
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router

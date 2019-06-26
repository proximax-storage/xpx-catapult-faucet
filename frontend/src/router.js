import Vue from 'vue'
import Router from 'vue-router'
Vue.use(Router)
function load(component) {
  return () => import(/* webpackChunkName: "[request]" */ `@/views/${component}.vue`)
}
function loadChildren(component, child) {
  return () => import(/* webpackChunkName: "[request]" */ `@/components/${component}/${child}.vue`)
}
export default new Router({
  routes: [
    {
      path: '/faucet',
      component: load('FaucetGet')
    },
    {
      path: '*',
      redirect: '/faucet'
    },
    {
      path: '/',
      component: load('FaucetGet')
    },
  ]
})

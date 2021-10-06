import Vue from 'vue'
import App from './App.vue'

import BootstrapVue from "bootstrap-vue";
import VueRouter from "vue-router";

import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import Signup from "./pages/Signup";
import Login from "./pages/Login";
import Workspaces from "./pages/Workspaces";

Vue.use(BootstrapVue)
Vue.use(VueRouter)
Vue.config.productionTip = false

const routes = [
  { path: '/login', component: Login },
  { path: '/signup', component: Signup },
  { path: '/', component: Workspaces },
]

const router = new VueRouter({
  mode: 'history',
  routes: routes
})

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')

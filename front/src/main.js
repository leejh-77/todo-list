import Vue from 'vue'
import App from './App.vue'

import BootstrapVue from "bootstrap-vue";
import VueRouter from "vue-router";
import Vuetify from 'vuetify'
import 'vuetify/dist/vuetify.min.css'

import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import Signup from "./pages/Signup";
import Login from "./pages/Login";
import Workspaces from "./pages/Workspaces";
import TodoList from "./pages/TodoList";
import vuetify from './plugins/vuetify'

Vue.use(BootstrapVue)
Vue.use(VueRouter)
Vue.use(Vuetify)
Vue.config.productionTip = false

const routes = [
  { path: '/login', component: Login },
  { path: '/signup', component: Signup },
  { path: '/todo', component: TodoList },
  { path: '/', component: Workspaces },
]

const router = new VueRouter({
  mode: 'history',
  routes: routes
})

new Vue({
  router,
  vuetify,
  render: h => h(App)
}).$mount('#app')

export default new Vuetify({})


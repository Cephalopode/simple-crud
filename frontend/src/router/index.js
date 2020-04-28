import Vue from 'vue'
import VueRouter from 'vue-router'
import Users from '../components/Users.vue'

Vue.use(VueRouter)

const router = new VueRouter({
  routes: [
    {
      path: '/',
      name: 'Users',
      component: Users,
    },
  ],
})

export default router

import { createRouter, createWebHistory } from 'vue-router'
import UserList from '../views/UserList.vue'
import CacheManager from '../views/CacheManager.vue'

const routes = [
  { path: '/', redirect: '/users' },
  { path: '/users', component: UserList },
  { path: '/cache', component: CacheManager }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router 
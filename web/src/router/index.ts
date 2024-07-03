import { createRouter, createWebHistory } from 'vue-router'
import Sidebar from '../components/Grass.vue'
import Battle from '../components/Battle.vue'
import Tools from '../components/Tools.vue'
import SettingsPanel from '../components/SettingsPanel.vue'


const routes = [
  { path: '/', redirect: '/one-click-grass' },
  { path: '/one-click-grass', component: Sidebar },
  { path: '/auto-battle', component: Battle },
  { path: '/tools', component: Tools },
  { path: '/settings', component: SettingsPanel }
]


const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
})

export default router

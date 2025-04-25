import { createRouter, createWebHistory } from 'vue-router'
import Sidebar from '../components/GrassPanel.vue'
import BattlePanel from '../components/BattlePanel.vue'
import ToolsPanel from '../components/ToolsPanel.vue'
import SettingsPanel from '../components/SettingsPanel.vue'


const routes = [
  { path: '/', redirect: '/one-click-grass' },
  { path: '/one-click-grass', component: Sidebar },
  { path: '/auto-battle', component: BattlePanel },
  { path: '/tools', component: ToolsPanel },
  { path: '/settings', component: SettingsPanel }
]


const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})

export default router

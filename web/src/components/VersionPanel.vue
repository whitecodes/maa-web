<template>
    <div class="version">
      <div class="title">
        MAA - {{ version }}
      </div>
      <div class="tabs">
        <router-link to="/one-click-grass" @click="selectTab('one-click-grass')">一键长草</router-link>
        <router-link to="/auto-battle" @click="selectTab('auto-battle')">自动战斗</router-link>
        <router-link to="/tools" @click="selectTab('tools')">小工具</router-link>
        <router-link to="/settings" @click="selectTab('settings')">设置</router-link>
      </div>
    </div>
  </template>

<script>
import { useMainStore } from '../stores'
import api from '../services/api'
import { defineComponent,ref,onMounted } from 'vue';

export default defineComponent ({
  name: "VersionPanel",
  setup() {
    const store = useMainStore()
    const version = ref('')

    const selectTab = (tab) => {
      store.setSelectedTab(tab)
    }

    const fetchVersion = async () => {
      try {
        const data = await api.getVersion()
        version.value = data.version
      }catch (error) {
        console.error('Failed to fetch version:', error)
      }
    }

    onMounted(() => {
      fetchVersion()
    })


    return { selectTab, version }
  }
})
</script>

<style scoped>
.version {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px;
  background: #f0f0f0;
}
.title {
  font-weight: bold;
}
.tabs {
  display: flex;
  gap: 10px;
}
.tabs a {
  text-decoration: none;
  color: inherit;
  padding: 5px 10px;
  border-radius: 3px;
  transition: background-color 0.3s;
}
.tabs a:hover {
  background-color: #ddd;
}
</style>
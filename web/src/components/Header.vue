<template>
    <div class="header">
      <div class="title">
        MAA - v4.21.0 
      </div>
      <div class="tabs">
        <router-link to="/one-click-grass" @click.native="selectTab('one-click-grass')">一键长草</router-link>
        <router-link to="/auto-battle" @click.native="selectTab('auto-battle')">自动战斗</router-link>
        <router-link to="/tools" @click.native="selectTab('tools')">小工具</router-link>
        <router-link to="/settings" @click.native="selectTab('settings')">设置</router-link>
      </div>
    </div>
  </template>

<script>
import { useMainStore } from '../stores'

export default {
  name: "Header",
  setup() {
    const store = useMainStore()

    const selectTab = (tab) => {
      store.setSelectedTab(tab)
    }

    return { selectTab }
  }
}
</script>

<style scoped>
.header {
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
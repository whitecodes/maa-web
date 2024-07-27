<template>
  <div id="app">
    <VersionPanel />
    <div class="main-content">
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
import VersionPanel from './components/VersionPanel.vue'

export default {
  name: 'App',
  components: {
    VersionPanel
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
.main-content {
  display: flex;
}
</style>

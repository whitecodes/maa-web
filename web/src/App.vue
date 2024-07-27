<template>
  <div id="app">
    <Header />
    <div class="main-content">
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
import Version from './components/Version.vue'

export default {
  name: 'App',
  components: {
    Version
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

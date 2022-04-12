<template>
  <div id="app">
    <FIleManage />
  </div>
</template>

<script>
import FIleManage from './pages/fileManage.vue'

export default {
  name: 'app',
  components: {
    FIleManage
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
</style>

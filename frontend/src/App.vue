<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script lang="ts">
// TODO: each section takes a full vertical screen with 100vh

import 'bulma'
import 'jquery'

export default {
  name: 'App',
}
</script>

<style lang="scss">
@import '@/assets/scss/core.scss';
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  // -webkit-font-smoothing: antialiased;
  // -moz-osx-font-smoothing: grayscale;
  // text-align: center;
  // color: #2c3e50;
}
</style>

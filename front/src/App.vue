<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>

<script>
import bus from './event-bus'
import {ServerError} from "./const";

export default {
  name: 'App',
  components: {
  },
  created() {
    bus.$on(ServerError.Unauthenticated, () => {
      this.$router.push('/login')
    })
    bus.$on('logout', () => {
      this.$router.push('/login')
    })
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
  height: 100%;
}

html, body {
  height: 100%;
}

</style>

<template>
  <div class="page-header d-flex">
    <div class="header-home" @click="actionGoToHome">
      <img class="app-icon" src="../assets/logo.png">
      <p class="app-title">Todo-List</p>
    </div>
    <div class="user-info" @click="actionClickUserInfo">
      <img class="user-icon" src="../assets/user_icon.png"/>
      <p class="user-name">{{ user.username }}</p>
    </div>
  </div>
</template>

<script>

import {mapGetters} from "vuex";

export default {
  name: "MainHeader",
  methods: {
    actionGoToHome() {
      console.log('go to home')
    },
    actionClickUserInfo() {

    }
  },
  computed: {
    ...mapGetters([
        'user'
    ])
  },
  mounted() {
    this.$store.dispatch('loadMe')
  }
}
</script>

<style scoped>

.page-header {
  flex: none;
  padding: 9px 10px 8px;
  height: 60px;
  border-bottom: 1px solid #eee;
}

.header-home {
  position: absolute;
  cursor: pointer;
  align-content: center;
  line-height: 40px;
}

.app-icon {
  float: left;
  width: 40px;
}

.app-title {
  float: left;
  margin-left: 10px;
  font-weight: bold;
}

.user-info {
  cursor: pointer;
  position: absolute;
  right: 0;
  align-content: center;
  line-height: 40px;
}

.user-icon {
  width: 40px;
  float: left;
  margin-right: 10px;
}

.user-name {
  float: left;
  margin-right: 30px;
}

</style>
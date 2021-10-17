<template>
  <div class="page-header d-flex">
    <div class="header-home" @click="actionGoToHome">
      <img class="app-icon" src="../assets/logo.png">
      <p class="app-title">Todo-List</p>
    </div>
    <div class="user-info">
      <div class="user-info-profile">
        <img class="user-icon" :src="getUserImageUrl"/>
        <p class="user-name">{{ user.username }}</p>
      </div>
      <div class="user-info-dropdown">
        <p class="user-info-dropdown-menu" @click="actionShowProfile">Profile</p>
        <p class="user-info-dropdown-menu" @click="actionLogout">Log out</p>
      </div>
    </div>
    <UserInfoModal v-if="showUserInfoModal" @close="closeModal"/>
  </div>
</template>

<script>

import {mapGetters} from "vuex";
import UserInfoModal from "../modals/UserInfoModal";

export default {
  name: "MainHeader",
  components: {UserInfoModal},
  data() {
    return {
      showUserInfoModal: false
    }
  },
  methods: {
    actionGoToHome() {
      console.log('go to home')
    },
    actionShowProfile() {
      this.showUserInfoModal = true
    },
    actionLogout() {

    },
    closeModal() {
      this.showUserInfoModal = false
    }
  },
  computed: {
    ...mapGetters([
        'user'
    ]),
    getUserImageUrl() {
      let url = this.user.imageUrl
      if (url == null) {
        url = require('../assets/user_icon.png')
      }
      return url
    }
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
  display: block;
  align-content: center;
  line-height: 40px;
  z-index: 1;
}

.user-info-profile {
  overflow: hidden;
  display: block;
}

.user-info-dropdown {
  position: relative;
  display: none;
  background: white;
}

.user-info:hover .user-info-dropdown {
  display: block;
}

.user-info-dropdown-menu:hover {
  cursor: pointer;
  background: #cccccc;
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
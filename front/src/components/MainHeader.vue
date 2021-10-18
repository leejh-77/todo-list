<template>
  <div class="page-header d-flex">
    <div class="header-home" @click="actionGoToHome">
      <img class="app-icon" src="../assets/logo.png">
      <p class="app-title">Todo-List</p>
    </div>
    <div class="user-info">
      <div class="user-info-profile">
        <img class="user-icon" :src="getImageSource"/>
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
import userService from '../service/user'
import $ from "jquery";

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
      this.$router.push('/')
    },
    actionShowProfile() {
      this.showUserInfoModal = true
    },
    actionLogout() {
      userService.logout()
      .then(() => {
        this.$router.push('/login')
      })
    },
    closeModal() {
      this.showUserInfoModal = false
    },
    showUserInfo() {
      $('.user-info').css({
        'opacity': '0',
        'display': 'block'
      }).show().animate({opacity: 1})
    }
  },
  computed: {
    ...mapGetters([
        'user'
    ]),
    getImageSource() {
      let image = this.user.image
      if (image == null) {
        image = require('../assets/user_icon.png')
      } else {
        image = "data:" + image.type + ";base64," + image.data
      }
      return image
    },
    getUser() {
      return this.$store.state.user
    }
  },
  watch: {
    getUser() {
      this.showUserInfo()
    }
  },
  mounted() {
    if (!this.user.authenticated) {
      this.$store.dispatch('loadMe')
    } else {
      this.showUserInfo()
    }
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
  display: none;
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
  border-radius: 10%;
  box-shadow: 0px 0px 20px #cccccc;
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
  border-radius: 50%;
  float: left;
  margin-right: 10px;
}

.user-name {
  float: left;
  margin-right: 30px;
}

</style>
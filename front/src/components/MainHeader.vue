<template>
  <div class="page-header d-flex">
    <div class="header-home" @click="actionGoToHome">
      <img class="app-icon" src="../assets/logo.png">
      <p class="app-title">Todo-List</p>
    </div>
  </div>
</template>

<script>
import userService from "../service/user";
import workspaceService from "../service/workspace";
import folderService from "../service/folder";

export default {
  name: "MainHeader",
  methods: {
    actionGoToHome () {
      console.log('go to home')
    },
    getMyData() {
      userService.getMe()
          .then(res => {
            console.log(res)
            this.$store.commit('setUser', res.data)
            this.getWorkspaceData()
          })
          .catch(e => {
            console.log(e)
          })
    },
    getWorkspaceData() {
      workspaceService.getWorkspace(this.$route.query.workspaceId)
          .then(res => {
            console.log(res.data)
            this.$store.commit('setWorkspace', res.data)
            this.getFolderData()
          })
          .catch(e => {
            console.log(e)
          })
    },
    getFolderData() {
      folderService.getFolders(this.workspace.id, res => {
        if (res.status === 200) {
          this.folders = res.data == null ? [] : res.data
          this.selectedFolder = this.folders[0]
        } else {
          alert('something wrong')
        }
      })
    }
  },
  mounted() {
    this.getMyData()
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
}

</style>
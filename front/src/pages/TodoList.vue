<template>
  <div>
    <v-sheet height="700" class="overflow-hidden" style="position: relative;">
      <v-navigation-drawer absolute>
        <v-list-item>
          <v-list-item-avatar>
            <v-img src="../assets/user_icon.png"></v-img>
          </v-list-item-avatar>

          <v-list-item-content>
            <v-list-item-title>{{ user.username }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>

        <v-divider></v-divider>

        <v-list dense>
          <v-list-item v-for="folder in folders" :key="folder.id" link>
            <v-list-item-content>
              <v-list-item-title>{{ folder.name }}</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
          <v-list-item>
            <v-list-item-content id="add-folder-button" v-on:click="actionShowAddFolderModal">
              <v-list-item-title v-if="folders.length === 0">Add Folder</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </v-list>
      </v-navigation-drawer>
    </v-sheet>
    <AddFolderModal v-if="this.showAddFolderModal" @close="this.actionCloseAddFolderModal" @created="onFolderCreated"/>
  </div>
</template>

<script>

import folderService from '../service/folder'
import todoService from '../service/todo'
import {mapGetters} from 'vuex'
import userService from "../service/user";
import workspaceService from "../service/workspace"
import AddFolderModal from "../components/AddFolderModal";

export default {
  name: "TodoList",
  components: {AddFolderModal},
  data: function () {
    return {
      folders: [],
      selectedFolder: null,
      showAddFolderModal: false
    }
  },
  computed: {
    ...mapGetters([
        'user',
        'workspace'
    ])
  },
  methods: {
    actionShowAddFolderModal() {
      this.showAddFolderModal = true
    },
    actionCloseAddFolderModal() {
      this.showAddFolderModal = false
    },
    getTodos() {
      todoService.getTodos(this.selectedFolder, res => {
        this.todos = res.data
      })
    },
    onFolderCreated(folder) {
      this.folders.push(folder)
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
  },
}
</script>

<style scoped>
#folder-list-nav {
  background: #2c3e50;
  height: 100%;
  width: 100%;
}

#add-folder-button:hover {
  background: #cccccc;
  cursor: pointer;
}
</style>
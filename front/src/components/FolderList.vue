<template>
  <v-card height="700" width="256" class="main-page">
    <v-navigation-drawer permanent>
      <v-list-item class="workspace-info">
        <v-list-item-content>
          <v-list-item-title class="text-h6">{{ workspace.name }}</v-list-item-title>
        </v-list-item-content>
        <v-list-item-icon class="workspace-expand-icon">
          <v-img src="../assets/expand_icon.png"/>
        </v-list-item-icon>
      </v-list-item>

      <v-divider style="margin: 0"/>

      <v-list dense nav>
        <v-list-item v-for="folder in folders" :key="folder.id" link
                     :class="isSelectedFolder(folder) ? 'selected-folder-item' : ''" @click="onSelectFolder(folder)">
          <v-list-item-content>
            <v-list-item-title>{{ folder.name }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
        <v-list-item class="add-folder-button" link @click="actionShowAddFolderModal">
          <v-list-item-content>
            <v-list-item-title class="add-folder-title">Add Folder</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
    <AddFolderModal v-if="showAddFolderModal" @close="actionCloseAddFolderModal" @created="onFolderCreated"/>
  </v-card>
</template>

<script>
import {mapGetters} from "vuex";
import workspaceService from "../service/workspace";
import folderService from "../service/folder";
import AddFolderModal from "../modals/AddFolderModal";
import $ from 'jquery'

export default {
  name: "FolderList",
  computed: {
    ...mapGetters([
      'user',
      'workspace',
      'folder'
    ]),
  },
  components: {AddFolderModal},
  data() {
    return {
      folders: [],
      showAddFolderModal: false
    }
  },
  created() {
      this.getWorkspaceData()
  },
  methods: {
    isSelectedFolder(folder) {
      return this.folder.id === folder.id
    },
    onSelectFolder(folder) {
      this.$store.commit('setFolder', folder)
    },
    getWorkspaceData() {
      workspaceService.getWorkspace(this.$route.query.workspaceId)
          .then(res => {
            this.$store.commit('setWorkspace', res.data)
            this.getFolderData()
          })
    },
    getFolderData() {
      folderService.getFolders(this.$store.state.workspace.id)
          .then(res => {
            this.folders = res.data == null ? [] : res.data
            this.$store.commit('setFolder', this.folders[0])
            $('.main-page').css({
              'opacity': '0',
              'display': 'block'
            }).show().animate({opacity: 1})
          })
          .catch(() => {
            alert('something wrong')
          })
    },
    actionShowAddFolderModal() {
      this.showAddFolderModal = true
    },
    actionCloseAddFolderModal() {
      this.showAddFolderModal = false
    },
    onFolderCreated(folder) {
      this.folders.push(folder)
    }
  }
}
</script>

<style scoped>

.workspace-info {
  background: #eeeeee;
  height: 30px;
}

.workspace-info:hover {
  cursor: pointer;
}

.workspace-expand-icon {
  width: 15px;
  height: 20px;
}

.add-folder-button {
  background: cornflowerblue;
}

.add-folder-title {
  color: white;
}

.main-page {
  display: none;
}

</style>
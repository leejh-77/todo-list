<template>
  <v-card height="700" width="256" class="main-page">
    <div class="workspace-selector">
      <v-list-item class="workspace-selector-item"
                   v-for="workspace in workspaces"
                   :key="workspace.id" @click="onSelectWorkspace(workspace)">
        <v-list-item-content>
          <v-list-item-title class="text-h6">{{ workspace.name }}</v-list-item-title>
        </v-list-item-content>
        <v-list-item-icon class="workspace-selector-icon" v-if="isSelectedWorkspace(workspace)">
          <v-img src="../assets/check.png"/>
        </v-list-item-icon>
      </v-list-item>
    </div>
    <v-navigation-drawer permanent>
      <v-list-item class="workspace-info" @click="actionShowWorkspaceSelector">
        <v-list-item-content>
          <v-list-item-title class="text-h6">{{ workspace.name }}</v-list-item-title>
        </v-list-item-content>
        <v-list-item-icon class="workspace-expand-icon">
          <v-img src="../assets/expand_icon.png"/>
        </v-list-item-icon>
      </v-list-item>

      <v-divider style="margin: 0"/>

      <v-list dense nav>
        <v-list-item v-for="folder in workspace.folders" :key="folder.id" link
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
import AddFolderModal from "../modals/AddFolderModal";
import $ from 'jquery'

export default {
  name: "FolderList",
  computed: {
    ...mapGetters([
      'user',
      'workspace',
      'workspaces',
      'folder'
    ]),
  },
  components: {AddFolderModal},
  data() {
    return {
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
            console.log(res)
            this.$store.commit('setWorkspace', res.data)
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
      this.workspace.folders.push(folder)
    },
    actionShowWorkspaceSelector() {
      $('.workspace-selector').css('display', 'block')
    },
    isSelectedWorkspace(workspace) {
      return this.workspace.id === workspace.id
    },
    onSelectWorkspace(workspace) {
      this.$router.push('todo?workspaceId=' + workspace.id)
      this.getWorkspaceData()
      $('.workspace-selector').css('display', 'none')
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

.selected-folder-item {
  background: lightgray;
}

.workspace-selector {
  margin-top: 10px;
  display: none;
  width: 100%;
  position: absolute;
  background: white;
  z-index: 10;
  box-shadow: 0 0 20px #cccccc;
}

.workspace-selector-item:hover {
  cursor: pointer;
  background: #cccccc;
}

.workspace-selector-icon {
  width: 15px;
  height: 20px;
}

</style>
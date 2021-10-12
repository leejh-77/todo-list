<template>
  <v-card height="700" width="256">
    <v-navigation-drawer permanent>
      <v-list-item class="workspace-info">
        <v-list-item-content>
          <v-list-item-title class="text-h6">{{ workspace.name }}</v-list-item-title>
        </v-list-item-content>
        <v-list-item-icon>
          <v-img width="5px" src="../assets/expand_icon.png"/>
        </v-list-item-icon>
      </v-list-item>

      <v-divider style="margin: 0"/>

      <v-list dense nav>
        <v-list-item v-for="folder in folders" :key="folder.id" link>
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
    <AddFolderModal v-if="this.showAddFolderModal" @close="this.actionCloseAddFolderModal" @created="onFolderCreated"/>
  </v-card>
</template>

<script>
import {mapGetters} from "vuex";
import workspaceService from "../service/workspace";
import folderService from "../service/folder";
import AddFolderModal from "./AddFolderModal";

export default {
  name: "FolderList",
  computed: {
    ...mapGetters([
      'workspace'
    ])
  },
  components: {AddFolderModal},
  data() {
    return {
      folders: [],
      selectedFolder: null,
      showAddFolderModal: false
    }
  },
  mounted() {
    this.getWorkspaceData()
  },
  methods: {
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
      folderService.getFolders(this.$store.state.workspace.id)
          .then(res => {
            if (res.status === 200) {
              this.folders = res.data == null ? [] : res.data
              this.selectedFolder = this.folders[0]
            } else {
              alert('something wrong')
            }
          })
    },
    actionShowAddFolderModal() {
      this.showAddFolderModal = true
    },
    actionCloseAddFolderModal() {
      this.showAddFolderModal = false
    },
    onFolderCreated(folder) {
      console.log(folder)
      this.folders.push(folder)
    }
  }
}
</script>

<style scoped>

.workspace-info {
  background: #eeeeee;
}

.workspace-info:hover {
  cursor: pointer;
}


.add-folder-button {
  background: cornflowerblue;
}

.add-folder-title {
  color: white;
}


</style>
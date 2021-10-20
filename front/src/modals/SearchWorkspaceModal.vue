<template>
  <Modal @close="actionClose">
    <div class="main-page">
      <div class="search-input-wrapper">
        <b-input class="search-input"
                 placeholder="workspace name"
                 v-model="searchInput"
                 @keyup.enter="actionSearchWorkspace"/>
        <b-button class="search-button" @click="actionSearchWorkspace">Search</b-button>
      </div>
      <v-list-item class="search-item"
                   v-for="workspace in searchedWorkspaces"
                   :key="workspace.id"
                   @click="actionSelectWorkspace(workspace)">
        <v-list-item-content>
          <v-list-item-title>{{ workspace.name }}</v-list-item-title>
        </v-list-item-content>
      </v-list-item>
      <p class="search-no-results">No workspaces searched</p>
    </div>
    <template slot="footer">
      <b-button @click="actionAddWorkspace" v-if="this.selectedWorkspace != null">Add</b-button>
    </template>
  </Modal>
</template>

<script>
import Modal from "./Modal";
import workspaceService from "../service/workspace"
import $ from 'jquery'

export default {
  name: "SearchWorkspaceModal",
  components: {Modal},
  data() {
    return {
      searchedWorkspaces: [],
      searchInput: '',
      selectedWorkspace: null
    }
  },
  methods: {
    actionClose() {
      this.$emit('close')
    },
    actionSearchWorkspace() {
      this.selectedWorkspace = null
      workspaceService.searchWorkspace(this.searchInput)
          .then(res => {
            this.searchedWorkspaces = res.data == null ? [] : res.data
            if (this.searchedWorkspaces.length > 0) {
              $('.search-no-results').css('display', 'none')
            } else {
              $('.search-no-results').css('display', 'block')
            }
          })
    },
    actionSelectWorkspace(workspace) {
      this.selectedWorkspace = workspace
    },
    actionAddWorkspace() {
      workspaceService.addWorkspaceMember(this.selectedWorkspace.id)
      .then(() => {
          this.$emit('onAddWorkspace')
      })
    }
  }
}
</script>

<style scoped>

.main-page {
  height: 500px;
}

.search-input-wrapper {
  margin-bottom: 10px;
}

.search-input {
  width: 400px;
  float: left;
}

.search-button {
  margin-left: 10px;
  color: white;
}

.search-item:hover {
  background: #cccccc;
  cursor: pointer;
}

.search-no-results {
  display: none;
  margin-top: 20px;
}

</style>
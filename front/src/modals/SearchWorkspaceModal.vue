<template>
  <Modal @close="actionsClose">
    <div class="main-page">
      <b-input class="search-input" placeholder="workspace name" v-model="searchInput"></b-input>
      <b-button class="search-button" @click="actionSearchWorkspace">Search</b-button>
      <v-list-item v-for="workspace in searchedWorkspaces" :key="workspace.id">
        <v-list-item-content>
          <v-list-item-title>{{ workspace.name }}</v-list-item-title>
        </v-list-item-content>
      </v-list-item>
    </div>
  </Modal>
</template>

<script>
import Modal from "./Modal";
import workspaceService from "../service/workspace"

export default {
  name: "SearchWorkspaceModal",
  components: {Modal},
  data() {
    return {
      searchedWorkspaces: [],
      searchInput: ''
    }
  },
  methods: {
    actionsClose() {
      this.$emit('close')
    },
    actionSearchWorkspace() {
      workspaceService.searchWorkspace(this.searchInput)
      .then((res) => {
        this.searchedWorkspaces = res.data
      })
    }
  }
}
</script>

<style scoped>

.main-page {
  height: 500px;
}

.search-input {
  width: 400px;
  float: left;
}

.search-button {
  float: left;
  margin-left: 10px;
  color: white;
}

</style>
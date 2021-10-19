<template>
  <div class="main-page">
    <h1 class="app-title">Todo List</h1>
    <b-list-group id="workspace-list">
      <b-list-group-item id="workspace-list-item" v-for="workspace in workspaces" :key="workspace.id"
                         @click="actionMoveToWorkspace(workspace.id)">
        {{ workspace.name }}
      </b-list-group-item>
    </b-list-group>
    <p id="workspace-empty" v-show="workspaces.length === 0">There is no workspaces you are in</p>
    <img class="search-icon" src="../assets/search.png" @click="actionShowSearchModal"/>
    <b-button class="add-button" @click="this.actionShowModal">Add Workspace</b-button>
    <Modal @close="this.actionCloseModal" v-if="showModal">
      <p>Write name for workspace</p>
      <div>
        <b-input v-model="newWorkspaceName"/></div>
      <template slot="footer">
        <b-button @click="actionAddWorkspace">Add</b-button>
      </template>
    </Modal>
    <SearchWorkspaceModal class="search-modal" @close="actionCloseSearchModal" v-if="showSearchModal"/>
  </div>
</template>

<script>

import service from "../service/workspace";
import Modal from "../modals/Modal";
import $ from 'jquery'
import SearchWorkspaceModal from "../modals/SearchWorkspaceModal";

export default {
  name: "Workspaces",
  components: {SearchWorkspaceModal, Modal},
  data: function () {
    return {
      workspaces: [],
      showModal: false,
      newWorkspaceName: '',
      showSearchModal: false
    }
  },
  methods: {
    actionShowModal: function () {
      this.showModal = true
    },
    actionCloseModal: function() {
      this.showModal = false
      this.newWorkspaceName = ''
    },
    actionAddWorkspace: function () {
      service.addWorkspace(this.newWorkspaceName)
          .then(() => {
            this.getWorkspaces()
          })
      this.showModal = false
    },
    actionMoveToWorkspace: function (id) {
      this.$router.push('todo?workspaceId=' + id)
    },
    getWorkspaces: function () {
      service.getWorkspaces()
      .then(res => {
        this.workspaces = res.data == null ? [] : res.data
        $('.main-page').css("display", "flex")
      })
    },
    actionShowSearchModal() {
      this.showSearchModal = true
    },
    actionCloseSearchModal() {
      this.showSearchModal = false
    }
  },
  created() {
    this.getWorkspaces()
  }
}
</script>

<style scoped>

.main-page {
  display: none;
  flex-direction: column;
  align-items: center;
}

.app-title {
  margin-top: 20px;
}

.search-icon {
  width: 30px;
  position: relative;
}

.search-icon:hover {
  cursor: pointer;
}

ul {
  list-style: none;
}

#workspace-list {
  margin-bottom: 10px;
  margin-right: 10px;
  margin-left: 10px;
  width: 100%;
}

#workspace-list-item:hover {
  cursor: pointer;
  background: #eeeeee;
}

.add-button {
  margin-top: 20px;
  width: 200px;
}

</style>
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
    <b-button @click="this.actionShowModal">Add Workspace</b-button>
    <Modal @close="this.actionCloseModal" v-if="showModal">
      <p>Write name for workspace</p>
      <div>
        <b-input v-model="newWorkspaceName"/></div>
      <template slot="footer">
        <b-button @click="actionAddWorkspace">Add</b-button>
      </template>
    </Modal>
  </div>
</template>

<script>

import service from "../service/workspace";
import Modal from "../modals/Modal";
import $ from 'jquery'

export default {
  name: "Workspaces",
  components: {Modal},
  data: function () {
    return {
      workspaces: [],
      showModal: false,
      newWorkspaceName: ''
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
        $('.main-page').css("display", "block")
      })
    },
  },
  created() {
    this.getWorkspaces()
  }
}
</script>

<style scoped>

.main-page {
  display: none;
}

.app-title {
  margin-top: 20px;
}

ul {
  list-style: none;
}

#workspace-list {
  margin-bottom: 20px;
}

#workspace-list-item {
  margin: 10px;
}

#workspace-list-item:hover {
  cursor: pointer;
  background: #eeeeee;
}

</style>
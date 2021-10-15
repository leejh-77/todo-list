<template>
  <div>
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
import Modal from "../components/Modal";

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
      service.addWorkspace(this.newWorkspaceName, res => {
        if (res.status === 200) {
          this.getWorkspaces()
        }
        this.showModal = false
      })
    },
    actionMoveToWorkspace: function (id) {
      this.$router.push('todo?workspaceId=' + id)
    },
    getWorkspaces: function () {
      service.getWorkspaces(res => {
        this.workspaces = res.data == null ? [] : res.data
      })
    },
  },
  created() {
    this.getWorkspaces()
  }
}
</script>

<style scoped>

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
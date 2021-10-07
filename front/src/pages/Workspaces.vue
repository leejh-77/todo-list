<template>
  <div>
    <b-list-group id="workspace-list">
      <b-list-group-item id="workspace-list-item" v-for="workspace in workspaces" v-bind:key="workspace.id"
                         v-on:click="actionMoveToWorkspace(workspace.id)">
        {{ workspace.name }}
      </b-list-group-item>
    </b-list-group>
    <p id="workspace-empty" v-show="workspaces.length === 0">There is no workspaces you are in</p>
    <b-button v-on:click="this.actionShowModal">Add Workspace</b-button>
    <Modal v-on:close="this.actionCloseModal" v-if="showModal">
      <p>Write name for workspace</p>
      <div>
        <b-input v-model="newWorkspaceName"/></div>
      <template slot="footer">
        <b-button v-on:click="actionAddWorkspace">Add</b-button>
      </template>
    </Modal>
  </div>
</template>

<script>

import service from "../service/service";
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
      service.addWorkspace(this.newWorkspaceName, (ret) => {
        if (ret != null) {
          this.getWorkspaces()
        }
        this.showModal = false
      })
    },
    actionMoveToWorkspace: function (id) {
      this.$router.push('todo?workspaceId=' + id)
    },
    getWorkspaces: function () {
      service.getWorkspaces(data => {
        if (data == null) {
          this.$router.push('/login')
        } else if (data.data != null) {
            this.workspaces = data.data
          }
      })
    }
  },
  created() {
    this.getWorkspaces()
  }
}
</script>

<style scoped>

ul {
  list-style: none;
}

#workspace-list {
  margin-bottom: 20px;
}

#workspace-list-item {
}

</style>
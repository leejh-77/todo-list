<template>
  <div>
    <ul id="workspace-list">
      <li v-for="workspace in workspaces" v-bind:key="workspace.id">
        {{ workspace.name }}
      </li>
    </ul>
  </div>
</template>

<script>

import service from "../service/service";

export default {
  name: "Workspaces",
  data: function () {
    return {
      workspaces: []
    }
  },
  created() {
    service.getWorkspaces(data => {
      if (data == null) {
        this.$router.push('/login')
      } else {
        this.workspaces = data
      }
    })
  }
}
</script>

<style scoped>

</style>
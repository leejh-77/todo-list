<template>
  <div>
    <div id="folder-list-nav"/>
  </div>
</template>

<script>
export default {
  name: "TodoList",
  methods: {
    getFolders() {

    },
    addFolder() {

    },
    getTodos() {

    },

  },
  created() {
    let id = this.$route.query.workspaceId
    console.log(id)
  }
}
</script>

<style scoped>
#folder-list-nav {
  background: #2c3e50;
  height: 100%;
  width: 100%;
}
</style>
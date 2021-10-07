<template>
  <div>
    <v-sheet height="400" class="overflow-hidden" style="position: relative;">
      <v-navigation-drawer absolute>
        <v-list-item>
          <v-list-item-avatar>
            <v-img src="https://randomuser.me/api/portraits/men/78.jpg"></v-img>
          </v-list-item-avatar>

          <v-list-item-content>
            <v-list-item-title>John Leider</v-list-item-title>
          </v-list-item-content>
        </v-list-item>

        <v-divider></v-divider>

        <v-list dense>
          <v-list-item v-for="folder in folders" :key="folder.id" link>
            <v-list-item-content>
              <v-list-item-title>{{ folder.name }}</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </v-list>
      </v-navigation-drawer>
    </v-sheet>
  </div>
</template>

<script>

import service from '../service/folder'

export default {
  name: "TodoList",
  data: function () {
    return {
      workspaceId : 0,
      folders: []
    }
  },
  methods: {
    getFolders() {
      service.getFolders(this.workspaceId, res => {
        if (res.status === 200) {
          this.folders = res.data == null ? [] : res.data
        } else {
          alert('something wrong')
        }
      })
    },
    addFolder() {
      service.addFolder()
    },
    getTodos() {

    },

  },
  created() {
    this.workspaceId = this.$route.query.workspaceId
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
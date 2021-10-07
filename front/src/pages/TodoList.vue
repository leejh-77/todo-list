<template>
  <div>
      <v-sheet
          height="400"
          class="overflow-hidden"
          style="position: relative;"
      >
        <v-container class="fill-height">
          <v-row
              align="center"
              justify="center"
          >
            <v-btn
                color="pink"
                dark
                @click.stop="drawer = !drawer"
            >
              Toggle
            </v-btn>
          </v-row>
        </v-container>

        <v-navigation-drawer
            v-model="drawer"
            absolute
            temporary
        >
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
            <v-list-item
                v-for="item in items"
                :key="item.title"
                link
            >
              <v-list-item-icon>
                <v-icon>{{ item.icon }}</v-icon>
              </v-list-item-icon>

              <v-list-item-content>
                <v-list-item-title>{{ item.title }}</v-list-item-title>
              </v-list-item-content>
            </v-list-item>
          </v-list>
        </v-navigation-drawer>
      </v-sheet>
  </div>
</template>

<script>
export default {
  name: "TodoList",
  data: function () {
    return {
      drawer: null,
      items: [
        { title: 'Home', icon: 'mdi-view-dashboard' },
        { title: 'About', icon: 'mdi-forum' },
      ],
    }
  },
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
<template>
  <v-card height="700" width="256">
    <v-navigation-drawer permanent>
      <v-list-item class="user-name">
        <v-list-item-content>
          <v-list-item-title class="text-h6">{{ user.username }}</v-list-item-title>
        </v-list-item-content>
      </v-list-item>

      <v-divider style="margin: 0"/>

      <v-list dense nav>
        <v-list-item
            v-for="folder in folders"
            :key="folder.id"
            link
        >
          <v-list-item-content>
            <v-list-item-title>{{ folder.name }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
  </v-card>
</template>

<script>
import {mapGetters} from "vuex";

export default {
  name: "FolderList",
  computed: {
    ...mapGetters([
      'user'
    ])
  },
  data () {
    return {
      folders: []
    }
  }
}
</script>

<style scoped>

.user-name {
  background: lightyellow;
}

</style>
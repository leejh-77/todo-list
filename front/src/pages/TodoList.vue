<template>
  <div class="main-container">
    <MainHeader/>
    <FolderList class="folder-list"/>
    <TodoCards class="todo-board"/>
    <AddFolderModal v-if="this.showAddFolderModal" @close="this.actionCloseAddFolderModal" @created="onFolderCreated"/>
  </div>
</template>

<script>

import todoService from '../service/todo'
import {mapGetters} from 'vuex'
import AddFolderModal from "../components/AddFolderModal";
import MainHeader from "../components/MainHeader";
import FolderList from "../components/FolderList";
import TodoCards from "../components/TodoLists";

export default {
  name: "TodoList",
  components: {TodoCards, FolderList, MainHeader, AddFolderModal},
  data: function () {
    return {
      folders: [],
      selectedFolder: null,
      showAddFolderModal: false
    }
  },
  computed: {
    ...mapGetters([
        'user',
        'workspace'
    ])
  },
  methods: {
    actionShowAddFolderModal() {
      this.showAddFolderModal = true
    },
    actionCloseAddFolderModal() {
      this.showAddFolderModal = false
    },
    getTodos() {
      todoService.getTodos(this.selectedFolder, res => {
        this.todos = res.data
      })
    },
    onFolderCreated(folder) {
      this.folders.push(folder)
    },

  },
}
</script>

<style scoped>

.folder-list {
  float: left;
}

.main-container {
  height: 100%;
}

</style>
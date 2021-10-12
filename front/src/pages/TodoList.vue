<template>
  <div class="main-container">
    <MainHeader/>
    <FolderList class="folder-list"/>
    <TodoCards class="todo-board"/>
  </div>
</template>

<script>

import MainHeader from "../components/MainHeader";
import FolderList from "../components/FolderList";
import TodoCards from "../components/TodoBoard";

export default {
  name: "TodoList",
  components: {TodoCards, FolderList, MainHeader},
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
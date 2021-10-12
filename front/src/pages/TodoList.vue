<template>
  <div class="main-container">
    <MainHeader/>
    <FolderList class="folder-list"/>
    <TodoBoard class="todo-board"/>
  </div>
</template>

<script>

import MainHeader from "../components/MainHeader";
import FolderList from "../components/FolderList";
import TodoBoard from "../components/TodoBoard";

export default {
  name: "TodoList",
  components: {TodoBoard, FolderList, MainHeader},
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
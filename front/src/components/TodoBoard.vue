<template>
  <div class="page-body">
    <div class="list" v-for="list in cardLists" v-bind:key="list.status">
      <p class="list-title">{{list.name}}</p>
      <div class="card" v-for="todo in list.cards" v-bind:key="todo.id">
        <p>{{todo.name}}</p>
      </div>
      <div class="card-add-button" @click="actionAddNewTodo(list.status)">
        <p class="card-add-button-text">Add Todo</p>
      </div>
    </div>
  </div>
</template>

<script>
import {TodoStatus} from "../const";

export default {
  name: "TodoCards",
  data() {
    return {
      cardLists : [
        {
          status: TodoStatus.NotStarted,
          name : 'Not Started',
          cards : []
        },
        {
          status: TodoStatus.InProgress,
          name : 'In Progress',
          cards : []
        },
        {
          status: TodoStatus.Completed,
          name : 'Completed',
          cards : []
        },
      ]
    }
  },
  methods: {
    actionAddNewTodo(status) {
      console.log(status)
    }
  }
}
</script>

<style scoped>

.page-body {
  height: 700px;
  margin-top: 10px;
  margin-left: 270px;
  position: relative;
  flex-grow: 1;
  overflow-y: auto;
}

.list {
  width: 32%;
  float: left;
  background: #eee;
  border-radius: 5px;
  box-sizing: border-box;
  display: flex;
  flex-direction: column;
  max-height: 100%;
  white-space: normal;
  position: relative;
  margin-right: 10px;
}

.list-title {
  text-align: left;
  margin-left: 10px;
  margin-top: 10px;
  font-weight: bold;
}

.card {
  margin: 5px;
  border-radius: 3px;
}

.card-add-button-text {
  text-align: right;
  margin-right: 15px;
  margin-bottom: 10px;
}

</style>
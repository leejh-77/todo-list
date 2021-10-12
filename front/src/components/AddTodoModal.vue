<template>
  <v-app>
    <Modal class="todo-modal" @close="actionClose">
      <div class="todo-subject">
        <h5 class="input-title">Subject</h5>
        <b-input v-model="todo.subject"/>
      </div>
      <div class="todo-status">
        <h5 class="input-title">Status</h5>
        <v-container fluid class="todo-status-select">
          <v-row align="center">
            <v-col class="d-flex" cols="12" sm="6">
              <v-select :items="statusItems"
                        v-model="selectedStatus"
                        item-text="name"
                        item-value="value"></v-select>
            </v-col>
          </v-row>
        </v-container>
      </div>
      <div class="todo-body">
        <h5 class="input-title">Body</h5>
        <b-textarea class="todo-body-text" v-model="todo.body"/>
      </div>
      <template slot="footer">
        <b-button @click="actionAddTodo">Add</b-button>
      </template>
    </Modal>
  </v-app>
</template>

<script>

import Modal from "./Modal";
import {TodoStatus} from "../const";
import todoService from '../service/todo'

export default {
  name: "AddTodoModal",
  components: {Modal},
  props: ['status'],
  data() {
    return {
      todo : {
        subject: '',
        body: '',
      },
      selectedStatus: {},
      statusItems: [
        { name: 'Not Started', value: TodoStatus.NotStarted },
        { name: 'In Progress', value: TodoStatus.InProgress },
        { name: 'Completed', value: TodoStatus.Completed },
      ],
    }
  },
  methods: {
    actionAddTodo() {
      let todo = {
        subject: this.todo.subject,
        body: this.todo.body,
        folderId: this.$store.state.folder.id,
        userId: this.$store.state.user.id,
        status: this.selectedStatus.value
      }
      console.log(todo)
      todoService.createTodo(todo)
          .then(res => {
            this.$emit('onTodoCreated', res.data)
          })
    },
    actionClose() {
      this.$emit('close')
    },
  },
  mounted() {
    console.log(this.todo == null)
    let status = this.status
    this.statusItems.forEach(s => {
      if (s.value === status) {
        this.selectedStatus = s
      }
    })
  }
}
</script>

<style scoped>

.todo-subject {
  margin-top: 10px;
  width: 500px;
}

.input-title {
  text-align: left;
}

.todo-status {
  margin-top: 20px;
}

.todo-status-select {
  width: 100%;
  padding: 0;
  margin-left: 10px;
}

.todo-body {
  margin-top: 20px;
}

.todo-body-text {
  width: 100%;
  height: 300px;
}

b-button:hover {
  color: black;
}

</style>
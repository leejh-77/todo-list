<template>
  <v-app>
    <Modal class="todo-modal" @close="actionClose">
      <div class="todo-subject">
        <h6 class="input-title">Subject</h6>
        <b-input v-model="todo.subject"/>
      </div>
      <div class="todo-status">
        <h6 class="input-title">Status</h6>
        <v-container fluid class="todo-status-select">
          <v-row align="center">
            <v-col class="d-flex" cols="12" sm="6">
              <v-select :items="statusItems"
                        v-model="todo.status"
                        item-text="name"
                        item-value="value"></v-select>
            </v-col>
          </v-row>
        </v-container>
      </div>
      <div class="todo-body">
        <h6 class="input-title">Body</h6>
        <b-textarea class="todo-body-text" v-model="todo.body"/>
      </div>
      <template slot="footer">
        <b-button @click="actionDeleteTodo" v-if="todo.id != null">Delete</b-button>
        <b-button @click="actionAddTodo">{{ buttonText }}</b-button>
      </template>
    </Modal>
  </v-app>
</template>

<script>

import Modal from "./Modal";
import {TodoStatus} from "../const";

export default {
  name: "EditTodoModal",
  components: {Modal},
  props: ['todo'],
  data() {
    return {
      buttonText: '',
      statusItems: [
        { name: 'Not Started', value: TodoStatus.NotStarted },
        { name: 'In Progress', value: TodoStatus.InProgress },
        { name: 'Completed', value: TodoStatus.Completed },
      ],
    }
  },
  methods: {
    actionAddTodo() {
      this.$emit('onFinishEdit', this.todo)
    },
    actionDeleteTodo() {
      this.$emit('onDelete', this.todo)
    },
    actionClose() {
      this.$emit('close')
    },
  },
  mounted() {
    if (this.todo.id == null) {
      this.buttonText = 'Add'
    } else {
      this.buttonText = 'Update'
    }
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
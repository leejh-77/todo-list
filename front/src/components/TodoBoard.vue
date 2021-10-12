<template>
  <div class="page-body">
    <div class="list" v-for="list in cardLists" v-bind:key="list.status">
      <p class="list-title">{{list.name}}</p>
      <draggable :list="list.todos" group="todos" @change="log">
        <div class="card" v-for="todo in list.todos" :key="todo.id">
          <p class="card-subject">{{todo.subject}}</p>
          <p class="card-body">{{todo.body}}</p>
        </div>
      </draggable>
      <div class="card-add-button" @click="actionShowModal(list.status)">
        <p class="card-add-button-text">Add Todo</p>
      </div>
    </div>
    <AddTodoModal class="add-todo-modal"
                  :status="this.todoStatus"
                  v-if="this.showTodoModal"
                  @close="actionCloseModal"
                  @onTodoCreated="onTodoCreated"/>
  </div>
</template>

<script>
import draggable from 'vuedraggable'
import {TodoStatus} from "../const";
import {mapGetters} from "vuex";
import todoService from '../service/todo'
import AddTodoModal from "./AddTodoModal";

export default {
  name: "TodoCards",
  components: {AddTodoModal, draggable},
  data() {
    return {
      todoStatus: 0,
      showTodoModal: false,
      cardLists : [
        {
          status: TodoStatus.NotStarted,
          name : 'Not Started',
          todos : []
        },
        {
          status: TodoStatus.InProgress,
          name : 'In Progress',
          todos : []
        },
        {
          status: TodoStatus.Completed,
          name : 'Completed',
          todos : []
        },
      ]
    }
  },
  computed: {
    ...mapGetters([
        'folder'
    ]),
    getFolder() {
      return this.folder
    }
  },
  watch: {
    getFolder () {
      this.loadTodos()
    }
  },
  methods: {
    actionShowModal(status) {
      this.todoStatus = status
      this.showTodoModal = true
    },
    actionCloseModal() {
      this.showTodoModal = false
    },
    onTodoCreated(todo) {
      this.actionCloseModal()
      let list
      if (this.todoStatus === TodoStatus.NotStarted) {
        list = this.cardLists[0].todos
      } else if (this.todoStatus === TodoStatus.InProgress) {
        list = this.cardLists[1].todos
      } else {
        list = this.cardLists[2].todos
      }
      list.push(todo)
    },
    log(evt) {
      console.log(evt)
    },
    loadTodos() {
      todoService.getTodos(this.folder.id)
      .then(res => {
        let notStarted = this.cardLists[0]
        let inProgress = this.cardLists[1]
        let completed = this.cardLists[2]

        notStarted.todos = []
        inProgress.todos = []
        completed.todos = []

        let todos = res.data
        if (todos == null) {
          return
        }

        console.log(res)
        todos.forEach(todo => {
          let status = todo.status
          if (status === TodoStatus.NotStarted) {
            notStarted.todos.push(todo)
          } else if (status === TodoStatus.InProgress) {
            inProgress.todos.push(todo)
          } else {
            completed.todos.push(todo)
          }
        })
      })
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

.card-add-button {
  margin-top: 10px;
}

.card-add-button:hover {
  cursor: pointer;
}

.card-add-button-text {
  text-align: right;
  margin-right: 15px;
  margin-bottom: 10px;
}

.card {
  padding: 10px;
  border-radius: 10px;
}

.card-subject {
  font-weight: bold;
  text-align: left;
}

.card-body {
  text-align: left;
  padding: 0;
}

</style>
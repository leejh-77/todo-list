<template>
  <div class="page-body">
    <div class="list" v-for="list in cardLists" v-bind:key="list.status">
      <p class="list-title">{{list.name}}</p>
      <draggable :list="list.todos" group="todos" @change="log">
        <div class="card" v-for="todo in list.todos" :key="todo.id" @click="actionUpdateTodo(todo)">
          <p class="card-subject">{{todo.subject}}</p>
          <p class="card-body">{{todo.body}}</p>
        </div>
      </draggable>
      <div class="card-add-button" @click="actionCreateNewTodo(list.status)">
        <p class="card-add-button-text">Add Todo</p>
      </div>
    </div>
    <EditTodoModal class="add-todo-modal"
                  :todo="editingTodo"
                  v-if="showTodoModal"
                  @close="actionCloseModal"
                   @onFinishEdit="onFinishEdit"
                   @onDelete="onDelete"/>
  </div>
</template>

<script>
import draggable from 'vuedraggable'
import {TodoStatus} from "../const";
import {mapGetters} from "vuex";
import todoService from '../service/todo'
import EditTodoModal from "./EditTodoModal";

export default {
  name: "TodoCards",
  components: {EditTodoModal, draggable},
  data() {
    return {
      editingTodo: {},
      showTodoModal: false,
      todos: [],
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
    actionCreateNewTodo(status) {
      this.editingTodo = {
        subject: '',
        body: '',
        status: status
      }
      this.showTodoModal = true
    },
    actionUpdateTodo(todo) {
      this.editingTodo = {
        id: todo.id,
        subject: todo.subject,
        body: todo.body,
        status: todo.status
      }
      this.showTodoModal = true
    },
    actionCloseModal() {
      this.showTodoModal = false
      this.editingTodo = null
    },
    onFinishEdit() {
      let todo = this.editingTodo
      todo.folderId = this.$store.state.folder.id
      todo.userId = this.$store.state.user.id

      if (todo.id == null) {
        todoService.createTodo(todo)
            .then(res => {
              this.todos.push(res.data)
              this.alignTodos()
              this.actionCloseModal()
            })
      } else {
        todoService.updateTodo(todo)
        .then(res => {
          let updated = res.data
          this.todos.forEach(item => {
            if (item.id === updated.id) {
              item.subject = updated.subject
              item.body = updated.body
              item.position = updated.position
              item.status = updated.status
            }
          })
          this.alignTodos()
          this.actionCloseModal()
        })
      }
    },
    onDelete(todo) {
      todoService.deleteTodo(todo.id)
      .then(res => {
        console.log(res)
        var i
        for (i = 0; i < this.todos.length; i++) {
          if (todo.id === this.todos[i].id) {
            this.todos.splice(i, 1)
            break
          }
        }
      })
    },
    log(evt, evt2) {
      console.log('todo moved', evt)
      console.log('todo moved', evt2)
    },
    loadTodos() {
      todoService.getTodos(this.folder.id)
      .then(res => {
        this.todos = res.data
        this.alignTodos()
      })
    },
    alignTodos() {
      let notStarted = this.cardLists[0]
      let inProgress = this.cardLists[1]
      let completed = this.cardLists[2]

      notStarted.todos = []
      inProgress.todos = []
      completed.todos = []

      if (this.todos == null) {
        return
      }
      this.todos.forEach(todo => {
        let status = todo.status
        if (status === TodoStatus.NotStarted) {
          notStarted.todos.push(todo)
        } else if (status === TodoStatus.InProgress) {
          inProgress.todos.push(todo)
        } else {
          completed.todos.push(todo)
        }
      })

      let sortTodos = (t1, t2) => {
        return t1.position - t2.position
      }

      notStarted.todos.sort(sortTodos)
      inProgress.todos.sort(sortTodos)
      completed.todos.sort(sortTodos)
    },
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
  border-radius: 10px;
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
  padding: 10px;
  border-radius: 10px;
}

.card:hover {
  cursor: pointer;
}

.card-subject {
  font-weight: bold;
  text-align: left;
}

.card-body {
  text-align: left;
  padding: 0;
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


</style>
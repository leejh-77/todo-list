<template>
  <div class="page-body">
    <div class="list" v-for="list in cardLists" v-bind:key="list.status">
      <p class="list-title">{{list.name}}</p>
      <draggable :list="list.todos" group="todos" @change="log">
        <div class="card" v-for="todo in list.todos" :key="todo.id" @click="actionUpdateTodo(todo)">
          <p class="card-subject">{{todo.subject}}</p>
          <p class="card-body">{{todo.body}}</p>
          <div>
            <img class="card-user-profile" :src="getImage(todo.userId)"/>
          </div>
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
import EditTodoModal from "../modals/EditTodoModal";
import $ from 'jquery'

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
        'folder',
        'workspace'
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
        status: status,
        position: -1
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
        console.log(todo)
        todo.position = this.lastPosition(todo.status) + 1
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
      .then(() => {
        for (var i = 0; i < this.todos.length; i++) {
          if (todo.id === this.todos[i].id) {
            this.todos.splice(i, 1)
            break
          }
        }
        this.alignTodos()
        this.actionCloseModal()
      })
    },
    log(evt) {
      console.log('todo moved', evt)
      console.log(this.cardLists[0])
      console.log(this.cardLists[1])
      console.log(this.cardLists[2])
    },
    loadTodos() {
      if (this.folder == null || this.folder.id === 0) {
        $('.list').css('display', 'none')
        this.cardLists[0].todos = []
        this.cardLists[1].todos = []
        this.cardLists[2].todos = []
      } else {
        $('.list').css('display', 'flex')
        todoService.getTodos(this.folder.id)
            .then(res => {
              this.todos = res.data
              this.alignTodos()
              $('.page-body').css({
                'opacity': '0',
                'display': 'block'
              }).show().animate({opacity: 1})
            })
      }
    },
    getImage(uid) {
      let image
      for (var i = 0; i < this.workspace.members.length; i++) {
        let member = this.workspace.members[i]
        if (member.userId === uid) {
          image = member.image
          break
        }
      }
      if (image == null) {
        image = require('../assets/user_icon.png')
      } else {
        image = "data:" + image.type + ";base64," + image.data
      }
      return image
    },
    lastPosition(status) {
      let list
      if (status === TodoStatus.NotStarted) {
        list = this.cardLists[0].todos
      } else if (status === TodoStatus.InProgress) {
        list = this.cardLists[1].todos
      } else {
        list = this.cardLists[2].todos
      }
      if (list.length === 0) {
        return -1
      }
      let last = list[list.length - 1]
      return last.position
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
  display: none;
  min-height: 700px;
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

.card-user-profile {
  width: 40px;
  float: right;
  border-radius: 50%;
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
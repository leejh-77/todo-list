import axios from "axios";

export default {
    getTodos(folderId) {
        return new Promise((resolve, reject) => {
            axios.get("/api/todos?folderId=" + folderId)
                .then(res => resolve(res))
                .catch(e => reject(e))
        })
    },
    createTodo(todo) {
        return new Promise((resolve, reject) => {
            axios.post('/api/todos', todo)
                .then(res => resolve(res))
                .catch(e => reject(e))
        })
    },
    updateTodo(todo) {
        return new Promise((resolve, reject) => {
            axios.put('/api/todos/' + todo.id, todo)
                .then(res => resolve(res))
                .catch(e => reject(e))
        })
    },
    deleteTodo(id) {
        return new Promise((resolve, reject) => {
            axios.delete('/api/todos/' + id)
                .then(res => resolve(res))
                .catch(e => reject(e))
        })
    }
}
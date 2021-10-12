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
    }
}
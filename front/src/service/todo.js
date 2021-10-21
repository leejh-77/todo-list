import axios from "axios";

export default {
    getTodos(folderId) {
        return axios.get("/api/todos?folderId=" + folderId)
    },
    createTodo(todo) {
        return axios.post('/api/todos', todo)
    },
    updateTodo(todo) {
        return axios.put('/api/todos/' + todo.id, todo)
    },
    deleteTodo(id) {
        return axios.delete('/api/todos/' + id)
    },
    updatePositions(status, positions) {
        return axios.put('/api/todos/positions', {
            status: status,
            moveData: positions
        })
    }
}
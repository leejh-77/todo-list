import axios from "axios";

export default {
    async getTodos(folderId, callback) {
        try {
            let res = await axios.get("/api/todos?folderId=" + folderId)
            callback(res)
        } catch (e) {
            callback(e)
        }
    }
}
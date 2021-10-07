import axios from "axios";

export default {
    async getWorkspaces(callback) {
        try {
            let ret = await axios.get("/api/workspaces")
            callback(ret)
        } catch (e) {
            callback(e)
        }
    },
    async addWorkspace(name, callback) {
        try {
            let data = {
                'name' : name
            }
            let ret = await axios.post("/api/workspaces", data)
            callback(ret)
        } catch (e) {
            callback(e)
        }
    }
}
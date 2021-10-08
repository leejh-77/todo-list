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
    getWorkspace(wid) {
        return new Promise((resolve, reject) => {
            axios.get('/api/workspaces/' + wid)
                .then(res => resolve(res))
                .catch(e => reject(e))
        })
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
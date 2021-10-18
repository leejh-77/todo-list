import axios from "axios";

export default {
    getWorkspaces() {
        return axios.get("/api/workspaces")
    },
    getWorkspace(wid) {
        return axios.get('/api/workspaces/' + wid)
    },
    addWorkspace(name) {
        return axios.post("/api/workspaces", {
            'name' : name
        })
    }
}
import axios from "axios";

axios.defaults.baseURL = "http://localhost:9090"
axios.defaults.headers.post['Content-Type'] = 'application/json'
axios.defaults.withCredentials = true

export default {
    async login(email, password, callback) {
        let data = {
            emailAddress: email,
            password: password
        }
        try {
            await axios.post("/login", JSON.stringify(data))
            callback(true)
        } catch (e) {
            callback(false)
            console.error(e.response.data)
        }
    },
    async signup(email, password, username, callback) {
        let data = {
            emailAddress: email,
            password: password,
            username: username
        }
        try {
            await axios.post("/signup", JSON.stringify(data))
            callback(true)
        } catch (e) {
            callback(false)
            console.error(e.response.data)
        }
    },
    async getWorkspaces(callback) {
        try {
            let ret = await axios.get("/api/workspaces")
            callback(ret)
        } catch (e) {
            callback(null)
            console.error(e.response.data)
        }
    }
}

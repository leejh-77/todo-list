import axios from "axios";

axios.defaults.baseURL = "http://localhost:9090"
axios.defaults.headers.post['Content-Type'] = 'application/json'

export default {
    async login(email, password, callback) {
        let data = {
            emailAddress: email,
            password: password
        }
        try {
            let ret = await axios.post("/login", JSON.stringify(data))
            callback(ret)
        } catch (e) {
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
            let ret = await  axios.post("/signup", JSON.stringify(data))
            callback(ret)
        } catch (e) {
            console.error(e.response.data)
        }
    }
}

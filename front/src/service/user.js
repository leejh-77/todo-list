import axios from "axios";

export default {
    async login(email, password, callback) {
        let data = {
            emailAddress: email,
            password: password
        }
        try {
            let res = await axios.post("/login", data)
            console.log(res)
            callback(res)
        } catch (e) {
            callback(e)
        }
    },
    async signup(email, password, username, callback) {
        let data = {
            emailAddress: email,
            password: password,
            username: username
        }
        try {
            let res = await axios.post("/signup", data)
            console.log(res)
            callback(res)
        } catch (e) {
            callback(e)
        }
    },
}

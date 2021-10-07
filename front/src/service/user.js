import axios from "axios";

export default {
    login(email, password) {
        let data = {
            emailAddress: email,
            password: password
        }
        return new Promise((resolve, reject) => {
            axios.post('/login', data)
                .then(res => resolve(res))
                .catch(e => reject(e))
        })
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

import axios from "axios";

export default {
    getMe() {
        return new Promise((resolve, reject) => {
            axios.get('/api/users/me')
                .then(res => resolve(res))
                .catch(e => reject(e))
        })
    },
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
    signup(email, password, username) {
        let data = {
            emailAddress: email,
            password: password,
            username: username
        }
        return new Promise((resolve, reject) => {
            axios.post("/signup", data)
                .then(res => resolve(res))
                .catch(e => reject(e))
        })
    },
    update(imageData, username) {
        let data = {
            image: imageData,
            username: username
        }
        return new Promise((resolve, reject) => {
            axios.put('/api/user', data)
                .then(res => resolve(res))
                .catch(e => reject(e))
        })
    }
}

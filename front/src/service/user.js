import axios from "axios";

export default {
    getMe() {
        return axios.get('/api/users/me')
    },
    update(image, username) {
        let data
        if (image) {
            data = {
                image: image,
                username: username
            }
        } else {
            data = {
                username: username
            }
        }
        return axios.put('/api/users', data)
    },
    login(email, password) {
        return axios.post('/login', {
            emailAddress: email,
            password: password
        })
    },
    logout() {
        return axios.post('/logout')
    },
    signup(email, password, username) {
        return axios.post("/signup", {
            emailAddress: email,
            password: password,
            username: username
        })
    },

}

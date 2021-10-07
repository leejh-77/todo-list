import axios from "axios";

axios.defaults.baseURL = "http://localhost:9090"
axios.defaults.headers.post['Content-Type'] = 'application/json'
axios.defaults.withCredentials = true

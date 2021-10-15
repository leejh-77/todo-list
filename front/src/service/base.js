import axios from "axios";
import error_handler from "./error_handler";

axios.defaults.baseURL = "http://localhost:9090"
axios.defaults.headers.post['Content-Type'] = 'application/json'
axios.defaults.withCredentials = true

axios.interceptors.response.use(
    res => { return res },
    error => { return Promise.reject(error_handler.handle(error))})

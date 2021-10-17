import bus from '../event-bus'
import {ServerError} from "../const";


export default {
    handle(e) {
        let response = e.response
        console.log(response)
        if (response) {
            let status = response.status
            if (status === 401) {
                bus.$emit(ServerError.Unauthenticated)
            } else if (status === 400 && response.data.message === 'missing or malformed jwt') {
                bus.$emit(ServerError.Unauthenticated)
            }
        }
        return e
    }
}
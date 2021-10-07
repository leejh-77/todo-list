import axios from "axios";

export default {
    async getFolders(workspaceId, callback) {
        try {
            let res = await axios.get('/api/folders?workspaceId=' + workspaceId)
            callback(res)
        } catch (e) {
            callback(e)
        }
    }
}
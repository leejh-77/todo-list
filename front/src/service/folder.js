import axios from "axios";

export default {
    async getFolders(workspaceId, callback) {
        try {
            let res = await axios.get('/api/folders?workspaceId=' + workspaceId)
            callback(res)
        } catch (e) {
            callback(e)
        }
    },
    async addFolder(workspaceId, folderName, callback) {
        let data = {
            "workspaceId" : workspaceId,
            "name" : folderName
        }
        await axios.post('/api/folders', data)
            .then(res => callback(res))
            .catch(e => { callback(e) })
    }
}
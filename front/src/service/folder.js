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
    addFolder(workspaceId, folderName) {
        let data = {
            "workspaceId" : workspaceId,
            "name" : folderName
        }
        return new Promise((resolve, reject) => {
            axios.post('/api/folders', data)
                .then(res => resolve(res))
                .catch(e => reject(e))
        })
    }
}
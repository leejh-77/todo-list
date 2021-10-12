import axios from "axios";

export default {
    getFolders(workspaceId) {
        return new Promise((resolve, reject) => {
            axios.get('/api/folders?workspaceId=' + workspaceId)
                .then(res => resolve(res))
                .catch(e => reject(e))
        })
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
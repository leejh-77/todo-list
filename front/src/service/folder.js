import axios from "axios";

export default {
    getFolders(workspaceId) {
        return axios.get('/api/folders?workspaceId=' + workspaceId)
    },
    addFolder(workspaceId, folderName) {
        return axios.post('/api/folders', {
            "workspaceId" : workspaceId,
            "name" : folderName
        })
    }
}
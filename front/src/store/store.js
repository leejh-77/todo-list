import Vuex from 'vuex'
import Vue from "vue";
import userService from "../service/user";

Vue.use(Vuex)

export default new Vuex.Store({
    getters: {
        user: state => state.user,
        workspace: state => state.workspace,
        workspaces: state => state.workspaces,
        folder: state => state.folder
    },
    state: {
        user: {
            id: 0,
            username: '',
            emailAddress: '',
            authenticated: false,
            image: null,
        },
        workspaces: [],
        workspace: {
            id: 0,
            name: '',
            folders: [],
            members: []
        },
        folder: {
            id: 0,
            name: ''
        }
    },
    mutations: {
        setUser(state, user) {
            console.log(user)
            state.user = user
            state.user.authenticated = true
        },
        setWorkspaces(state, workspaces) {
            state.workspaces = workspaces
        },
        setWorkspace(state, workspace) {
            if (workspace.folders == null) {
                workspace.folders = []
            }
            state.workspace = workspace
            if (workspace.folders.length > 0) {
                state.folder = workspace.folders[0]
            } else {
                state.folder = null
            }
        },
        setFolder(state, folder) {
            state.folder = folder
        }
    },
    actions: {
        loadMe({commit}) {
            return userService.getMe().then(res => {
                commit('setUser', res.data)
                commit('setWorkspaces', res.data.workspaces)
            })
        }
    }
})
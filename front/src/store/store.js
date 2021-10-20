import Vuex from 'vuex'
import Vue from "vue";
import userService from "../service/user";
import bus from "../event-bus";

Vue.use(Vuex)

const EMPTY_USER = {
    id: 0,
    username: '',
    emailAddress: '',
    authenticated: false,
    image: null,
}

const EMPTY_WORKSPACE = {
    id: 0,
    name: '',
    folders: [],
    members: []
}

const EMPTY_FOLDER = {
    id: 0,
    name: ''
}

export default new Vuex.Store({
    getters: {
        user: state => state.user,
        workspace: state => state.workspace,
        workspaces: state => state.workspaces,
        folder: state => state.folder
    },
    state: {
        user: EMPTY_USER,
        workspaces: [],
        workspace: EMPTY_WORKSPACE,
        folder: EMPTY_FOLDER
    },
    mutations: {
        setUser(state, user) {
            state.user = user

            if (user !== EMPTY_USER) {
                state.user.authenticated = true
            }
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
                console.log(res)
                commit('setUser', res.data)
                commit('setWorkspaces', res.data.workspaces)
            })
        },
        logout({commit}) {
            userService.logout()
                .then(() => {
                    commit('setUser', EMPTY_USER)
                    commit('setWorkspaces', [])
                    commit('setFolder', EMPTY_FOLDER)
                    commit('setWorkspace', EMPTY_WORKSPACE)

                    bus.$emit('logout')
                })
        }
    }
})
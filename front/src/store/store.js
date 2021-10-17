import Vuex from 'vuex'
import Vue from "vue";
import userService from "../service/user";

Vue.use(Vuex)

export default new Vuex.Store({
    getters: {
        user: state => state.user,
        workspace: state => state.workspace,
        folder: state => state.folder
    },
    state: {
        user: {
            id: 0,
            username: '',
            emailAddress: '',
            imageUrl: null,
            authenticated: false
        },
        workspace: {
            id: 0,
            name: ''
        },
        folder: {
            id: 0,
            name: ''
        }
    },
    mutations: {
        setUser(state, user) {
            user.authenticated = true
            state.user = user
        },
        setWorkspace(state, workspace) {
            state.workspace = workspace
        },
        setFolder(state, folder) {
            state.folder = folder
        }
    },
    actions: {
        loadMe({commit}) {
            return userService.getMe().then(res => {
                commit('setUser', res.data)
            })
        }
    }
})
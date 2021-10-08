import Vuex from 'vuex'
import Vue from "vue";

Vue.use(Vuex)

export default new Vuex.Store({
    getters: {
        user: state => {
            console.log('getUser' + state.user.username)
            return state.user
        },
        workspace: state => state.workspace
    },
    state: {
        user: {
            id: 0,
            username: '',
            emailAddress: ''
        },
        workspace: {
            id: 0,
            name: ''
        }
    },
    mutations: {
        setUser(state, user) {
            console.log('setUser' + user.username)
            state.user = user
        },
        setWorkspace(state, workspace) {
            state.workspace = workspace
        }
    }
})
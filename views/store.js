const store = new Vuex.Store({

    state: {
        userPath: '',
        userAuthorized: false,
    },

    getters: {
        PATH: state => { return state.userPath; },
        USER_AUTHORIZED: state => { return state.userAuthorized; },
    },

    mutations: {
        SET_PATH: (state, value) => {
            state.userPath = value;
        },

        NEXT_PATH: (state, value) => {
            state.userPath += '/' + value;
        },

        PREV_PATH: (state) => {
            var pathSplit = state.userPath.split("/");
            state.userPath = pathSplit.slice(0, pathSplit.length - 1).join("/");
        },

        SET_USER_AUTHORIZED_STATUS: (state, value) => {
            state.userAuthorized = value;
        },
    },

    actions: {
    },

});
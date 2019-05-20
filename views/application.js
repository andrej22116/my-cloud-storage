var application = new Vue({
    store,
    el: "#application",
    components: {
        'filesystem': FILESYSTEM_COMPONENT,
        'login': LOGIN_COMPONENT,
        'logout': LOGOUT_COMPONENT,
        'registration': REGISTRATION_COMPONENT,
        'modal': MODAL_COMPONENT,
    },
    data: {
        user_nick: "NO USER",
        modal: false,
        login: false,
        parentPathProperty: "",
    },
    methods: {
        showLoginDialog: function() {
            this.modal = true;
            this.login = true;
        },

        showRegistrationDialog: function() {
            this.modal = true;
            this.login = false;
        },

        closeModal: function() {
            this.modal = false;
        },

        onLoginSuccess: function(userNickname) {
            this.user_nick = userNickname;
            this.$store.commit('SET_USER_AUTHORIZED_STATUS', true);
            this.$store.commit("SET_PATH", '');
        },

        onLogout: function() {
            this.$store.commit('SET_USER_AUTHORIZED_STATUS', false);
        },
    },

    created: function() {
        var token = window.localStorage['token'];
        if ( token == '' || token == undefined ) { 
            window.localStorage['token'] = '';
            return;
        }

        axios
            .post('http://' + SERVER_ADDRES + '/testtoken', {
                token: token
            })
            .then( response => { 
                this.$store.commit('SET_USER_AUTHORIZED_STATUS', true);
                this.user_nick = response.data.nickname;
            })
            .catch( err => {} );
    },

    computed: {
        currentPath() {
            return this.$store.getters.PATH;
        }
    }
});
var LOGOUT_COMPONENT = {
    template: `
        <div>
            <button @click="onLogout">Logout</button>
        </div>
    `,

    methods: {
        // Если есть инфа о токене - пробуем его уничтожить
        onLogout: function () {
            axios
                .post('http://' + SERVER_ADDRES + '/logout', {
                    token: window.localStorage["token"], 
                })
                .then(response => this.$emit('user-logout'))
                .catch(error => console.log(error));
        }
    }
}
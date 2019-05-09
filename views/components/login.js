var LOGIN_COMPONENT = {
    template: `
        <div>
            <input v-model="login" placeholder="Enter Your login">
            <input v-model="password" placeholder="Enter Your password">
            <button @click="onLogin">Login</button>
        </div>
    `,
    data: () => {
        return {
            login: "",
            password: "",
        }
    },
    methods: {
        onLogin: function () {
            axios
                .post('http://' + SERVER_ADDRES + '/authorization', {
                    login: this.login,
                    password: this.password,
                })
                .then(response => { window.localStorage["token"] = response.data.token; })
                .catch(error => console.log(error));
        }
    }
}
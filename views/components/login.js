var LOGIN_COMPONENT = {
    template: `
        <div>
            <input v-model="login" placeholder="Enter Your login">
            <input v-model="password" placeholder="Enter Your password" type="password">
            <p class="dialog-error" v-if="error">Invalid login or password!</p>
            <button @click="onLogin">Login</button>
        </div>
    `,
    data: () => {
        return {
            login: "",
            password: "",
            error: false
        }
    },
    methods: {
        onLogin: function () {
            axios
                .post('http://' + SERVER_ADDRES + '/authorization', {
                    login: this.login,
                    password: this.password,
                })
                .then(response => { 
                    window.localStorage["token"] = response.data.token;
                    this.$emit("login-success", response.data.nickname);
                    this.$emit("close-modal");
                })
                .catch(error => {
                    this.error = true;
                    console.log(error);
                });
        }
    }
}
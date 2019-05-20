var LOGIN_COMPONENT = {
    template: `
        <div>
            <input v-model="login" placeholder="Enter Your login">
            <input v-model="password" placeholder="Enter Your password" type="password">
            <p class="dialog-error" v-if="error">Invalid login or password!</p>
            <div class="dialog-ok"><button @click="onLogin">Login</button></div>
        </div>
    `,
    // Состояние
    data: () => {
        return {
            login: "",
            password: "",
            error: false
        }
    },
    // Методы
    methods: {
        onLogin: function () {
            // Делаем запрос
            axios
                .post('http://' + SERVER_ADDRES + '/authorization', {
                    login: this.login,
                    password: this.password,
                })
                // ОЕсли всё ок - запоминаем токен и сигналим
                .then(response => { 
                    window.localStorage["token"] = response.data.token;
                    // сигналим, что всё ок
                    this.$emit("login-success", response.data.nickname);
                    // и что надо закрыть диалог
                    this.$emit("close-modal");
                })
                .catch(error => {
                    // Иначе отобразим сообщение об ошибке.
                    this.error = true;
                    console.log(error);
                });
        }
    }
}
var REGISTRATION_COMPONENT = {
    template: `
        <div>
            <input v-model="login" placeholder="Enter Your login">
            <input v-model="password" placeholder="Enter Your password" type="password">
            <input v-model="repeatPassword" placeholder="Repeat password" type="password">
            <p class="dialog-error" v-if="error">Invalid data!</p>
            <div class="dialog-ok"><button @click="onRegistration">Registration!</button></div>
        </div>
    `,

    // состояние
    data: () => {
        return {
            login: "",
            password: "",
            repeatPassword: "",
            error: false,
        }
    },

    // Методы
    methods: {
        onRegistration() {
            // Проверяем поля
            if ( this.password != this.repeatPassword || this.login.length == 0 ) {
                this.error = true;
                return;
            }

            // Затем делаем запрос
            axios
                .post('http://' + SERVER_ADDRES + '/registration', {
                    login: this.login,
                    password: this.password,
                })
                .then( () => {
                    // Если всё ок - закрываем диалог
                    this.$emit("close-modal");
                })
                .catch(error => {
                    this.error = true;
                    console.log(error)
                });
        }
    }
}
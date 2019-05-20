var REGISTRATION_COMPONENT = {
    template: `
        <div>
            <input v-model="login" placeholder="Enter Your login">
            <input v-model="password" placeholder="Enter Your password" type="password">
            <input v-model="repeatPassword" placeholder="Repeat password" type="password">
            <p class="dialog-error" v-if="error">Invalid login or password!</p>
            <button @click="onRegistration">Registration!</button>
        </div>
    `,
    data: () => {
        return {
            login: "",
            password: "",
            repeatPassword: "",
        }
    },
    methods: {
        onRegistration: function () {
            if ( this.password != this.repeatPassword ) {
                alert("Different passwords are indicated!");
                return;
            }

            axios
                .post('http://' + SERVER_ADDRES + '/registration', {
                    login: this.login,
                    password: this.password,
                })
                .then( () => {
                    this.$emit("close-modal");
                })
                .catch(error => {
                    this.error = true;
                    console.log(error)
                });
        }
    }
}
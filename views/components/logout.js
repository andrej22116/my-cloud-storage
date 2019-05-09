var LOGOUT_COMPONENT = {
    template: `
        <div>
            <button @click="onLogout">Logout</button>
        </div>
    `,
    data: () => {
        return {
            login: "",
            password: "",
        }
    },
    methods: {
        onLogout: function () {
            axios
                .post('http://' + SERVER_ADDRES + '/logout', {
                    token: window.localStorage["token"], 
                })
                .then(response => { alert("Ok") })
                .catch(error => console.log(error));
        }
    }
}
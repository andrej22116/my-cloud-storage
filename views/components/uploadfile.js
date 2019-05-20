var UPLOAD_COMPONENT = {
    template: `
        <div class="upload-file">
            <input type="file" id="file" ref="file" @change="change" hidden/>
            <button @click="select">{{name}}</button>
            <button @click="upload">Upload</button>
        </div>
    `,
    data: () => {
        return {
            file: '',
            name: 'Select'
        }
    },
    methods: {
        change() {
            this.file = this.$refs.file.files[0];
            this.name = this.file.name
        },

        select() {
            document.getElementById("file").click()
        },

        upload() {
            var formData = new FormData();
            formData.append('file', this.file);

            axios
                .post('http://' + SERVER_ADDRES + '/upload', {
                    token: window.localStorage["token"],
                    path: this.$store.getters.PATH,
                    name: 'lol',
                })
                .then( response => {
                    axios
                        .post('http://' + SERVER_ADDRES + '/upload/' + response.data.uploadToken, formData, {
                            headers: {
                                'Content-Type': 'multipart/form-data'
                            }
                        })
                        .then(() => {
                            this.$emit('file-uploaded');
                        })
                        .catch(() => {
                            alert('FAILURE 2!!');
                            this.$emit('file-uploaded');
                        })
                })
                .catch(function(){
                    alert('FAILURE 1!!');
                });
        },
    }
}
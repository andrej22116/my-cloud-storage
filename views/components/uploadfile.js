var UPLOAD_COMPONENT = {
    props: ["uploadPath"],
    template: `
        <div>
            <input type="file" id="file" ref="file" @change="change"/>
            <button @click="upload">upload</button>
        </div>
    `,
    data: () => {
        return {
            file: '',
        }
    },
    methods: {
        change: function () {
            this.file = this.$refs.file.files[0];
        },

        upload: function () {
            var formData = new FormData();
            formData.append('file', this.file);

            axios
                .post('http://' + SERVER_ADDRES + '/upload', {
                    token: window.localStorage["token"],
                    path: this.uploadPath,
                    name: 'lol',
                })
                .then( response => {
                    axios
                        .post('http://' + SERVER_ADDRES + '/upload/' + response.data.uploadToken, formData, {
                            headers: {
                                'Content-Type': 'multipart/form-data'
                            }
                        })
                        .then(function(){
                            console.log('SUCCESS!!');
                        })
                        .catch(function(){
                            console.log('FAILURE!!');
                        })
                })
                .catch(function(){
                    console.log('FAILURE!!');
                });
        },
    }
}
var UPLOAD_COMPONENT = {
    prop: ["upload"],
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
            formData.append('token', JSON.stringify(window.localStorage["token"]));

            axios
                .post('http://' + SERVER_ADDRES + '/upload', formData, {
                    headers: {
                        'Content-Type': 'multipart/form-data'
                    }
                })
                .then(function(){
                    console.log('SUCCESS!!');
                })
                .catch(function(){
                    console.log('FAILURE!!');
                });
        },
    }
}
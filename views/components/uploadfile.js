var UPLOAD_COMPONENT = {
    template: `
        <div class="upload-file">
            <input type="file" id="file" ref="file" @change="change" hidden/>
            <button @click="select">{{name}}</button>
            <button @click="upload">Upload</button>
        </div>
    `,

    // состояние
    data: () => {
        return {
            file: '',
            name: 'Select'
        }
    },

    // Методы
    methods: {
        // Пользователь выбраз загружаемый файл
        change() {
            this.file = this.$refs.file.files[0];
            this.name = this.file.name
        },

        // Пользователь нажал на кнопку для выбора загружаемого файла
        select() {
            // Находим стандартный контроллер и инициируем его работу
            document.getElementById("file").click()
        },

        // Загружаем файл на сервер
        upload() {
            // Подготавливаем форму
            var formData = new FormData();
            formData.append('file', this.file);

            // Шлём запрос
            axios
                .post('http://' + SERVER_ADDRES + '/upload', {
                    token: window.localStorage["token"],
                    path: this.$store.getters.PATH,
                    name: 'lol',
                })
                // Если ответ полождительный - переходим ко второй части запроса
                .then( response => {
                    axios
                        .post('http://' + SERVER_ADDRES + '/upload/' + response.data.uploadToken, formData, {
                            headers: {
                                'Content-Type': 'multipart/form-data'
                            }
                        })
                        .then(() => {
                            // Ура, файл загружен
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
var FILESYSTEM_COMPONENT = {
    components: {
        'file': FILE_COMPONENT,
        'add': ADD_COMPONENT,
    },
    template: `
        <div class="filesystem">
            <button @click="update">update</button>
            <button @click="onGoParentFolder" v-if="!needHideBackButtonTest">Back</button>
            <div class="filelist">
                {{userAuthorizedTest}}
                <add @create-folder="onCreateNewFolder" @update="update"></add>
                <file
                    @on_next_folder="nextFolder"
                    @on_download_file="downloadFile"
                    @on_delete_file_obj="deleteFile"
                    @change-file-name="changeFileName"
                    v-for="(file, index) in filesList" 
                    :file="file"
                    :id="index"
                    :key="file.id"
                    >
                </file>
            </div>
        </div>
    `,

    // Состояние
    data: function() {
        return {
            newFolderName: "",
            filesList: [],
        }
    },

    // Методы
    methods: {
        // Переход в дочернюю папку
        nextFolder: function( folder ) {
            // Обновляем данные в хранилище
            this.$store.commit("NEXT_PATH", folder);
            this.update();
        },

        // Скачивание файла
        downloadFile: function( file ) {
            // Запрос
            axios
                .post('http://' + SERVER_ADDRES + '/load', {
                    token: window.localStorage["token"],    // токен
                    path: this.$store.getters.PATH,         // путь к файлу
                    name: file,                             // название файла
                })
                .then( response => {
                    // В случае успеха пробуем получить файл
                    window.open('http://' + SERVER_ADDRES + '/load/' + response.data.loadToken);
                })
                .catch( error => alert(error) ); // Иначе укажем на ошибку
        },

        // Удаляет файл
        deleteFile: function( fileItem ) {
            // Делаем запрос на сервер
            axios
                .post('http://' + SERVER_ADDRES + '/remove', {
                    token: window.localStorage["token"],
                    path: this.$store.getters.PATH,
                    name: fileItem.file.name,
                })
                // Если всё ок - просто обновляем список
                .then( () => {
                    this.update();
                })
                .catch( error => alert(error) );
        },

        // Обновление списка файлов в текущем каталоге
        update: function() {
            // Опустошаем список
            this.filesList = [];
            // Запрашиваем новый
            axios
                .post('http://' + SERVER_ADDRES + '/files', {
                    token: window.localStorage["token"],
                    path: this.$store.getters.PATH,
                })
                .then(response => {
                    // Список устанавливаем
                    this.filesList = response.data 
                })
                .catch( error => alert(error) );
        },

        // Создание новой папки
        onCreateNewFolder: function(name) {
            // Делаем запрос
            axios
                .post('http://' + SERVER_ADDRES + '/add/folder', {
                    token: window.localStorage["token"],
                    path: this.$store.getters.PATH,
                    name: name
                })
                // Обновляем
                .then( () => this.update() )
                .catch( error => alert(error) );
        },

        // Подняться на дерикторию выше 
        onGoParentFolder: function() {
            this.$store.commit("PREV_PATH");
            this.update();
        },

        // Запрос на изменение имени файла
        changeFileName(fileItem, newName) {
            axios
                .post('http://' + SERVER_ADDRES + '/modify', {
                    token: window.localStorage["token"],
                    path: this.$store.getters.PATH,
                    oldName: fileItem.file.name,
                    newName: newName,
                })
                // Если одобрили - обновляем значение
                .then( () => fileItem.file.name = newName )
                .catch( error => {alert(error); console.log(error);} );
        },
    },

    computed: {
        // С помощью этой шняги можно вовремя выполнить запрос на получение списка.
        userAuthorizedTest() {
            if ( this.$store.getters.USER_AUTHORIZED ) {
                this.update();
            }
            return;
        },

        // Тест на необходимость отображать кнопку возврата в родительский каталог
        needHideBackButtonTest() {
            return this.$store.getters.PATH == '';
        }
    },
};


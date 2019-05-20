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
                    v-for="(file, index) in filesList" 
                    :file="file"
                    :id="index"
                    :key="file.id"
                    >
                </file>
            </div>
        </div>
    `,
    data: function() {
        return {
            newFolderName: "",
            filesList: [],
        }
    },
    methods: {
        nextFolder: function( folder ) {
            this.$store.commit("NEXT_PATH", folder);
            this.update();
        },

        downloadFile: function( file ) {
            axios
                .post('http://' + SERVER_ADDRES + '/load', {
                    token: window.localStorage["token"],
                    path: this.$store.getters.PATH,
                    name: file,
                })
                .then( response => {
                    window.open('http://' + SERVER_ADDRES + '/load/' + response.data.loadToken);
                })
                .catch( error => alert(error) );
        },

        deleteFile: function( fileItem ) {
            axios
                .post('http://' + SERVER_ADDRES + '/remove', {
                    token: window.localStorage["token"],
                    path: this.$store.getters.PATH,
                    name: fileItem.file.name,
                })
                .then( () => {
                    this.filesList.splice(fileItem.id, 1);
                    this.update();
                })
                .catch( error => alert(error) );
        },

        update: function() {
            this.filesList = [];
            axios
                .post('http://' + SERVER_ADDRES + '/files', {
                    token: window.localStorage["token"],
                    path: this.$store.getters.PATH,
                })
                .then(response => {
                    this.filesList = response.data 
                })
                .catch( error => alert(error) );
        },

        onCreateNewFolder: function(name) {
            axios
                .post('http://' + SERVER_ADDRES + '/add/folder', {
                    token: window.localStorage["token"],
                    path: this.$store.getters.PATH,
                    name: name
                })
                .then( () => this.update() )
                .catch( error => alert(error) );
        },

        onGoParentFolder: function() {
            this.$store.commit("PREV_PATH");
            this.update();
        },
    },

    computed: {
        userAuthorizedTest() {
            if ( this.$store.getters.USER_AUTHORIZED ) {
                this.update();
            }
            return;
        },

        needHideBackButtonTest() {
            return this.$store.getters.PATH == '';
        }
    },
};


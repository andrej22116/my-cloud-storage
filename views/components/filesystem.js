var FILESYSTEM_COMPONENT = {
    props: ['filesystem'],
    components: {
        'file': FILE_COMPONENT,
        'upload': UPLOAD_COMPONENT,
    },
    template: `
        <div class="filesystem">
            <div class="menu">{{path}}</div>
            <button @click="update">Get files</button>
            <input v-model="newFolderName">
            <button @click="onCreateNewFolder">CreateFolder</button>
            <button @click="onGoParentFolder">back</button>
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
            <upload :upload-path="path"></upload>
        </div>
    `,
    data: function() {
        return {
            path: "",
            newFolderName: "",
            filesList: [],
        }
    },
    methods: {
        nextFolder: function( folder ) {
            console.log("Next folder: " + folder);
            this.path += "/" + folder;
            this.update();
        },

        downloadFile: function( file ) {
            axios
                .post('http://' + SERVER_ADDRES + '/load', {
                    token: window.localStorage["token"],
                    path: this.path,
                    name: file,
                })
                .then( response => {
                    window.open('http://' + SERVER_ADDRES + '/load/' + response.data.loadToken);
                })
                .catch(function(){
                    console.log('FAILURE!!');
                });
        },

        deleteFile: function( id ) {
            this.filesList.splice(id, 1);
        },

        update: function() {
            this.filesList = [];
            axios
                .post('http://' + SERVER_ADDRES + '/files', {
                    token: window.localStorage["token"],
                    path: this.path,
                })
                .then(response => {
                    this.filesList = response.data 
                })
                .catch(error => console.log(error));
        },

        onCreateNewFolder: function() {
            axios
                .post('http://' + SERVER_ADDRES + '/add/folder', {
                    token: window.localStorage["token"],
                    path: this.path,
                    name: this.newFolderName
                })
                .then(response => console.log('SUCCESS!!'))
                .catch(error => console.log(error));
        },

        onGoParentFolder: function() {
            var pathSplit = this.path.split("/");
            this.path = pathSplit.slice(0, pathSplit.length - 1).join("/");
            this.update();
        },
    }
};


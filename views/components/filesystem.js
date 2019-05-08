var FILESYSTEM_COMPONENT = {
    props: ['filesystem'],
    components: {
        'file': FILE_COMPONENT,
    },
    template: `
        <div class="filesystem">
            <button @click="update">Get files</button>
            <file 
                @on_next_folder="nextFolder"
                @on_download_file="downloadFile"
                @on_delete_file_obj="deleteFile"
                v-for="item in filesList"
                v-bind:key="item.id"
                v-bind:file="item">
            </file>
        </div>
    `,
    data: function() {
        return {
            filesList: [],
        }
    },
    methods: {
        nextFolder: function( folder ) {
            console.log("Next folder: " + folder);
        },
        downloadFile: function( file ) {
            console.log("Download file: " + file);
        },
        deleteFile: function( id ) {
            //this.filesList.remove();
            this.filesList = this.filesList.filter( file => {
                return +file.id != +id;
            });
        },
        update: function() {
            axios
                .get('http://' + SERVER_ADDRES + '/files')
                .then(response => { this.filesList = response.data.Data })
                .catch(error => console.log(error));
        },
    }
};


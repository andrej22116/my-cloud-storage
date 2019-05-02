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
            filesList: [
                { id: 0, name: 'file_1', date: '1.1.2019', isFolder: true },
                { id: 1, name: 'file_2', date: '2.2.2019', isFolder: false },
                { id: 2, name: 'file_3', date: '3.3.2019', isFolder: true },
                { id: 3, name: 'file_4', date: '4.4.2019', isFolder: false },
            ],
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


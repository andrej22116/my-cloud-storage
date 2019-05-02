var FILE_COMPONENT = {
    props: ['file'],
    template: `
        <div @dblclick="onDblClick" v-bind:class="[ file.isFolder ? 'folder' : 'file' ]">
            <div class="icon"></div>
            <div class="name">
                <div>{{ file.name }}</div>
                <!--<input v-model="message">-->
            </div>
            <div class="date">{{ file.date }}</div>
            <div class="options">
                <button @click="onBeginEdit" class="reneame">e</button>
                <button @click="onDelete" class="delete">x</button>
            </div>
        </div>
    `,
    /*data: function () {
        return {
            fileSize: 0,
        }
    },*/
    methods: {
        onDblClick: function() {
            this.$emit( this.file.isFolder ? 'on_next_folder' : 'on_download_file', this.file.name );
        },
        onDelete: function() {
            this.$emit( 'on_delete_file_obj', this.file.id );
            // send msg to server.
        },
        onBeginEdit: function() {

        },
        onEndEdit: function() {

        }
    }
};

// change data:
//объектВью.$children[порядковый номер].$data -> вернёт data;

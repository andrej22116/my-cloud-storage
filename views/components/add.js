var ADD_COMPONENT = {
    components: {
        'upload': UPLOAD_COMPONENT,
    },

    template: `
        <div class="add-container" 
        @mouseover="hover = true"
        @mouseleave="mouseLeave">
            <div class="add-background">+</div>
            <div class="add-content" v-if="true">
                <button v-if="state == 'base'" @click="state = 'file'">File</button>
                <button v-if="state == 'base'" @click="state = 'folder'">Folder</button>

                <div v-if="state == 'folder'" class="add-input">
                    <input v-model="newFolderName" placeholder="Folder name">
                    <button @click="onCreateNewFolder">CreateFolder</button>
                </div>
                
                <div v-if="state == 'file'" class="add-input">
                    <upload @file-uploaded="fileUploaded"></upload>
                </div>

                <button v-if="state != 'base'" @click="state = 'base'">Cancel</button>
            </div>
        </div>
    `,

    data: function () {
        return {
            hover: false,
            state: 'base',
            newFolderName: '',
        }
    },

    methods: {
        mouseLeave() {
            this.hover = false;
            //this.state = 'base';
        },

        onCreateNewFolder() {
            this.$emit('create-folder', this.newFolderName);
            this.newFolderName = '';
        },

        fileUploaded() {
            this.$emit('update');
            this.state = 'base';
        },
    }
};
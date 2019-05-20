var FILE_COMPONENT = {
    // Свойства объекта
    props: ['file', 'id'],
    template: `
        <div @dblclick="onDblClick" v-bind:class="[ {'folder' : file.isFolder}, {'file' : !file.isFolder && !isImage }, {'image' : isImage}, 'folder-item' ]">
            <div class="options" v-if="!edit">
                <button @click="onBeginEdit" class="reneame">e</button>
                <button @click="onDelete" class="delete">&times;</button>
            </div>
            <div>
                <div v-if="!edit" class="file-name">{{ file.name }}</div>
                <div v-if="edit">
                    <input v-model="fileName" placeholder="New name">
                    <button @click="onEndEdit">Ok</button>
                    <button @click="onBreakEdit">&times;</button>
                </div>
                <!--<input v-model="message">-->
            </div>
            <div class="date">{{ parsedDateTime }}</div>
        </div>
    `,

    // состояние
    data: function () {
        return {
            index: 0,
            edit: false,
            fileName: '',
        }
    },

    // Методы
    methods: {
        // Двойной клик
        onDblClick: function() {
            // Испускаем нужный сигнал в зависимости от ситуации
            this.$emit( this.file.isFolder ? 'on_next_folder' : 'on_download_file', this.file.name );
        },

        // Если нужно удалить файл
        onDelete: function() {
            // Сигналим с просьбой удалить файл
            this.$emit( 'on_delete_file_obj', this );
        },

        // Перед началом редактирования
        onBeginEdit: function() {
            // Фиксируем текущее имя файла (чтобы отобразить в текстовом поле)
            this.fileName = this.file.name;
            // Включаем режим редактирования
            this.edit = true;
        },

        // Завершение редактирования
        onEndEdit: function() {
            // Сигналим, что больше не редактируем
            this.edit = false;
            // И сообщаем родителю, что имя файла надо изменить
            this.$emit('change-file-name', this, this.fileName)
            
        },

        // Если решили отменить редактирование - просто отменяем, без сигналов
        onBreakEdit: function() {
            this.edit = false;
        },

        // Дрыц тыц
        getID(index) {
            this.index = index;
        }
    },

    // Вычисляемые значения
    computed: {
        // Преобразуем дату в адекватный вид перед выводом.
        parsedDateTime() {
            return new Date(this.file.date).toUTCString();
        }
    }
};

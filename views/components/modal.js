var MODAL_COMPONENT = {
    template: `
    <div class="modal-mask">
        <div class="modal-wrapper">
            <div class="modal-container">

                <div class="modal-header">
                    <slot name="title">
                        Modal title
                    </slot>
                    <button class="modal-close" @click="$emit('close-modal')">
                        &times;
                    </button>
                </div>

                <div class="modal-body">
                    <slot name="body">
                        Dialog content
                    </slot>
                </div>
            </div>
        </div>
    </div>
    `,
}
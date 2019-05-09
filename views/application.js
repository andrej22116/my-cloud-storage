var application = new Vue({
    el: "#content",
    components: {
        'filesystem': FILESYSTEM_COMPONENT,
        'login': LOGIN_COMPONENT,
        'logout': LOGOUT_COMPONENT,
        'registration': REGISTRATION_COMPONENT,
        'upload': UPLOAD_COMPONENT,
    },
    data: {
        user_nick: "NO USER",
        path: "/",
    }
});
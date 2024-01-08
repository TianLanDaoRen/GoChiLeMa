function onOpenLinkInExternalBrowser(url) {
    ipc.emit('requestOpenLinkInExternalBrowser', [url]);
}

function onMinWindow() {
    // emit request MinWindow
    ipc.emit("requestWindowState", [0]);
    console.log("requestWindowMin");
}

function onCloseWindow() {
    // emit request CloseWindow
    ipc.emit("requestWindowClose", []);
    console.log("requestWindowClose");
}
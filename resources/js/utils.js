function onOpenLinkInEdge(url) {
    ipc.emit('requestOpenLinkInEdge', [url]);
}
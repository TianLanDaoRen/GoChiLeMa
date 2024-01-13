import * as app from '../wailsjs/go/main/App.js'
import * as encrypto from '../wailsjs/go/encrypto/Encrypto.js'

// Make GoJs is available in the global scope
window.GoJs = {
    OpenLinkInExternalBrowser: app.OpenLinkInExternalBrowser,
    Encrypt: encrypto.Encrypt,
    Decrypt: encrypto.Decrypt,
    PKCS5Padding: encrypto.PKCS5Padding,
}
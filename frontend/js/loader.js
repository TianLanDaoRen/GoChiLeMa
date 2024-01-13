import * as wails from '../wailsjs/go/main/App.js'

// Make App is available in the global scope
window.App = {
    OpenLinkInExternalBrowser: wails.OpenLinkInExternalBrowser,
}
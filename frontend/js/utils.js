let API_ROOT = "http://localhost:23456/api";

// Callbacks
function onOpenLinkInExternalBrowser(url) {
    console.log("onOpenLinkInExternalBrowser");
    window.GoJs.OpenLinkInExternalBrowser(url);
}

function onMinWindow() {
    console.log("requestWindowMin");
    window.runtime.WindowMinimise();
}

function onToggleWindow() {
    window.runtime.WindowIsMaximised().then((isMaximised) => {
        if (isMaximised) {
            window.runtime.WindowUnmaximise();
        } else {
            window.runtime.WindowMaximise();
        }
    });
}

function onCloseWindow() {
    console.log("requestWindowClose");
    window.runtime.Quit();
}

// Route utils
function Fetch(url, method, body) {
    return fetch(url, {
        method: method,
        headers: {
            "Content-Type": "application/json",
            "Access-Control-Allow-Headers": "*",
            "Access-Control-Allow-Origin": "*",
            "Access-Control-Allow-Methods": "POST,GET,OPTIONS",
        },
        body: body,
    })
}

const utils = (function () {
    function getApiUrlAndBody(route) {
        return new Promise((resolve) => {
            let ts = new Date().getTime();
            let sn = Math.random() * 1e5;
            let api_url = `${API_ROOT}/${route}`;
            window.GoJs.Encrypt(`${sn}apiaccesskey||${ts}apiaccesskey||${route}`)
                .then((sign) => {
                    let body = {
                        "ts": ts,
                        "sn": sn,
                        "sign": sign,
                    };
                    resolve([api_url, body]);
                });
        });
    }

    // Public
    return {
        getApiUrlAndBody: getApiUrlAndBody,
    };
})();

// Set interval to update the toggle button
setInterval(() => {
    window.runtime.WindowIsMaximised().then((isMaximised) => {
        if (isMaximised) {
            document.getElementById("toggle-button").innerText = "2";
        } else {
            document.getElementById("toggle-button").innerText = "1";
        }
    });
}, 10);
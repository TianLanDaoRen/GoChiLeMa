let API_ROOT = "http://localhost:23456/api";

function Fetch(url, method = "GET") {
    return fetch(url, {
        method: method,
        headers: {
            "Access-Control-Allow-Headers": "*",
            "Access-Control-Allow-Origin": "*",
            "Access-Control-Allow-Methods": "POST,GET",
        },
    })
}

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

const utils = (function () {
    // Public
    return {
        getApiUrl(route) {
            return new Promise((resolve) => {
                let ts = new Date().getTime();
                let sn = Math.random() * 1e5;
                let api_url = `${API_ROOT}/${route}?ts=${ts}&sn=${sn}`;
                window.GoJs.Encrypt(`${sn}apiaccesskey||${ts}apiaccesskey||${route}`)
                    .then((sign) => {
                        api_url += `&sign=${sign}`;
                        resolve(api_url);
                    });
            });
        }
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
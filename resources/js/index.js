let API_ROOT = "http://localhost:23458/api";

(async function () {
    onRequestIpAndLocation().then(function () {
        console.log("onRequestIpAndLocationDone");
    });

    onRequestWeather().then(function () {
        console.log("onRequestWeatherDone");
    });
})();


// Callback functions
async function onContactUs() {
    console.log("requestContactUs");
    let api_url = utils.getApiUrl("contactus");
    let response = await fetch(api_url);
    let data = await response.json();
    console.log(data);
    // Retrieve email and vx
    let email = data.email;
    let vx = data.vx;
    // Update email and vx
    document.getElementById("emailInfo").innerText = email;
    document.getElementById("vxInfo").innerText = vx;
}

async function onRequestIpAndLocation() {
    console.log("requestIP");
    let api_url = utils.getApiUrl("ip");
    let response = await fetch(api_url);
    let data = await response.json();
    console.log(data);
    // Retrieve
    let ip = data.ip;
    let city = data.city;
    let lat = data.lat;
    let lon = data.lon;
    let timezone = data.timezone;
    let isp = data.isp;
    let as = data.as;
    // Update
    document.getElementById("ipInfo").innerText = ip;
    document.getElementById("locationInfo").innerText = city;
    document.getElementById("latInfo").innerText = lat;
    document.getElementById("lonInfo").innerText = lon;
    document.getElementById("timezoneInfo").innerText = timezone;
    document.getElementById("ispInfo").innerText = isp;
    document.getElementById("asInfo").innerText = as;
    document.getElementById("ipNotification").classList.remove("is-hidden");
}

async function onRequestWeather() {
    console.log("requestWeather");
    let api_url = utils.getApiUrl("weather");
    let response = await fetch(api_url);
    let data = await response.json();
    console.log(data);
    // Retrieve weather info
    let temp = data.temp;
    let description = data.description;
    // Update weather info
    document.getElementById("timeInfo").innerText = new Date().toString();
    document.getElementById("tempInfo").innerText = temp + "℃";
    document.getElementById("tempDesc").innerText = description;
    document.getElementById("weatherNotification").classList.remove("is-hidden");
}

// on receive Tick
ipc.on("receiveTick", function () {
    console.log("receiveTick");
    (function () {
        // Check if weather is out of date
        var timeInfo = document.getElementById("timeInfo").innerText;
        var time = new Date(timeInfo);
        var now = new Date();
        // If time is out of date for 1 hour
        if (now.getTime() - time.getTime() > 3600 * 1000) {
            // Request weather
            onRequestWeather();
        }
    })();
});

// on receive OsInfo
ipc.on("receiveOsInfo", function (os) {
    console.log("receiveOsInfo");
    document.getElementById("osInfo").innerText = os;
    document.getElementById("osNotification").classList.remove("is-hidden");
});
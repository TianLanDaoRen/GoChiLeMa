// Callback functions
function onContactUs() {
    // emit request ContactUs
    ipc.emit("requestContactUs", []);
    console.log("requestContactUs");
}

function onChiLeMa() {
    // emit request onChiLeMa
    // ipc.emit("requestContactUs", []);
}

function onRequestWeather() {
    // emit request Weather
    ipc.emit("requestWeather", []);
    console.log("requestWeather");
}

// IPC events
onRequestWeather();

// on receive AllowContactUs
ipc.on("receiveAllowContactUs", function (email, vx) {
    console.log("receiveAllowContactUs");
    document.getElementById("emailInfo").innerText = email;
    document.getElementById("vxInfo").innerText = vx;
});

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

// on receive Ip
ipc.on("receiveIp", function (ip) {
    console.log("receiveIp");
    document.getElementById("ipInfo").innerText = ip;
});

// on receive Location
ipc.on("receiveLocation", function (city, lat, lon, timezone, isp, as) {
    console.log("receiveLocation");
    document.getElementById("locationInfo").innerText = city;
    document.getElementById("latInfo").innerText = lat;
    document.getElementById("lonInfo").innerText = lon;
    document.getElementById("timezoneInfo").innerText = timezone;
    document.getElementById("ispInfo").innerText = isp;
    document.getElementById("asInfo").innerText = as;
    document.getElementById("ipNotification").classList.remove("is-hidden");
});

// on receive Weather
ipc.on("receiveWeather", function (temp, description) {
    console.log("receiveWeather");
    // Retrieve time info from local machine
    document.getElementById("timeInfo").innerText = new Date().toLocaleTimeString();
    // Update weather info
    document.getElementById("tempInfo").innerText = temp + "℃";
    document.getElementById("tempDesc").innerText = description;
    document.getElementById("weatherNotification").classList.remove("is-hidden");

});
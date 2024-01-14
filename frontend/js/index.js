// Callback functions
function onRequestOsInfo() {
    console.log("requestOsInfo");
    utils.getApiUrlAndBody("osinfo").then((api) => {
        let api_url = api[0], body = api[1];
        // Fetch data
        Fetch(api_url, "POST", body)
            .then(response => response.json())
            .then(data => {
                console.log(data);
                // Retrieve os info
                let info = data.info;
                // Update os info
                document.getElementById("osInfo").innerText = info;
                document.getElementById("osNotification").classList.remove("is-hidden");
            }
            )
            .catch(error => {
                console.error("Error:", error);
            });
    });
}

function onRequestContactUs() {
    console.log("requestContactUs");
    utils.getApiUrlAndBody("contactus").then((api) => {
        let api_url = api[0], body = api[1];
        // Fetch data
        Fetch(api_url, "POST", body)
            .then(response => response.json())
            .then(data => {
                // Retrieve email and vx
                let email = data.email;
                let vx = data.vx;
                // Update email and vx
                document.getElementById("emailInfo").innerText = email;
                document.getElementById("vxInfo").innerText = vx;
            })
            .catch(error => {
                console.error("Error:", error);
            });
    });
}

function onRequestIpAndLocation() {
    console.log("requestIP");
    utils.getApiUrlAndBody("ip").then((api) => {
        let api_url = api[0], body = api[1];
        // Fetch data
        Fetch(api_url, "POST", body)
            .then(response => response.json())
            .then(data => {
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
            })
            .catch(error => {
                console.error("Error:", error);
            });
    });
}

function onRequestWeather() {
    console.log("requestWeather");
    utils.getApiUrlAndBody("weather").then((api) => {
        let api_url = api[0], body = api[1];
        // Fetch data
        Fetch(api_url, "POST", body)
            .then(response => response.json())
            .then(data => {
                console.log(data);
                // Retrieve weather info
                let temp = data.temp;
                let description = data.description;
                // Update weather info
                document.getElementById("timeInfo").innerText = new Date().toString();
                document.getElementById("tempInfo").innerText = temp + "℃";
                document.getElementById("tempDesc").innerText = description;
                document.getElementById("weatherNotification").classList.remove("is-hidden");
            })
            .catch(error => {
                console.error("Error:", error);
            });
    });
}

// Set dom loaded event
document.addEventListener("DOMContentLoaded", () => {
    console.log("DOMContentLoaded");
    // Request os info
    onRequestOsInfo();
    // Request ip and location
    onRequestIpAndLocation();
    // Request weather
    onRequestWeather();
});
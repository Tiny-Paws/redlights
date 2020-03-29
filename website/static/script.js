redlight = new Vue({
    el: "#red",
    data: {
        color: "off"
    }
})

orangelight = new Vue({
    el: "#orange",
    data: {
        color: "off"
    }
})

greenlight = new Vue({
    el: "#green",
    data: {
        color: "off"
    }
})

const red = 1;
const orange = 2;
const green = 4;

function updateLightStatus() {
    let hs = window.location.hostname
    axios.get("http://"+hs+":5600/status").then(function(response) {
        redlight.color = computeColorPresence(red, response.data.state) ? "red" : "off";
        orangelight.color = computeColorPresence(orange, response.data.state) ? "orange" : "off";
        greenlight.color = computeColorPresence(green, response.data.state) ? "green" : "off";
    }).catch(function(error) {
        redlight.color = "off";
        orangelight.color = "off";
        greenlight.color = "off";
    });
}

function computeColorPresence(color, apidata) {
    return (color & apidata) == color
}

setInterval(updateLightStatus, 500)

redlight = new Vue({
    el: "#red",
    data: {
        color: "red"
    }
})

orangelight = new Vue({
    el: "#orange",
    data: {
        color: "orange"
    }
})

greenlight = new Vue({
    el: "#green",
    data: {
        color: "green"
    }
})

const red = 1;
const orange = 2;
const green = 4;

function updateLightStatus() {
    axios.get("http://localhost:5600/status").then(function(response) {
        redlight.color = computeColorPresence(red, response.data.state) ? "red" : "off";
        orangelight.color = computeColorPresence(orange, response.data.state) ? "orange" : "off";
        greenlight.color = computeColorPresence(green, response.data.state) ? "green" : "off";
    });
}

function computeColorPresence(color, apidata) {
    return (color & apidata) == color
}

setInterval(updateLightStatus, 500)

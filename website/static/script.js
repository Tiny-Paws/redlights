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
const hs = window.location.hostname

function updateLightStatus(data) {
    redlight.color = computeColorPresence(red, data) ? "red" : "off";
    orangelight.color = computeColorPresence(orange, data) ? "orange" : "off";
    greenlight.color = computeColorPresence(green, data) ? "green" : "off";
}

function computeColorPresence(color, apidata) {
    return (color & apidata) == color;
}

e = new EventSource('/status');
e.onmessage = function(event) {
    updateLightStatus(JSON.parse(event.data).state);
};
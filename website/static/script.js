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

function getStatus() {
    axios.get("http://localhost:5600/status").then(function(response) {
        console.log(response);
    });
}

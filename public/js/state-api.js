const $ = document.querySelector.bind(document); // We don't need much of jQuery - usually just this.

var pwnagotchi = pwnagotchi || {};

pwnagotchi.stateRetrieval = (function(){
    let _retrieval = function(callback){
        let xhttp = new XMLHttpRequest();
        xhttp.onreadystatechange = function() {
            if (this.readyState === 4 && this.status === 200) {
                let response = JSON.parse(this.responseText);

                let initialised = response.initialised !== "false";

                if (initialised === false){
                    const snore = $("#snore");
                    snore.innerText = snore.dataset.zeds.charAt(snore.dataset.index++);
                    if (snore.dataset.index>2) { snore.dataset.index = "0";}
                    return;
                }
                $("#initialiser").style.display = "none";
                $("#maindisplay").style.display = "block";

                callback(response)
            }
        };
        xhttp.open("GET", "/api/get", true);
        xhttp.send();
    };

    let initialise = function(callback){
        setInterval(_retrieval, 2000, callback)
    };

    return {
        initialise: initialise
    }
}());

var pwnagotchi = pwnagotchi || {};

pwnagotchi.populateDisplay = function(result){
    $("#channel").innerText = result.channel_text;
    $("#aps").innerText = result.aps_text;
    $("#uptime").innerText = result.uptime;

    $("#name").innerText = result.name + ">";
    $("#face").innerText = result.face;
    $("#status").innerText = result.status;

    $("#friend_face_text").innerText = result.friend_face_text !== null ? result.friend_face_text : "";
    $("#friend_name_text").innerText = result.friend_name_text !== null ? result.friend_name_text : "";

    $("#shakes").innerText = result.pwnd_run + "(" + result.pwnd_tot + ")";
    $("#mode").innerText = result.mode;

    $("#cpu").innerText = (result.cpu * 100).toFixed(2) + "%";
    $("#temperature").innerText = result.temperature.toFixed(2) +"c";
    $("#memory").innerText = (result.memory * 100).toFixed(2) + "%"
};
const $ = document.querySelector.bind(document); // We don't need much of jQuery - usually just this.

var pwnagotchi = pwnagotchi || {};

pwnagotchi.stateRetrieval = (function(){
    let _retrieval = function(callback){
        var hash = window.location.hash.slice(1);

        if (hash == "") {
            return;
        }

        let xhttp = new XMLHttpRequest();
        xhttp.onreadystatechange = function() {
            if (this.readyState === 4 && this.status === 200) {
                let response = JSON.parse(this.responseText);

                let initialised = response.initialised === true;

                if (initialised === false){
                    $("#initialiser").style.display = "block";
                    $("#maindisplay").style.display = "none";

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
        xhttp.open("GET", "/api/get?hash="+hash, true);
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
    window.document.title = result.name + ">";
    
    $("#face").innerText = result.face;
    $("#status").innerText = !result.status ? "" : result.status;

    $("#friend_face_text").innerText = !result.friend_face_text ? "" : result.friend_face_text;
    $("#friend_name_text").innerText = !result.friend_name_text ? "" : result.friend_name_text;

    $("#bluetooh").innerText = !result.bluetooth ? "" : "ᛒ"+result.bluetooth;

    if ((result.total_messages && result.total_messages>0) || (result.unread_messages && result.unread_messages>0)) {
        var mail_text = " ✉";
        if (result.unread_messages && result.unread_messages>0) {
            mail_text += result.unread_messages + "/";
        }

        if (result.total_messages && result.total_messages>0) {
            mail_text += result.total_messages;
        }

        $("#mail").innerText = mail_text;
    }

    if (result.pwnd_deauth) {
        result.pwnd_run += "/" + result.pwnd_deauth
    }

    var pwnd_last = ""
    if (result.pwnd_last) {
         pwnd_last = " [" + result.pwnd_last +"]"
    }

    $("#shakes").innerText = result.pwnd_run + "(" + result.pwnd_tot + ")"+pwnd_last;
    $("#mode").innerText = result.mode;

    $("#cpu").innerText = !result.cpu ? "" : (result.cpu * 100).toFixed(0) + "%";
    $("#temperature").innerText = !result.temperature ? "" : result.temperature.toFixed(0) +"°C";
    $("#memory").innerText = !result.memory ? "" : (result.memory * 100).toFixed(0) + "%"
};
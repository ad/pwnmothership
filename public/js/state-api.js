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

var prevStatus = "";

pwnagotchi.populateDisplay = function(result){
    $("#channel").innerText = result.channel_text;
    $("#aps").innerText = result.aps_text.replace(/\s+/g, '');
    $("#uptime").innerText = result.uptime;

    $("#name").innerText = result.name + ">";
    window.document.title = result.name + ">";

    if (result.level) {
        $("#level").innerText = result.level;
    }
    if (result.exp) {
        $("#exp").innerText = result.exp;
    }
    
    $("#face").innerText = result.face;

    $("#friend_face_text").innerText = !result.friend_face_text ? "" : result.friend_face_text;
    $("#friend_name_text").innerText = !result.friend_name_text ? "" : result.friend_name_text;

    if ((result.total_messages && result.total_messages>0) || (result.unread_messages && result.unread_messages>0)) {
        var mail_text = " ✉";
        if (result.unread_messages && result.unread_messages>0) {
            mail_text += result.unread_messages + "/";
        }

        if (result.total_messages && result.total_messages>0) {
            mail_text += result.total_messages;
        }

        $("#mail").innerText = mail_text;
    } else {
        $("#mail").innerText = "";
    }

    if (result.pwnd_deauth) {
        result.pwnd_run += "/" + result.pwnd_deauth
    }

    var pwnd_last = ""
    if (result.pwnd_last) {
         pwnd_last = "[" + result.pwnd_last +"]"
    }

    $("#shakes").innerText = result.pwnd_run + "(" + result.pwnd_tot + ")"+pwnd_last;
    $("#mode").innerText = result.mode;

    $("#cpu").parentNode.style.display = (result.cpu) ? 'block' : 'none';
    $("#cpu").innerText = !result.cpu ? "?%" : (result.cpu * 100).toFixed(0) + "%";

    $("#temperature").parentNode.style.display = (result.temperature) ? 'block' : 'none';
    $("#temperature").innerText = !result.temperature ? "?" : result.temperature.toFixed(0) +"°C";

    $("#memory").parentNode.style.display = (result.memory) ? 'block' : 'none';
    $("#memory").innerText = !result.memory ? "?%" : (result.memory * 100).toFixed(0) + "%";

    $("#bluetooth").parentNode.style.display = (result.bluetooth) ? 'block' : 'none';
    $("#bluetooth").innerText = !result.bluetooth ? "" : result.bluetooth;

    $("#ups").parentNode.style.display = (result.ups) ? 'block' : 'none';
    $("#ups").innerText = !result.ups ? "" : result.ups;

    if (result.status) {
        if (result.status != prevStatus && result.status != "...") {
            var ul = $('#logs').childNodes;

            if (ul.length > 10) {
                $('#logs').removeChild(ul[ul.length-1]);
            }

            prevStatus = result.status;
            var node = document.createElement('li');
            var nodeText = document.createTextNode(result.status);
            node.appendChild(nodeText);
            $("#logs").prepend(node);
        }
    }
};
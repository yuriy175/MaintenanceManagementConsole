const socket = new WebSocket("ws://localhost:8080/echo");

    socket.onopen = function () {
        console.log("Status: Connected\n");
        socket.send("789 from ui");
    };

    socket.onmessage = function (e) {
        console.log("Server: " + e.data + "\n");
    };
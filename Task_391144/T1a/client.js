const token = "valid-token"; // Replace with "invalid-token" to test unauthorized access
const socket = new WebSocket(`ws://localhost:8080/ws?token=${token}`);

socket.onopen = function(event) {
    console.log("Connection established");
    socket.send("Hello Server!");
};

socket.onmessage = function(event) {
    console.log("Received from server: ", event.data);
};

socket.onerror = function(error) {
    console.error("WebSocket error: ", error);
};

socket.onclose = function(event) {
    console.log("Connection closed");
};
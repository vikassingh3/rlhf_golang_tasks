// Replace the token below with the one generated by the server
const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzQxNzQ2NDYsInVzZXJuYW1lIjoidXNlcjEifQ.FUJq3J8yk5T4-dZNrg5M9HJwPY7Y-6CS6zSFEpLZEtM";
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
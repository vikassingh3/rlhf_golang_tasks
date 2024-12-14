const socket = new WebSocket('ws://localhost:8080/ws');
const token = 'your_valid_jwt_token_here';

socket.addEventListener('open', () => {
    socket.send(JSON.stringify({ action: 'hello' }));
});

socket.addEventListener('message', (event) => {
    const response = JSON.parse(event.data);
    console.log('Server response:', response);
});

socket.addEventListener('error', (error) => {
    console.error('WebSocket error:', error);
});

socket.addEventListener('close', () => {
    console.log('WebSocket connection closed');
});
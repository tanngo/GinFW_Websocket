const WebSocket = require('ws');

// Replace 'ws://localhost:8080' with the WebSocket server endpoint you want to connect to
const socket = new WebSocket('ws://localhost:7007/ws');

// Connection opened
socket.addEventListener('open', (event) => {
  console.log('WebSocket connection opened');
  socket.send('Hello, Server!');
});

// Listen for messages from the server
socket.addEventListener('message', (event) => {
  console.log('Received message from server:', event.data);
});

// Listen for errors
socket.addEventListener('error', (event) => {
  console.error('WebSocket encountered an error:', event.error);
});

// Connection closed
socket.addEventListener('close', (event) => {
  console.log('WebSocket connection closed:', event.code, event.reason);
});

// Gracefully close the connection after 5 seconds
setTimeout(() => {
  socket.close();
}, 5000);

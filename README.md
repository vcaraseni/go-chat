# Go-Chat

This is a pet project to practice Golang. The current version works without storing messages in the database.

## Installation
Make shure that port 8080 is not used
```bash
git clone git@github.com:vcaraseni/go-chat.git
npm install -g wscat
cd go-chat/
go run cmd/main.go
```
Open another terminal tab and run next commands:
```bash
wscat -c ws://localhost:8080/ws
```
Now you can send messages and get them back.

## Todo List:
1. Add message storage to the database.
2. Connect the frontend (a simple HTML page with JavaScript for WebSocket).
3. Implement a user system (registration, JWT authorization).
4. Connect multiple chats (by room ID).
package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func handleWebsocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		fmt.Println("Connection error: %v", err)
	}

	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Read error: %v", err)
			break
		}
		fmt.Println("Message received: %s", string(msg))

		// Send the message back to the client
		conn.WriteMessage(websocket.TextMessage, msg)
	}
}

func StartServer(port int) {
	r := gin.Default()

	r.GET("/ws", handleWebsocket)

	address := fmt.Sprintf(":%d", port)
	log.Printf("Server is running on %s", address)
	r.Run(address)
}

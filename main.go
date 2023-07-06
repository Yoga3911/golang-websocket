// package main

// import (
// 	"net/http"

// 	"github.com/pusher/pusher-http-go/v5"
// )

// func handler(w http.ResponseWriter, r *http.Request) {
// 	pusherClient := pusher.Client{
// 		AppID:   "1524380",
// 		Key:     "52d69074566efbf20f59",
// 		Secret:  "d76b2d94f1af7fd6f3f7",
// 		Cluster: "ap1",
// 		Secure:  true,
// 	}

// 	pusherClient.Trigger("my-channel", "my-event", "There's a new order :)")
// }

// func main() {
// 	http.HandleFunc("/", handler)

// 	http.ListenAndServe(":9999", nil)
// }

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func reader(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// print out that message for clarity
		log.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}

	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	// upgrade this connection to a WebSocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client Connected")
	err = ws.WriteMessage(1, []byte("Hi Client!"))
	if err != nil {
		log.Println(err)
	}

	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	fmt.Println("Hello World")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

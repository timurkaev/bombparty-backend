package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func handleConnetions(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Ошиька подключения", err)
		return
	}
	defer conn.Close()
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Ошибка чтения", err)
			break
		}
		log.Printf("Получено сообщение: %s", msg)
		conn.WriteMessage(websocket.TextMessage, []byte("Hello from server"))
	}
}

func main() {
	http.HandleFunc("/ws", handleConnetions)
	fmt.Println("Сервер запущен на  :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

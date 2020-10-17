package websocket

import (
	"net/http"
	"fmt"
	"time"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/azizjon12/youtube-stats/youtube"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return ws, err
	}
	return ws, err
}

func Writer(conn *websocket.Conn) {
	for {
		ticker := time.NewTicker(5 * time.Second)

		for t := range ticker.C {
			fmt.Printf("Updating Stats: %+v\n", t)
			items, err := youtube.GetSubscribers()
			if err != nil {
				fmt.Println(err)
			}

			jsonString, err := json.Marshal(items)
			if err != nil {
				fmt.Println(err)
			}

			if err := conn.WriteMessage(websocket.TextMessage, []byte(jsonString)); 
			err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}
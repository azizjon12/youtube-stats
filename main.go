package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/azizjon12/youtube-stats/websocket"
	//"github.com/azizjon12/youtube-stats/youtube"

)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func stats(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}
	go websocket.Writer(ws)
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/stats", stats)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	fmt.Println("YouTube Subscriber Monitor")
	fmt.Println("--------------------------")

	setupRoutes()

	// item, err := youtube.GetSubscribers()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("%+v\n", item)

}
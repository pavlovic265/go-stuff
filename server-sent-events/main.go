package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var msgChan chan string

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/event", sseHandler)
	router.HandleFunc("/time", getTime)

	log.Fatal(http.ListenAndServe(":8090", router))
}

func sseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	msgChan = make(chan string)
	defer func() {
		close(msgChan)
		msgChan = nil
		fmt.Println("Client close connection")
	}()

	flusher, ok := w.(http.Flusher)
	if !ok {
		fmt.Println("Could not init http Flusher")
	}

	for {
		select {
		case message := <-msgChan:
			fmt.Fprintf(w, "data: %s\n\n", message)
			flusher.Flush()
		case <-r.Context().Done():
			fmt.Println("Client close connection")
			return
		}
	}

}

func getTime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if msgChan != nil {
		msg := time.Now().Format("15:04:05")
		msgChan <- msg
	}
}

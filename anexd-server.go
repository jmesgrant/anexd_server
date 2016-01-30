package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/googollee/go-socket.io"
)

func main() {
	//nil uses ["polling", "websocket"] as default string array args
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", server)
}
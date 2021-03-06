package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"web-mpv/players"
)

const port = 6969

func IsUrlValid(s string) bool {
	_, err := url.ParseRequestURI(s)
	if err != nil {
		return false
	} else {
		return true
	}
}

func main() {
	socketPath := fmt.Sprintf("/tmp/mpv-socket-%d", rand.Int())
	logPath := fmt.Sprintf("/tmp/mpv-log-%d", rand.Int())
	clientCh := make(chan *players.MpvClient)
	go players.StartMPV(socketPath, logPath, clientCh)
	client := <-clientCh
	close(clientCh)
	for {
		if <-client.StatusCh == players.MPV_START {
			break
		}
	}
	log.Printf("Player started")

	const PORT = 6969
	go ServeHttp(client, PORT)

	select {}
}

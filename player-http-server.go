package main

import (
	"fmt"
	"log"
	"net/http"
	"web-mpv/players"
)

type httpClient struct {
	player *players.MpvClient
}

func (c *httpClient) pauseHandler(w http.ResponseWriter, r *http.Request) {
	// pause?set=true ?set=false // sets
	// pause? // toggles
	keys := r.URL.Query()["set"]
	if len(keys) > 0 {
		key := keys[0]
		if key == "true" {
			c.player.PauseSet(true)
		} else if key == "false" {
			c.player.PauseSet(false)
		} else {
			http.Error(w, "set can only be used with \"true\" or \"false\"", http.StatusBadRequest)
		}
	} else {
		c.player.PauseToggle()
	}
}
func (c *httpClient) muteHandler(w http.ResponseWriter, r *http.Request) {
	// mute?set=true ?set=false // sets
	// mute? // toggles
	keys := r.URL.Query()["set"]
	if len(keys) > 0 {
		key := keys[0]
		if key == "true" {
			c.player.MuteSet(true)
		} else if key == "false" {
			c.player.MuteSet(false)
		} else {
			http.Error(w, "set can only be used with \"true\" or \"false\"", http.StatusBadRequest)
		}
	} else {
		c.player.MuteToggle()
	}
}
func (c *httpClient) fullscreenHandler(w http.ResponseWriter, r *http.Request) {
	// fullscreen?set=true ?set=false // sets
	// fullscreen? // toggles
	keys := r.URL.Query()["set"]
	if len(keys) > 0 {
		key := keys[0]
		if key == "true" {
			c.player.FullscreenSet(true)
		} else if key == "false" {
			c.player.FullscreenSet(false)
		} else {
			http.Error(w, "set can only be used with \"true\" or \"false\"", http.StatusBadRequest)
		}
	} else {
		c.player.FullscreenToggle()
	}
}
func (c *httpClient) hideHandler(w http.ResponseWriter, r *http.Request) {
	// hide?set=true ?set=false // sets
	// TODO: hide? // toggles
	keys := r.URL.Query()["set"]
	if len(keys) > 0 {
		key := keys[0]
		if key == "true" {
			c.player.Hide()
		} else if key == "false" {
			c.player.Show()
		} else {
			http.Error(w, "set can only be used with \"true\" or \"false\"", http.StatusBadRequest)
		}
	} else {
		http.Error(w, "toggling not yet implemented", http.StatusNotImplemented)
	}
}
func (c *httpClient) playHandler(w http.ResponseWriter, r *http.Request) {
	// play?url=http://...... // plays media
	keys := r.URL.Query()["url"]
	if len(keys) > 0 {
		key := keys[0]
		if IsUrlValid(key) {
			c.player.Play(key)
		} else {
			http.Error(w, "url isn't valid", http.StatusBadRequest)
		}
	} else {
		http.Error(w, "url not set", http.StatusBadRequest)
	}
}

func ServeHttp(client *players.MpvClient, port int) {
	address := fmt.Sprintf(":%d", port)
	wrappedClient := httpClient{player: client}
	http.HandleFunc("/play", wrappedClient.playHandler)
	http.HandleFunc("/pause", wrappedClient.pauseHandler)
	http.HandleFunc("/mute", wrappedClient.muteHandler)
	http.HandleFunc("/fullscreen", wrappedClient.fullscreenHandler)
	http.HandleFunc("/hide", wrappedClient.hideHandler)
	log.Fatal(http.ListenAndServe(address, nil))
}

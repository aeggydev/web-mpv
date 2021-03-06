package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"time"
	"web-mpv/players"
)

const port = 6969

func RunMPV(socketPath string, logFile *os.File, status chan<- int) {
	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(logFile, &stdBuffer)

	cmd := exec.Command("mpv", "--idle",
		fmt.Sprintf("--input-ipc-server=%s", socketPath))
	//cmd.Stdout = mw
	cmd.Stderr = mw
	cmd.Start()

	log.Printf("CONFIG: Log file will be created in %s\n", logFile.Name())
	// TODO: Make program detect MPV starting, not just wait 5s
	time.Sleep(1 * time.Second)
	status <- 0

	err := cmd.Wait()
	log.Printf("ERROR: MPV exited with error \"%s\"\n", err)
	status <- err.(*exec.ExitError).ExitCode()
}

func IsUrlValid(s string) bool {
	_, err := url.ParseRequestURI(s)
	if err != nil {
		return false
	} else {
		return true
	}
}
func WriteJson(w http.ResponseWriter, data interface{}) {
	jsonData, _ := json.Marshal(data)
	fmt.Fprintf(w, "%s\n", jsonData)
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

	for {
	}

	/*
		socketPath := fmt.Sprintf("/tmp/mpvsocket-%d", rand.Int())
		log.Printf("CONFIG: Socket file will be created in %s\n", socketPath)

		mpvStatusCh := make(chan int)
		outputFile, _ := ioutil.TempFile("", "mpv-output-*")
		go RunMPV(socketPath, outputFile, mpvStatusCh)

		mpvStatus := <-mpvStatusCh
		if mpvStatus != 0 {
			log.Fatalf("ERROR: MPV didn't start properly\n")
		}
		log.Printf("CONFIG: MPV started\n")

		c := mpv.NewClient(mpv.NewIPCClient(socketPath))
		c.Loadfile("https://www.youtube.com/watch?v=dm-Ge-kHKjw", mpv.LoadFileModeReplace)

		http.HandleFunc("/play", func(w http.ResponseWriter, r *http.Request){
			url := r.URL.Query()["url"][0]
			if IsUrlValid(url) {
				log.Printf("REQUEST: Replacing video with \"%s\"\n", url)
				c.Loadfile(url, mpv.LoadFileModeReplace)
			} else {
				log.Printf("ERROR: URL \"%s\" isn't valid.\n", url)
			}
		})
		http.HandleFunc("/hide", func(w http.ResponseWriter, r *http.Request){
			Hide(c)
		})
		http.HandleFunc("/show", func(w http.ResponseWriter, r *http.Request){
			// TODO: Make window retain its size
			Show(c)
		})
		http.HandleFunc("/pause", func(w http.ResponseWriter, r *http.Request){
			status, _ := c.Pause()
			keys := r.URL.Query()["toggle"]
			if len(keys) > 0 {
				key := keys[0]
				if key == "true" {
					c.SetPause(true)
				} else if key == "false" {
					c.SetPause(false)
				}
				return
			}
			// TODO: Send back error message
			c.SetPause(!status)
			log.Printf("MEDIA: Paused was toggled to %t\n", !status)
		})
		http.HandleFunc("/mute", func(w http.ResponseWriter, r *http.Request){
			status, _ := c.Mute()
			keys := r.URL.Query()["toggle"]
			if len(keys) > 0 {
				key := keys[0]
				if key == "true" {
					c.SetMute(true)
				} else if key == "false" {
					c.SetMute(false)
				}
				return
			}
			// TODO: Send back error message
			c.SetMute(!status)
			log.Printf("MEDIA: Muted was toggled to %t\n", !status)
		})
		http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request){
			duration, _ := c.Duration()
			position, _ := c.Position()
			paused, _ := c.Pause()
			muted, _ := c.Mute()
			data := Status{
				Muted:    muted,
				Paused:   paused,
				Position: position,
				Duration: duration,
			}
			WriteJson(w, data)
		})

		stopProgram := make(chan os.Signal)
		signal.Notify(stopProgram, syscall.SIGTERM)
		signal.Notify(stopProgram, syscall.SIGINT)
		signal.Notify(stopProgram, syscall.SIGHUP)
		go func(){
			sig := <-stopProgram
			log.Printf("ERROR: Caught signal: %+v\n", sig)
			log.Printf("EXIT: Cleaning up temporary files: %s, %s\n", socketPath, outputFile.Name())
			os.Remove(socketPath)
			os.Remove(outputFile.Name())
		}()

		go func(){
			for {
				select {
				case mpvStatus := <-mpvStatusCh:
					log.Printf("ERROR: MPV exited with code %d, closing program\n", mpvStatus)
					stopProgram <- syscall.SIGHUP
					time.Sleep(1 * time.Second)
					os.Exit(1)
				}
			}
		}()

		address := fmt.Sprintf(":%d", port)
		log.Printf("CONFIG: Starting server on %s.\n", address)
		if err := http.ListenAndServe(address, nil); err != nil {
			log.Fatal(err)
		}
	*/
}

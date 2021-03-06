package players

import (
	"fmt"
	"github.com/blang/mpv"
	"os/exec"
	"time"
)

const (
	MPV_START = iota
	MPV_EXIT
)

type MpvClient struct {
	socketPath string
	logPath    string
	client     *mpv.Client

	StatusCh chan int // MPV_... status codes are thrown here
}

func (c *MpvClient) PauseGet() (bool, error) {
	return c.client.Pause()
}
func (c *MpvClient) PauseSet(value bool) error {
	return c.client.SetPause(value)
}
func (c *MpvClient) PauseToggle() (bool, error) {
	current, err := c.PauseGet()
	if err != nil {
		return current, err
	}

	return !current, c.PauseSet(!current)
}

func (c *MpvClient) MuteGet() (bool, error) {
	return c.client.Mute()
}
func (c *MpvClient) MuteSet(value bool) error {
	return c.client.SetMute(value)
}
func (c *MpvClient) MuteToggle() (bool, error) {
	current, err := c.MuteGet()
	if err != nil {
		return current, err
	}

	return !current, c.MuteSet(!current)
}

func (c *MpvClient) FullscreenGet() (bool, error) {
	return c.client.Fullscreen()
}
func (c *MpvClient) FullscreenSet(value bool) error {
	return c.client.SetFullscreen(value)
}
func (c *MpvClient) FullscreenToggle() (bool, error) {
	current, err := c.FullscreenGet()
	if err != nil {
		return current, err
	}

	return !current, c.FullscreenSet(!current)
}

func (c *MpvClient) Position() (Position, error) {
	data := Position{}
	currentIdle, err := c.client.Idle()
	if err != nil {
		return data, err
	}

	data.OnMedia = !currentIdle
	duration, err := c.client.Duration()
	if err != nil {
		return data, err
	}
	position, err := c.client.Position()
	if err != nil {
		return data, err
	}
	data.Duration = duration
	data.Position = position
	return data, nil
}
func (c *MpvClient) Play(path string) error {
	return c.client.Loadfile(path, mpv.LoadFileModeReplace)
}
func (c *MpvClient) Hide() error {
	return c.client.SetProperty("video", "no")
}
func (c *MpvClient) Show() error {
	return c.client.SetProperty("video", "auto")
}

func StartMPV(socketPath string, logPath string, clientCh chan *MpvClient) {
	client := MpvClient{socketPath: socketPath, logPath: logPath, StatusCh: make(chan int)}
	clientCh <- &client

	cmd := exec.Command("mpv", "--idle",
		fmt.Sprintf("--input-ipc-server=%s", client.socketPath))
	cmd.Start()

	// TODO: Find a better way to find out if Mpv is started
	time.Sleep(2 * time.Second)
	client.StatusCh <- MPV_START

	client.client = mpv.NewClient(mpv.NewIPCClient(client.socketPath))

	cmd.Wait()
	client.StatusCh <- MPV_EXIT
}

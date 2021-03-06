package players

type MediaPlayer interface {
	PauseGet() (bool, error)
	PauseSet(value bool) error
	PauseToggle() (bool, error) // Returns what the property was set to

	MuteGet() (bool, error)
	MuteSet(value bool) error
	MuteToggle() (bool, error) // Returns what the property was set to

	FullscreenGet() (bool, error)
	FullscreenSet(value bool) error
	FullscreenToggle() (bool, error)

	Position() (Position, error)
	Play(path string) error
}
type Position struct {
	OnMedia  bool    // Is something playing
	Position float64 // Current position in seconds
	Duration float64 // Full length in seconds
}

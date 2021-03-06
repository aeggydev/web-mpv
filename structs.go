package main

type MutedResponse struct {
	Muted bool `json:"muted"`
}
type PausedResponse struct {
	Paused bool `json:"paused"`
}
type PositionResponse struct {
	Position float64 `json:"position"`
	Duration float64 `json:"duration"`
}
type Status struct {
	Muted  bool `json:"muted"`
	Paused bool `json:"paused"`

	Position float64 `json:"position"`
	Duration float64 `json:"duration"`
}

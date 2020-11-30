package game

const (
	Kickoff = iota
	Playing
)

type PlayerStateComponent struct {
	State int
}

func NewPlayerStateComponent() *PlayerStateComponent {
	return &PlayerStateComponent{State: Kickoff}
}

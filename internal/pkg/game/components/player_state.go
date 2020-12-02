package components

import "github.com/kevholditch/breakout/internal/pkg/ecs"

const (
	Kickoff = iota
	Playing
)

var HasPlayingState *PlayingState

type PlayingState interface {
	PlayingStateComponent() *PlayerStateComponent
}

type PlayerStateComponent struct {
	State int
}

func NewPlayerStateComponent() *PlayerStateComponent {
	return &PlayerStateComponent{State: Kickoff}
}

func (s *PlayerStateComponent) PlayingStateComponent() *PlayerStateComponent {
	return s
}

func init() {
	ecs.RegisterComponent(&PlayerStateComponent{})
}

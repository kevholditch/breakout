package game

const (
	Kickoff = iota
	Playing
)

type GameState struct {
	State int
}

func NewGameState() *GameState {
	return &GameState{State: Kickoff}
}

package game

type renderPlayer struct {
	player renderEntityHolder
}

func NewPlayerRenderer(player renderEntityHolder) *renderPlayer {
	return &renderPlayer{player: player}
}

func (r *renderPlayer) initialise() {

}

func (r *renderPlayer) render() {

}

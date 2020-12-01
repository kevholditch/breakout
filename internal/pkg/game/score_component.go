package game

type ScoreComponent struct {
	Score int
}

func NewScoreComponent(score int) *ScoreComponent {
	return &ScoreComponent{Score: score}
}

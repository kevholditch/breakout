package game

var IsSpeedControllable *SpeedControllable

type SpeedControllable interface {
	SpeedControllableComponent() *SpeedControlComponent
}

type SpeedControlComponent struct {
	Speed float32
}

func NewControlComponent(initialSpeed float32) *SpeedControlComponent {
	return &SpeedControlComponent{Speed: initialSpeed}
}

func (c *SpeedControlComponent) SpeedControllableComponent() *SpeedControlComponent {
	return c
}

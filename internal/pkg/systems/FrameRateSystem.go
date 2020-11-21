package systems

import (
	"fmt"
	"github.com/EngoEngine/ecs"
)

type FrameRateSystem struct {
	FrameRate         float32
	syncEveryXUpdates int
	sum               float32
	count             int
}

func NewFrameRateSystem() *FrameRateSystem {
	return &FrameRateSystem{
		syncEveryXUpdates: 50,
	}
}

func (f *FrameRateSystem) Reset() {
	f.sum = 0
	f.count = 0
}

func (f *FrameRateSystem) Update(dt float32) {
	f.sum += 1 / dt
	f.count++

	if f.count == f.syncEveryXUpdates {
		f.FrameRate = f.sum / float32(f.syncEveryXUpdates)
		fmt.Printf("framerate is: %.1f\n", f.FrameRate)
		f.Reset()
	}

}

func (f *FrameRateSystem) Remove(e ecs.BasicEntity) {

}

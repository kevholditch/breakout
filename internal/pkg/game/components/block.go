package components

import (
	"github.com/liamg/ecs"
)

var IsBlock *Block

type Block interface {
	BlockComponent() *BlockComponent
}

type BlockComponent struct{}

func NewBlockComponent() *BlockComponent {
	return &BlockComponent{}
}

func (c *BlockComponent) BlockComponent() *BlockComponent {
	return c
}

func init() {
	ecs.RegisterComponent(&BlockComponent{})
}

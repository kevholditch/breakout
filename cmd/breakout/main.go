package main

import (
	"fmt"
	"github.com/kevholditch/breakout/internal/pkg/game"
	"os"
)

func main() {
	err := game.Run()
	if err != nil {
		fmt.Printf("error running breakout: %v", err)
		os.Exit(-1)
	}
}

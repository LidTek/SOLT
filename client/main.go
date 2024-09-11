package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := NewGame()

	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Embedded Client")

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}

package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Embedded Client")

	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}

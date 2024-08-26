package main

import "github.com/hajimehoshi/ebiten/v2"

// State represents the state of the game, waoow
type State interface {
	// Enter is called when first entering the state.
	Enter(g *Game)
	// Leave is called when leaving the state.
	Leave(g *Game)
	// Update is update.
	Update(g *Game) error
	// Draw is y'know.
	Draw(g *Game, screen *ebiten.Image)
}

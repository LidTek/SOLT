package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// Game is our overall structure for Ebitengine.
type Game struct {
	bigFont   *text.GoTextFace
	state     State
	nextState State
}

// NewGame creates a new game and sets the first state.
func NewGame() *Game {
	g := &Game{}
	g.NextState(&StatePending{})
	return g
}

// NextState sets the state to switch to at the end of the current tick.
func (g *Game) NextState(state State) {
	g.nextState = state
}

// Update updates the current state and then switches to the next state if one is pending.
func (g *Game) Update() error {
	if g.state != nil {
		return g.state.Update(g)
	}
	if g.nextState != nil {
		if g.state != nil {
			g.state.Leave(g)
		}
		g.state = g.nextState
		g.nextState = nil
		g.state.Enter(g)
	}
	return nil
}

// Draw draws the current state if one exists.
func (g *Game) Draw(screen *ebiten.Image) {
	if g.state != nil {
		g.state.Draw(g, screen)
	}
}

// Layout just sets the same size as outside dimensions.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

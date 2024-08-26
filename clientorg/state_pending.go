package main

import (
	"image/color"

	"github.com/LidTek/client/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// StatePending represents the state when awaiting to connect to... peers? a room?
type StatePending struct {
	ticks   int
	message TickMessage
}

// Enter is called when first entering the state.
func (s *StatePending) Enter(g *Game) {
	s.message = TickMessage{
		ticks:    90,
		messages: []string{"Pending", "Pending.", "Pending..", "Pending..."},
	}
}

// Leave is called when leaving the state.
func (s *StatePending) Leave(g *Game) {
}

// Update is update.
func (s *StatePending) Update(g *Game) error {
	s.ticks++
	s.message.Update()
	return nil
}

// Draw draws the pending text to the screen.
func (s *StatePending) Draw(g *Game, screen *ebiten.Image) {
	screen.Fill(color.White)
	opts := &text.DrawOptions{}
	opts.LayoutOptions.SecondaryAlign = text.AlignCenter
	opts.ColorScale.ScaleWithColor(color.Black)

	w, h := text.Measure("Pending  ", assets.FontGroups[assets.DisplayFont].GetFace(text.WeightBold), 0)

	opts.GeoM.Translate(float64(screen.Bounds().Dx())/2-w/2, float64(screen.Bounds().Dy())/2-h/2)

	face := assets.FontGroups[assets.DisplayFont].GetFace(text.WeightBold)

	text.Draw(screen, s.message.Message(), face, opts)
}

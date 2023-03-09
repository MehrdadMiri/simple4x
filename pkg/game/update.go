package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Update() error {
	g.Debugger.Set("fps", ebiten.ActualFPS())
	x, y := ebiten.CursorPosition()
	g.Debugger.Set("cursor", fmt.Sprintf("%v %v", x, y))
	select {
	case <-g.TickerChan:
		g.Debugger.Inc("clock")
	default:
		break
	}
	return nil
}

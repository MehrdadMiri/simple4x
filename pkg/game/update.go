package game

import (
	"fmt"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Update() error {
	g.Debugger.Set("fps", ebiten.ActualFPS())
	x, y := ebiten.CursorPosition()
	g.Debugger.Set("cursor", fmt.Sprintf("%v %v", x, y))
	select {
	case <-g.TickerChan:
		i := rand.Intn(g.Rows)
		j := rand.Intn(g.Columns)
		g.Grid.Select(i, j)
		i = rand.Intn(g.Rows)
		j = rand.Intn(g.Columns)
		g.Grid.Unselect(i, j)
		g.Debugger.Inc("clock")
	default:
		break
	}
	return nil
}

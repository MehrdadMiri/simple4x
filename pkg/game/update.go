package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Update() error {
	g.Debugger.Set("fps", ebiten.ActualFPS())
	x, y := ebiten.CursorPosition()
	g.Debugger.Set("cursor", fmt.Sprintf("%v %v", x, y))
	cell, err := g.Grid.FindCell(x, y)
	if err != nil {
		g.Debugger.Set("out of bounds", "true")
	} else {
		g.Debugger.Set("out of bounds", "false")
		g.Debugger.Set("cell", fmt.Sprintf("%v %v", cell.Row(), cell.Column()))
		g.Grid.UnselectAll()
		cell.Select()
		adjacents, err := g.Grid.GetAdjacantHexes(cell.Row(), cell.Column())
		if err != nil {
			return err
		}
		for _, hex := range adjacents {
			hex.Select()
		}
	}
	g.Debugger.Inc("clock")
	select {
	case <-g.TickerChan:
	default:
		break
	}
	return nil
}

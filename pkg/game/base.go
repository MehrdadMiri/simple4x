package game

import (
	"time"

	"tidy/pkg/debugger"
	"tidy/pkg/grid"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Rows, Columns int
	Grid          grid.Grid
	TickerChan    <-chan time.Time
	Debugger      *debugger.Debugger
}

func NewGame(rows, columns int) *Game {
	_, height := ebiten.WindowSize()
	g := &Game{
		Rows:       rows,
		Columns:    columns,
		TickerChan: time.NewTicker(200 * time.Millisecond).C,
		Grid:       grid.NewGrid(rows, columns, height),
		Debugger:   debugger.NewDebugger(),
	}
	g.Debugger.Set("clock", 0)
	return g
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

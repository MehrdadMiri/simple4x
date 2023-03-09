package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func (g *Game) Draw(screen *ebiten.Image) {
	g.Grid.Draw(screen, false)
	ebitenutil.DebugPrint(screen, g.Debugger.ToString())
}

package game

import (
	"tidy/pkg/configs"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func (g *Game) Draw(screen *ebiten.Image) {
	g.Grid.Draw(screen, configs.Debug)
	ebitenutil.DebugPrint(screen, g.Debugger.ToString())
}

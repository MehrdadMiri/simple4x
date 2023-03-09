package game

import (
	"math"
	"tidy/pkg/configs"
	"tidy/pkg/geometry"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func (g *Game) Draw(screen *ebiten.Image) {
	_, h := ebiten.WindowSize()
	hexes := getHexes(g.Columns, g.Rows, configs.HexRad, configs.OffsetH, configs.OffsetV, h)
	for _, hex := range hexes {
		hex.DrawLines(screen, configs.Debug)
	}
	ebitenutil.DebugPrint(screen, g.Debugger.ToString())
}

func getHexes(columns, rows, rad, offsetH, offsetV, height int) []*geometry.Hex {
	length := math.Sin(math.Pi/3) * float64(rad)
	hexes := make([]*geometry.Hex, 0, rows*columns)
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			var x, y float64
			if i%2 == 0 {
				x = float64(j)*2*length + float64(offsetH)/2
			} else {
				x = float64(j)*2*length + length + float64(offsetH)/2
			}
			y = float64(height) - (float64(i)*1.5*float64(rad) + float64(offsetV)/2)
			hexes = append(hexes, geometry.NewHex(x, y, float64(rad), i, j))
		}
	}
	return hexes
}

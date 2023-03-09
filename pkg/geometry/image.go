package geometry

import (
	"fmt"
	"image/color"
	"tidy/pkg/configs"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func (h Hex) DrawLines(image *ebiten.Image, color color.Color, debug bool) {
	for _, line := range h.GetLines() {
		x1, y1 := line.Start().XY()
		x2, y2 := line.End().XY()
		ebitenutil.DrawLine(image, x1, y1, x2, y2, color)
	}
	if debug {
		ebitenutil.DebugPrintAt(image, fmt.Sprintf("%v,%v", h.row, h.column), int(h.x)-configs.HexRad/2, int(h.y))
	}
}

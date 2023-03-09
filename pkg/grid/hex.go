package grid

import (
	"image/color"
	"tidy/pkg/configs"
	"tidy/pkg/geometry"

	"github.com/hajimehoshi/ebiten/v2"
)

type Hex struct {
	hex    *geometry.Hex
	color  color.Color
	column int
	row    int
	center geometry.Point
}

func NewHex(hex *geometry.Hex, color color.Color, row, column int, center geometry.Point) Hex {
	return Hex{
		hex:    hex,
		color:  color,
		row:    row,
		column: column,
		center: center,
	}
}

// Draws hexes with color
func (h Hex) Draw(image *ebiten.Image, debug bool) {
	h.hex.DrawLines(image, h.color, debug)
}

// Sets color to selected color
func (h *Hex) Select() {
	h.color = configs.SelectedColor
}

// Sets color to unselected color
func (h *Hex) Unselect() {
	h.color = configs.DefaultColor
}

func (h *Hex) Hex() *geometry.Hex {
	return h.hex
}

func (h *Hex) Color() color.Color {
	return h.color
}

func (h *Hex) Column() int {
	return h.column
}

func (h *Hex) Row() int {
	return h.row
}

func (h *Hex) Center() geometry.Point {
	return h.center
}

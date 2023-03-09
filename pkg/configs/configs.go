package configs

import "image/color"

const (
	// HexRad is the radius of a hexagon.
	HexRad = 5
	// Debug is a flag to enable debug mode.
	Debug = true
	// Offset is the difference from buttom of the screen and the first hexagon.
	OffsetH = 50
	OffsetV = 50
)

var (
	DefaultColor  = color.Gray{100}
	SelectedColor = color.RGBA{0, 0, 255, 255}
)

package grid

import (
	"image/color"
	"math"
	"tidy/pkg/configs"
	"tidy/pkg/geometry"

	"github.com/hajimehoshi/ebiten/v2"
)

type Grid struct {
	// contains filtered or unexported fields
	hexes   [][]Hex
	rows    int
	columns int
}

func NewGrid(rows, columns, height int) Grid {
	return Grid{
		hexes:   getHexes(columns, rows, configs.HexRad, configs.OffsetH, configs.OffsetV, height),
		rows:    rows,
		columns: columns,
	}
}

func (g *Grid) Select(row, column int) {
	g.hexes[row][column].Select()
}

func (g *Grid) Unselect(row, column int) {
	g.hexes[row][column].Unselect()
}

func getHexes(columns, rows, rad, offsetH, offsetV, height int) [][]Hex {
	length := math.Sin(math.Pi/3) * float64(rad)
	hexes := make([][]Hex, rows)
	for i := range hexes {
		hexes[i] = make([]Hex, columns)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			var x, y float64
			if i%2 == 0 {
				x = float64(j)*2*length + float64(offsetH)/2
			} else {
				x = float64(j)*2*length + length + float64(offsetH)/2
			}
			y = float64(height) - (float64(i)*1.5*float64(rad) + float64(offsetV)/2)
			hexes[i][j] = NewHex(
				geometry.NewHex(x, y, float64(rad), i, j),
				color.White,
				i, j,
				*geometry.NewPoint(x, y),
			)
			//append(hexes, geometry.NewHex(x, y, float64(rad), i, j
		}
	}
	return hexes
}

func (g Grid) Draw(image *ebiten.Image, debug bool) {
	for _, row := range g.hexes {
		for _, hex := range row {
			hex.Draw(image, debug)
		}
	}
}

package grid

import (
	"errors"
	"image/color"
	"math"
	"tidy/pkg/configs"
	"tidy/pkg/geometry"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	// errors
	ErrInvalidHex  = "invalid hex"
	ErrOutOfBounds = "out of bounds"
)

type Grid struct {
	// contains filtered or unexported fields
	hexes   [][]Hex
	rows    int
	columns int
	offsetH int
	offsetV int
}

func NewGrid(rows, columns, height int) Grid {
	return Grid{
		hexes:   getHexes(columns, rows, configs.HexRad, configs.OffsetH, configs.OffsetV, height),
		rows:    rows,
		columns: columns,
		offsetH: configs.OffsetH,
		offsetV: configs.OffsetV,
	}
}

func (g *Grid) Select(row, column int) {
	g.hexes[row][column].Select()
}

func (g *Grid) Unselect(row, column int) {
	g.hexes[row][column].Unselect()
}

func (g *Grid) UnselectAll() {
	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.columns; j++ {
			g.hexes[i][j].Unselect()
		}
	}
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

func (g *Grid) Draw(image *ebiten.Image, debug bool) {
	for _, row := range g.hexes {
		for _, hex := range row {
			hex.Draw(image, debug)
		}
	}
}

func (g *Grid) FindCell(x, y int) (*Hex, error) {
	minx := g.offsetH/2 - int(math.Cos(math.Pi/6)*configs.HexRad)
	miny := g.offsetV/2 + configs.HexRad
	maxx := int(2*math.Cos(math.Pi/6)*configs.HexRad*float64(g.columns) + configs.OffsetH/2)
	maxy := int(1.5*configs.HexRad*float64(g.rows)+configs.OffsetV/2) + configs.HexRad
	if float64(x) < float64(minx) || float64(x) > float64(maxx) || float64(y) < float64(miny) || float64(y) > float64(maxy) {
		return nil, errors.New(ErrOutOfBounds)
	}
	i := int(math.Floor((float64(maxy) - float64(y)) / (1.5 * configs.HexRad)))
	j := int(math.Floor((float64(x) - float64(minx)) / (2 * math.Cos(math.Pi/6) * configs.HexRad)))
	if i < 0 {
		i = 0
	} else if i > g.rows-1 {
		i = g.rows - 1
	}
	if j < 0 {
		j = 0
	} else if j > g.columns-1 {
		j = g.columns - 1
	}
	dis := geometry.Distance(*geometry.NewPoint(float64(x), float64(y)), g.hexes[i][j].center)
	adj, err := g.GetAdjacantHexes(i, j)
	if err != nil {
		return nil, err
	}
	for _, hex := range adj {
		if geometry.Distance(*geometry.NewPoint(float64(x), float64(y)), hex.center) < dis {
			dis = geometry.Distance(*geometry.NewPoint(float64(x), float64(y)), hex.center)
			i = hex.row
			j = hex.column
		}
	}
	return &g.hexes[i][j], nil
}

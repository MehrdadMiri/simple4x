package grid

import "errors"

var indicesEven = [][]int{
	{0, -1},
	{0, 1},
	{-1, -1},
	{-1, 0},
	{1, -1},
	{1, 0},
}

var indicesOdd = [][]int{
	{0, -1},
	{0, 1},
	{-1, 0},
	{-1, 1},
	{1, 0},
	{1, 1},
}

func (g *Grid) GetAdjacantHexes(i, j int) ([]*Hex, error) {
	if i < 0 || i >= g.rows || j < 0 || j >= g.columns {
		return nil, errors.New("invalid index")
	}
	var adjacents []*Hex
	if i%2 == 0 {
		for ind := range indicesEven {
			i, j := i+indicesEven[ind][0], j+indicesEven[ind][1]
			if i < 0 || i >= g.rows {
				continue
			}
			j = (j + g.columns) % g.columns
			adjacents = append(adjacents, &g.hexes[i][j])
		}
	} else {
		for ind := range indicesOdd {
			i, j := i+indicesOdd[ind][0], j+indicesOdd[ind][1]
			if i < 0 || i >= g.rows {
				continue
			}
			j = (j + g.columns) % g.columns
			adjacents = append(adjacents, &g.hexes[i][j])
		}
	}
	return adjacents, nil
}

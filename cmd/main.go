package main

import (
	"fmt"
	"math"
	"tidy/pkg/configs"
	"tidy/pkg/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	fmt.Println("Hello, playground")
	n, m := 80, 20
	w := int(2*math.Cos(math.Pi/6)*configs.HexRad*float64(m) + configs.OffsetH)
	h := int(1.5*float64(n)*configs.HexRad + configs.OffsetV)
	ebiten.SetWindowSize(w, h)
	ebiten.SetWindowTitle("Tidy")
	g := game.NewGame(n, m)
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}

package entities

import (
	"github.com/hajimehoshi/ebiten"
	"snake-ebiten/util"
)

type Point struct {
	X int
	Y int
}

func (p Point) Xf() float64 {
	return float64(p.X * util.GridSize)
}

func (p Point) Yf() float64 {
	return float64(p.Y * util.GridSize)
}

type Object interface {
	Update() error
	Render(screen *ebiten.Image) error
}

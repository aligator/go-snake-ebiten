package entities

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"log"
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

func (p Point) Equals(p2 Point) bool {
	return p.X == p2.X && p.Y == p2.Y
}

type Object interface {
	Update() error
	Render(screen *ebiten.Image) error
}

func mustLoadTexture(path string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(path, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	return img
}

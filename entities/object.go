package entities

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"log"
	"snake-ebiten/util"
)

type Point struct {
	X float64
	Y float64
}

func (p Point) GridX() int {
	return int(p.X / util.GridSize)
}

func (p Point) GridY() int {
	return int(p.Y / util.GridSize)
}

func (p *Point) IncGridX() {
	p.X += util.GridSize
}

func (p *Point) IncGridY() {
	p.Y += util.GridSize
}

func (p *Point) DecGridX() {
	p.X -= util.GridSize
}

func (p *Point) DecGridY() {
	p.Y -= util.GridSize
}

func NewGridPoint(point Point) Point {
	return Point{
		X: point.X * util.GridSize,
		Y: point.Y * util.GridSize,
	}
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

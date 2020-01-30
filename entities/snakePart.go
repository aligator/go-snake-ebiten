package entities

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
)

type PartType int

const (
	Head = iota
	Body
	Tail
)

var imgHead, imgBody, imgTail *ebiten.Image

func init() {
	fmt.Println("Init SnakePart")

	imgHead = mustLoadTexture("assets/textures/snake_front.png")
	imgBody = mustLoadTexture("assets/textures/snake.png")
	imgTail = mustLoadTexture("assets/textures/snake_back.png")
}

func resolveImage(t PartType) *ebiten.Image {
	switch t {
	case Head:
		return imgHead
	case Body:
		return imgBody
	case Tail:
		return imgTail
	default:
		return nil
	}
}

type SnakePart struct {
	Position Point
	Type     PartType
}

func (s SnakePart) Update() error {
	return nil
}

func (s SnakePart) Render(screen *ebiten.Image) error {
	opt := ebiten.DrawImageOptions{}
	opt.GeoM.Translate(s.Position.Xf(), s.Position.Yf())

	screen.DrawImage(resolveImage(s.Type), &opt)

	return nil
}

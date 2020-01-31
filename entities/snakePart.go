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

func resolveSnakeImage(t PartType) *ebiten.Image {
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
	position Point
	partType PartType
	isEating bool
}

func (s SnakePart) Update() error {
	// todo: smooth move
	return nil
}

func (s SnakePart) Render(screen *ebiten.Image) error {
	opt := ebiten.DrawImageOptions{}
	opt.GeoM.Translate(s.position.Xf(), s.position.Yf())

	screen.DrawImage(resolveSnakeImage(s.partType), &opt)
	return nil
}

package entities

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"snake-ebiten/util"
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
	position       Point
	movingTo       Point
	movingProgress float64
	partType       PartType
	isEating       bool
}

func (s *SnakePart) SetPos(pos Point) {
	s.position = pos
	s.movingTo = pos
	s.movingProgress = 0
}

func (s *SnakePart) Move(dir Direction) {
	s.movingTo = dir.translate(s.position)
	s.movingProgress = 0
}

func (s *SnakePart) IsMoving() bool {
	return s.position != s.movingTo
}

func (s *SnakePart) Update() error {
	if !s.IsMoving() {
		return nil
	}
	// calculate grid step
	step := float64(util.GridSize) / float64(util.SnakeSpeed-1)

	// move by step
	s.movingProgress = s.movingProgress + step

	// if moving progress is more or same than util.GridSize, set position to movingTo
	// this ends the moving animation
	if s.movingProgress >= util.GridSize {
		s.position = s.movingTo
	}
	return nil
}

func (s *SnakePart) Render(screen *ebiten.Image) error {
	opt := ebiten.DrawImageOptions{}
	opt.GeoM.Translate(s.position.X, s.position.Y)

	if s.IsMoving() {
		opt.GeoM.Translate(
			float64(s.movingTo.GridX()-s.position.GridX())*s.movingProgress,
			float64(s.movingTo.GridY()-s.position.GridY())*s.movingProgress)
	}

	screen.DrawImage(resolveSnakeImage(s.partType), &opt)
	return nil
}

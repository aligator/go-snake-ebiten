package entities

import (
	"github.com/hajimehoshi/ebiten"
	"snake-ebiten/util"
)

type direction int

const (
	up direction = iota + 1
	down
	left
	right
)

type Snake struct {
	parts   []SnakePart
	lastDir direction

	updateCount uint
}

func NewSnake() *Snake {
	center := Point{
		X: util.GridWidth / 2,
		Y: util.GridHeight / 2,
	}
	s := Snake{parts: []SnakePart{
		{
			Position: Point{
				X: center.X - 1,
				Y: center.Y,
			},
			Type: Tail,
		},
		{
			Position: center,
			Type:     Body,
		},
		{
			Position: Point{
				X: center.X + 1,
				Y: center.Y,
			},
			Type: Head,
		},
	}}

	return &s
}

func (s *Snake) Update() error {
	switch {
	case ebiten.IsKeyPressed(ebiten.KeyUp):
		s.lastDir = up
	case ebiten.IsKeyPressed(ebiten.KeyDown):
		s.lastDir = down
	case ebiten.IsKeyPressed(ebiten.KeyLeft):
		s.lastDir = left
	case ebiten.IsKeyPressed(ebiten.KeyRight):
		s.lastDir = right
	}

	if s.lastDir != 0 {

		if s.updateCount == 30 {
			s.updateCount = 0
			newHead := SnakePart{
				Position: s.parts[len(s.parts)-1].Position,
				Type:     Head,
			}

			switch s.lastDir {
			case up:
				newHead.Position.Y--
			case down:
				newHead.Position.Y++
			case left:
				newHead.Position.X--
			case right:
				newHead.Position.X++
			}

			s.parts = append(s.parts, newHead)

			s.parts[len(s.parts)-2].Type = Body

			s.parts = append(s.parts[:0], s.parts[0+1:]...)

			s.parts[0].Type = Tail
		}

		s.updateCount++
	}

	for _, p := range s.parts {
		err := p.Update()
		if err != nil {
			return err
		}
	}
	return nil
}

func (s Snake) Render(screen *ebiten.Image) error {
	for _, o := range s.parts {
		err := o.Render(screen)
		if err != nil {
			return err
		}
	}
	return nil
}

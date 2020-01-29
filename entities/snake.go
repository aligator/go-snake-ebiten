package entities

import (
	"github.com/hajimehoshi/ebiten"
	"snake-ebiten/util"
)

type Snake struct {
	parts []SnakePart
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

func (s Snake) Update() error {
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

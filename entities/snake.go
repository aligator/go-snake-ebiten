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
	game    *Game
	parts   []SnakePart
	lastDir direction

	updateCount uint
}

func NewSnake(g *Game) *Snake {
	center := Point{
		X: util.GridWidth / 2,
		Y: util.GridHeight / 2,
	}
	s := Snake{
		game: g,
		parts: []SnakePart{
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
	// check keys
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
		if s.updateCount == util.SnakeSpeed {
			s.updateCount = 0

			// create new head
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

			// check if head collides with something
			// 1. with borders
			if newHead.Position.X > util.GridWidth-1 || newHead.Position.X < 0 ||
				newHead.Position.Y > util.GridHeight-1 || newHead.Position.Y < 0 {
				s.game.End()
			} else {
				// 2. with snake itself
				for _, part := range s.parts {
					if newHead.Position.Equals(part.Position) {
						s.game.End()
						break
					}
				}
			}

			// check for collision with cooky
			if s.game.cooky.position.Equals(newHead.Position) {
				s.game.cooky.Respawn()
			}

			// no need to move snake if game ended
			if !s.game.IsRunning() {
				return nil
			}

			s.parts = append(s.parts, newHead)

			// transform old head to body
			s.parts[len(s.parts)-2].Type = Body
			// remove old tail
			s.parts = append(s.parts[:0], s.parts[0+1:]...)
			// transform new tail to tail
			s.parts[0].Type = Tail
		}

		s.updateCount++
	}

	// call update for all snake parts
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

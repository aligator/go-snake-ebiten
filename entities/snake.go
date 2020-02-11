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
	center := NewGridPoint(Point{
		X: util.GridWidth / 2,
		Y: util.GridHeight / 2,
	})
	s := Snake{
		game: g,
		parts: []SnakePart{
			{
				position: Point{
					X: center.X - util.GridSize,
					Y: center.Y,
				},
				partType: Tail,
			},
			{
				position: center,
				partType: Body,
			},
			{
				position: Point{
					X: center.X + util.GridSize,
					Y: center.Y,
				},
				partType: Head,
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
				position: s.parts[len(s.parts)-1].position,
				partType: Head,
			}
			switch s.lastDir {
			case up:
				newHead.position.DecGridY()
			case down:
				newHead.position.IncGridY()
			case left:
				newHead.position.DecGridX()
			case right:
				newHead.position.IncGridX()
			}

			// check if head collides with something
			// 1. with borders
			if newHead.position.GridX() > util.GridWidth-1 || newHead.position.GridX() < 0 ||
				newHead.position.GridY() > util.GridHeight-1 || newHead.position.GridY() < 0 {
				s.game.End()
			} else {
				// 2. with snake itself
				for _, part := range s.parts {
					if newHead.position == part.position {
						s.game.End()
						break
					}
				}
			}

			// check for collision with cooky
			if s.game.cooky.position == newHead.position {
				s.game.cooky.respawn()
				s.game.incScore()

				// eat
				newHead.isEating = true
			}

			// no need to move snake if game ended
			if !s.game.IsRunning() {
				return nil
			}

			s.parts = append(s.parts, newHead)

			// transform old head to body
			s.parts[len(s.parts)-2].partType = Body

			// remove tail only if no need to grow
			if !s.parts[0].isEating {
				// remove old tail
				s.parts = append(s.parts[:0], s.parts[0+1:]...)
			} else {
				// stop eating so that in the next iteration the tail will be removed
				s.parts[0].isEating = false
			}

			// transform new tail to tail
			s.parts[0].partType = Tail
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

package entities

import (
	"github.com/hajimehoshi/ebiten"
	"snake-ebiten/util"
)

type Snake struct {
	game    *Game
	parts   []*SnakePart
	lastDir Direction

	updateCount uint
}

func NewSnake(g *Game) *Snake {
	center := NewGridPoint(Point{
		X: util.GridWidth / 2,
		Y: util.GridHeight / 2,
	})
	s := Snake{
		game: g,
		parts: []*SnakePart{
			{
				position: Point{
					X: center.X - util.GridSize,
					Y: center.Y,
				},
				movingTo: Point{
					X: center.X - util.GridSize,
					Y: center.Y,
				},
				partType: Tail,
			},
			{
				position: center,
				movingTo: center,
				partType: Body,
			},
			{
				position: Point{
					X: center.X + util.GridSize,
					Y: center.Y,
				},
				movingTo: Point{
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
		s.lastDir = Up
	case ebiten.IsKeyPressed(ebiten.KeyDown):
		s.lastDir = Down
	case ebiten.IsKeyPressed(ebiten.KeyLeft):
		s.lastDir = Left
	case ebiten.IsKeyPressed(ebiten.KeyRight):
		s.lastDir = Right
	}

	if s.lastDir != 0 {
		if s.updateCount == util.SnakeSpeed {
			s.updateCount = 0

			// create new head
			newHead := &SnakePart{
				partType: Head,
			}

			newHead.SetPos(s.parts[len(s.parts)-1].position)
			newHead.Move(s.lastDir)

			// check if head collides with something
			// 1. with borders
			if newHead.movingTo.GridX() > util.GridWidth-1 || newHead.movingTo.GridX() < 0 ||
				newHead.movingTo.GridY() > util.GridHeight-1 || newHead.movingTo.GridY() < 0 {
				s.game.End()
			} else {
				// 2. with snake itself
				for _, part := range s.parts {
					if newHead.movingTo == part.position {
						s.game.End()
						break
					}
				}
			}

			// check for collision with cooky
			if s.game.cooky.position == newHead.movingTo {
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

			// translate old head to body
			s.parts[len(s.parts)-2].partType = Body

			// remove tail only if no need to grow
			if !s.parts[0].isEating {
				// remove old tail (but only after it finished smooth-moving)
				if !s.parts[0].IsMoving() && s.parts[0].position == s.parts[1].position {
					s.parts = append(s.parts[:0], s.parts[0+1:]...)
				}

				// smooth-move the tail
				s.parts[0].MoveTo(s.parts[1].position)

			} else {
				// stop eating so that in the next iteration the tail will be removed
				s.parts[0].isEating = false
			}

			// translate new tail to tail
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

	// paint tail again, because it has to be on the top (for smooth-move)
	s.parts[0].Render(screen)
	return nil
}

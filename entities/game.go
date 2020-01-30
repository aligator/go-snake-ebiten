package entities

import (
	"github.com/hajimehoshi/ebiten"
)

type Game struct {
	snake *Snake
	cooky *Cooky

	running bool
	points  int
}

func NewGame() Game {
	g := Game{
		running: true,
		points:  0,
	}

	g.snake = NewSnake(&g)
	g.cooky = NewCooky(&g)
	return g
}

func (g *Game) End() {
	g.running = false
}

func (g Game) IsRunning() bool {
	return g.running
}

func (g *Game) Update() error {
	if g.IsRunning() {
		if err := executeUpdates(g.snake.Update, g.cooky.Update); err != nil {
			return err
		}
	}
	return nil
}

func (g *Game) Render(screen *ebiten.Image) error {
	if err := executeRenderers(screen, g.snake.Render, g.cooky.Render); err != nil {
		return err
	}
	return nil
}

func executeUpdates(fns ...func() error) error {
	for _, fn := range fns {
		if err := fn(); err != nil {
			return err
		}
	}

	return nil
}

func executeRenderers(screen *ebiten.Image, fns ...func(screen *ebiten.Image) error) error {
	for _, fn := range fns {
		if err := fn(screen); err != nil {
			return err
		}
	}

	return nil
}

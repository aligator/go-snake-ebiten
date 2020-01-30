package main

import (
	"github.com/hajimehoshi/ebiten"
	"log"
	"snake-ebiten/entities"
	"snake-ebiten/util"
)

var game entities.Game

func init() {
	game = entities.NewGame()
}

func update(screen *ebiten.Image) error {
	if err := game.Update(); err != nil {
		return err
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	return game.Render(screen)
}

func main() {
	if err := ebiten.Run(update, util.Width, util.Height, 1, "Snake"); err != nil {
		log.Fatal(err)
	}
}

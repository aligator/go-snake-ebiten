package main

import (
	"github.com/hajimehoshi/ebiten"
	"log"
	"math/rand"
	"snake-ebiten/entities"
	"snake-ebiten/util"
	"time"
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
	rand.Seed(time.Now().UnixNano())

	if err := ebiten.Run(update, util.Width, util.Height, 1, "Snake"); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"github.com/hajimehoshi/ebiten"
	"log"
	"snake-ebiten/entities"
	"snake-ebiten/util"
)

var objects []entities.Object

func init() {
	objects = append(objects, entities.NewSnake())
}

func update(screen *ebiten.Image) error {
	for _, o := range objects {
		err := o.Update()
		if err != nil {
			return err
		}
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	for _, o := range objects {
		err := o.Render(screen)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	if err := ebiten.Run(update, util.Width, util.Height, 1, "Hello, World!"); err != nil {
		log.Fatal(err)
	}
}

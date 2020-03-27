package entities

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font/basicfont"
	"image/color"
	"strconv"
)

func init() {
	fmt.Println("Init Hud")
}

type Hud struct {
	game *Game
}

func NewHud(g *Game) *Hud {
	h := Hud{g}

	return &h
}

func (h *Hud) Update() error {
	return nil
}

func textDimension(text string) (w int, h int) {
	return 7 * len(text), 13
}

func (h *Hud) Render(screen *ebiten.Image) error {
	text.Draw(screen, "Score: "+strconv.Itoa(h.game.points), basicfont.Face7x13, 20, 20, color.White)

	if !h.game.running {
		gameOverText := "GAME OVER"
		textW, textH := textDimension(gameOverText)
		screenW := screen.Bounds().Dx()
		screenH := screen.Bounds().Dy()

		text.Draw(screen, gameOverText, basicfont.Face7x13, screenW/2-textW/2, screenH/2+textH/2, color.White)
	}

	return nil
}

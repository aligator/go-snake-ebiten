package entities

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font/basicfont"
	"image/color"
	"math/rand"
	"strconv"
	"time"
)

func init() {
	fmt.Println("Init Hud")

	rand.Seed(time.Now().UnixNano())
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

func (h *Hud) Render(screen *ebiten.Image) error {
	text.Draw(screen, "Score: "+strconv.Itoa(h.game.points), basicfont.Face7x13, 20, 20, color.White)
	return nil
}

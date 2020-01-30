package entities

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"math/rand"
	"snake-ebiten/util"
	"time"
)

var imgCooky *ebiten.Image

func init() {
	fmt.Println("Init Cooky")

	imgCooky = mustLoadTexture("assets/textures/cooky.png")

	rand.Seed(time.Now().UnixNano())
}

type Cooky struct {
	position Point
}

func NewCooky(g *Game) *Cooky {
	c := Cooky{}
	c.Respawn()

	return &c
}

func (c *Cooky) Update() error {
	return nil
}

func (c *Cooky) Render(screen *ebiten.Image) error {
	opt := ebiten.DrawImageOptions{}
	opt.GeoM.Translate(c.position.Xf(), c.position.Yf())

	screen.DrawImage(imgCooky, &opt)

	return nil
}

func (c *Cooky) Respawn() {
	maxX := util.GridWidth
	maxY := util.GridHeight

	x := rand.Intn(maxX)
	y := rand.Intn(maxY)

	c.position = Point{
		X: x,
		Y: y,
	}
}

package entity

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const moveSpeed = 5

type Paddle struct {
	X      int
	Y      int
	width  int
	height int
}

func CreatePaddle() *Paddle {
	p := &Paddle{
		X:      0,
		Y:      0,
		width:  20,
		height: 100,
	}

	return p
}

func (p *Paddle) moveDown() {
	_, h := ebiten.WindowSize()

	if p.Y+p.height < h {
		p.Y += moveSpeed
	}
}

func (p *Paddle) moveUp() {
	if p.Y > 0 {
		p.Y -= moveSpeed
	}
}

func (p *Paddle) Update(game *Game) {
	for _, key := range game.keys {
		switch key {
		case ebiten.KeyArrowDown:
			p.moveDown()
		case ebiten.KeyArrowUp:
			p.moveUp()
		}
	}
}

func (p *Paddle) Draw(screen *ebiten.Image) {
	geom := ebiten.GeoM{}
	geom.Translate(float64(p.X), float64(p.Y))

	vector.DrawFilledRect(screen, float32(p.X), float32(p.Y), float32(p.width), float32(p.height), color.White, true)
}

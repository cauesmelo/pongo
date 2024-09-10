package entity

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type ball struct {
	X      float64
	Y      float64
	size   int
	sprite *ebiten.Image
}

func CreateBall() *ball {
	img, _, err := ebitenutil.NewImageFromFile("./asset/pongo.png")
	if err != nil {
		log.Fatal(err)
	}

	w, h := ebiten.WindowSize()
	size := img.Bounds().Size()

	b := &ball{
		X:      float64((w / 2) - 20),
		Y:      float64((h / 2) - 20),
		sprite: img,
		size:   size.X,
	}

	return b
}

func (b *ball) Update() {
	// pass
}

func (b *ball) Draw(screen *ebiten.Image) {
	geom := ebiten.GeoM{}
	geom.Translate(b.X, b.Y)

	screen.DrawImage(b.sprite, &ebiten.DrawImageOptions{
		GeoM: geom,
	})
}

package entity

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type ball struct {
	X      int
	Y      int
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
		X:      w/2 - 20,
		Y:      h/2 - 20,
		sprite: img,
		size:   size.X,
	}

	return b
}

func (b *ball) Update(keys []ebiten.Key, state GameState) {
	// pass
}

func (b *ball) Draw(screen *ebiten.Image) {
	geom := ebiten.GeoM{}
	geom.Translate(float64(b.X), float64(b.Y))

	screen.DrawImage(b.sprite, &ebiten.DrawImageOptions{
		GeoM: geom,
	})
}

package entity

import (
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type ball struct {
	X         int
	Y         int
	size      int
	sprite    *ebiten.Image
	velocityX int
	velocityY int
}

const ballSpeed = 3

func CreateBall() *ball {
	img, _, err := ebitenutil.NewImageFromFile("./asset/pongo.png")
	if err != nil {
		log.Fatal(err)
	}

	w, h := ebiten.WindowSize()
	size := img.Bounds().Size()

	b := &ball{
		X:         w/2 - 20,
		Y:         h/2 - 20,
		sprite:    img,
		size:      size.X,
		velocityX: 0,
		velocityY: 0,
	}

	return b
}

func (b *ball) isStoped() bool {
	return b.velocityX == 0 && b.velocityY == 0
}

func getRng() int {
	if rand.Intn(2) == 0 {
		return ballSpeed
	} else {
		return ballSpeed * -1
	}
}

func (b *ball) launch() {
	b.velocityY = getRng()
	b.velocityX = getRng()
}

func (b *ball) checkBoundaries(newGame func()) {
	w, h := ebiten.WindowSize()

	if b.Y <= 0 || b.Y >= h-b.size {
		b.velocityY *= -1
	}

	if b.X <= 0 || b.X >= w-b.size {
		newGame()
	}
}

func (b *ball) checkPaddleCollision(entities []Entity) {
	for _, e := range entities {
		paddle, isPaddle := e.(*Paddle)

		if isPaddle {
			if b.X <= paddle.X+paddle.width && b.Y >= paddle.Y && b.Y <= paddle.Y+paddle.height {
				b.velocityX *= -1
			}
		}
	}
}

func (b *ball) Update(game *Game) {
	if game.state.Started && b.isStoped() {
		b.launch()
		return
	}

	b.checkBoundaries(game.NewGame)
	b.checkPaddleCollision(game.entities)

	b.X += b.velocityX
	b.Y += b.velocityY

}

func (b *ball) Draw(screen *ebiten.Image) {
	geom := ebiten.GeoM{}
	geom.Translate(float64(b.X), float64(b.Y))

	screen.DrawImage(b.sprite, &ebiten.DrawImageOptions{
		GeoM: geom,
	})
}

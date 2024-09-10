package main

import (
	"log"

	"github.com/cauesmelo/pongo/entity"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	entities []entity.Entity
	keys     []ebiten.Key
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])

	for _, e := range g.entities {
		e.Update(g.keys)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, e := range g.entities {
		e.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func startup() *Game {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Pongo")

	entities := []entity.Entity{
		entity.CreateBall(),
		entity.CreatePaddle(),
	}

	return &Game{
		entities,
		[]ebiten.Key{},
	}
}

func main() {
	game := startup()

	err := ebiten.RunGame(game)

	if err != nil {
		log.Fatal(err)
	}
}

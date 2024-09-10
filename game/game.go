package game

import (
	"log"

	"github.com/cauesmelo/pongo/entity"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	entities []entity.Entity
}

func (g *Game) Update() error {
	for _, e := range g.entities {
		e.Update()
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

	entities := []entity.Entity{}
	ball := entity.CreateBall()

	entities = append(entities, ball)

	return &Game{
		entities,
	}
}

func Run() {
	game := startup()

	err := ebiten.RunGame(game)

	if err != nil {
		log.Fatal(err)
	}
}

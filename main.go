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
	state    entity.GameState
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])

	if !g.state.Started && ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.state.Started = true
	}

	for _, e := range g.entities {
		e.Update(g.keys, g.state)
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
		entity.CreateText("Press SPACE to start"),
	}

	initialState := entity.GameState{
		Started: false,
	}

	return &Game{
		entities: entities,
		keys:     []ebiten.Key{},
		state:    initialState,
	}
}

func main() {
	game := startup()

	err := ebiten.RunGame(game)

	if err != nil {
		log.Fatal(err)
	}
}

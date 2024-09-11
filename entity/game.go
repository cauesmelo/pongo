package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type GameState struct {
	Started bool
}

type Game struct {
	entities []Entity
	keys     []ebiten.Key
	state    GameState
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])

	if !g.state.Started && ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.state.Started = true
	}

	for _, e := range g.entities {
		e.Update(g)
	}

	return nil
}

func (g *Game) GameOver() {
	g.NewGame()
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, e := range g.entities {
		e.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func (g *Game) NewGame() {
	initialState := GameState{
		Started: false,
	}

	entities := []Entity{
		CreateBall(),
		CreatePaddle(PaddlePositionLeft),
		CreatePaddle(PaddlePositionRight),
		CreateText("Press SPACE to start"),
	}

	g.entities = entities
	g.state = initialState
}

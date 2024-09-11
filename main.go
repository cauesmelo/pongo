package main

import (
	"log"

	"github.com/cauesmelo/pongo/entity"
	"github.com/hajimehoshi/ebiten/v2"
)

func startup() *entity.Game {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Pongo")

	game := &entity.Game{}
	game.NewGame()

	return game
}

func main() {
	game := startup()

	err := ebiten.RunGame(game)

	if err != nil {
		log.Fatal(err)
	}
}

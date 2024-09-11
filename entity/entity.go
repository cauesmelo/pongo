package entity

import "github.com/hajimehoshi/ebiten/v2"

type Entity interface {
	Update(game *Game)
	Draw(screen *ebiten.Image)
}

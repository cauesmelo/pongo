package entity

import "github.com/hajimehoshi/ebiten/v2"

type Entity interface {
	Update(keys []ebiten.Key, state GameState)
	Draw(screen *ebiten.Image)
}

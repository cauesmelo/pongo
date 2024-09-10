package entity

import "github.com/hajimehoshi/ebiten/v2"

type Entity interface {
	Update(keys []ebiten.Key)
	Draw(screen *ebiten.Image)
}

package entity

import (
	"bytes"
	_ "embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

//go:embed FiraSans-Regular.ttf
var firaSans []byte

type info struct {
	text string
	font *text.GoTextFaceSource
}

func CreateText(txt string) *info {
	r := bytes.NewReader(firaSans)

	s, err := text.NewGoTextFaceSource(r)
	if err != nil {
		log.Fatal(err)
	}

	return &info{
		text: txt,
		font: s,
	}
}

func (i *info) Update(game *Game) {
	if game.state.Started && i.text != "" {
		i.text = ""
	}
}

func (i *info) Draw(screen *ebiten.Image) {
	if i.text != "" {
		f := &text.GoTextFace{
			Source: i.font,
			Size:   24,
		}

		op := &text.DrawOptions{}
		op.GeoM.Translate(float64(w/2), float64(h/2+100))
		op.PrimaryAlign = text.AlignCenter

		text.Draw(screen, i.text, f, op)
	}
}

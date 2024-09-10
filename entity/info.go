package entity

import (
	"bytes"
	_ "embed"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type info struct {
	text string
	font *text.GoTextFaceSource
}

func CreateText(txt string) *info {
	f, err := os.ReadFile("./asset/FiraSans-Regular.ttf")
	if err != nil {
		log.Fatal(err)
	}

	r := bytes.NewReader(f)

	s, err := text.NewGoTextFaceSource(r)
	if err != nil {
		log.Fatal(err)
	}

	return &info{
		text: txt,
		font: s,
	}
}

func (i *info) Update(keys []ebiten.Key, state GameState) {
	if state.Started && i.text != "" {
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
		w, h := ebiten.WindowSize()
		op.GeoM.Translate(float64(w/2), float64(h/2+100))
		op.PrimaryAlign = text.AlignCenter

		text.Draw(screen, i.text, f, op)
	}
}

package component

import (
	"fmt"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/basicfont"
)

type OptionsViewer struct {
	Options []string
	CanBack bool
}

func (*OptionsViewer) Update() error {
	return nil
}

func (o *OptionsViewer) StringOutput(screen *ebiten.Image) string {
	optionsText := strings.Builder{}
	for i, option := range o.Options {
		// do not put a \n for the last option
		optionsText.WriteString(fmt.Sprintf("%d) %s", i+1, option))
		if i < len(o.Options)-1 {
			optionsText.WriteString("\n")
		}

	}
	if o.CanBack {
		optionsText.WriteString("\n0) Back")
	}

	return optionsText.String()
}

func (o *OptionsViewer) Draw(screen *ebiten.Image) {
	bottomOffset := len(o.Options)*20 + 50
	op := &text.DrawOptions{}
	op.LineSpacing = 20
	op.GeoM.Translate(float64(screen.Bounds().Min.X)+10, float64(screen.Bounds().Max.Y-bottomOffset))
	text.Draw(screen, o.StringOutput(screen), text.NewGoXFace(basicfont.Face7x13), op)
}

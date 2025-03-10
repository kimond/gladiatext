package component

import (
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/basicfont"
)

type StringOutputter interface {
	StringOutput(screen *ebiten.Image) string
}

func JoinOutputter(screen *ebiten.Image, outputters []StringOutputter, sep string) string {
	outputText := strings.Builder{}
	for i, outputter := range outputters {
		outputText.WriteString(outputter.StringOutput(screen))
		if i < len(outputters)-1 {
			outputText.WriteString(sep)
		}
	}

	return outputText.String()
}

func DrawFromBottom(screen *ebiten.Image, output string, paddingX, paddingY int) {
	offset := ((len(strings.Split(output, "\n")) - 1) * 20) + paddingY

	op := &text.DrawOptions{}
	op.LineSpacing = 20
	op.GeoM.Translate(float64(screen.Bounds().Min.X+paddingX), float64(screen.Bounds().Max.Y-offset))
	text.Draw(screen, output, text.NewGoXFace(basicfont.Face7x13), op)
}

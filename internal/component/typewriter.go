package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/basicfont"
)

type Typewriter struct {
	runes   []rune
	Text    string
	counter int
}

func NewTypewriter() *Typewriter {
	return &Typewriter{}
}

func (t *Typewriter) Update() error {
	t.runes = ebiten.AppendInputChars(t.runes[:0])
	t.Text += string(t.runes)

	if repeatingKeyPressed(ebiten.KeyBackspace) {
		if len(t.Text) > 0 {
			t.Text = t.Text[:len(t.Text)-1]
		}
	}

	t.counter++
	return nil
}

func (t *Typewriter) StringOutput(screen *ebiten.Image) string {
	inputText := t.Text
	if t.counter%60 < 30 {
		inputText += "_"
	}
	inputText = "> " + inputText

	return inputText
}

func (t *Typewriter) Draw(screen *ebiten.Image) {
	drawOptions := &text.DrawOptions{}
	drawOptions.GeoM.Translate(float64(screen.Bounds().Min.X)+10, float64(screen.Bounds().Max.Y-30))
	text.Draw(screen, t.StringOutput(screen), text.NewGoXFace(basicfont.Face7x13), drawOptions)
}

func (t *Typewriter) Clear() {
	t.Text = ""
}

// repeatingKeyPressed return true when key is pressed considering the repeat state.
func repeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 30
		interval = 3
	)
	d := inpututil.KeyPressDuration(key)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)%interval == 0 {
		return true
	}
	return false
}

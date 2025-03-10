package internal

import (
	"kimond/gladiatext/internal/actor"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/basicfont"
)

type CombatResultScene struct {
	Result          string
	displayedTime   time.Time
	continueEnabled bool
}

func NewCombatResultScene(result string) *CombatResultScene {
	return &CombatResultScene{
		Result:          result,
		displayedTime:   time.Now(),
		continueEnabled: false,
	}
}

func (c *CombatResultScene) Update(state *GameState) error {
	if c.continueEnabled {
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			player := actor.NewCharacter("Hero", 10, 8, 12, 10)
			state.SceneManager.GoTo(NewCombatScene(player, actor.NewRandomNPC()))
		}
	}

	if time.Since(c.displayedTime) > 3*time.Second {
		c.continueEnabled = true
	}

	return nil
}

func (c *CombatResultScene) Draw(screen *ebiten.Image) {
	screenWidth := screen.Bounds().Dx()
	screenHeight := screen.Bounds().Dy()
	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(screenWidth/2-len(c.Result)), float64(screenHeight/2+20))
	text.Draw(screen, c.Result, text.NewGoXFace(basicfont.Face7x13), op)

	if c.continueEnabled {
		op.GeoM.Translate(0, 20)
		text.Draw(screen, "Press Enter to continue", text.NewGoXFace(basicfont.Face7x13), op)
	}
}

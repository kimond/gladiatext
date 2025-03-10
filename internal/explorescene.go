package internal

import (
	"image/color"
	"kimond/gladiatext/internal/actor"
	"kimond/gladiatext/internal/component"

	"github.com/hajimehoshi/ebiten/v2"
)

type ExploreScene struct {
	LogViewer  *component.LogViewer
	exloreMenu *component.Menu
	nextAction string
}

func NewExploreScene() *ExploreScene {
	es := &ExploreScene{
		LogViewer: &component.LogViewer{Title: "", Log: []string{"What to do next?"}},
	}

	es.exloreMenu = component.NewMenu(
		[]string{"Explore", "Fight", "Rest", "Options"},
		component.MenuActions{
			1: func() { es.LogViewer.Log = append(es.LogViewer.Log, "Exploring...") },
			2: func() {
				es.LogViewer.Log = append(es.LogViewer.Log, "Fighting...")
				es.nextAction = "fight"
			},
			3: func() {
				es.nextAction = "rest"
			},
			4: func() { es.LogViewer.Log = append(es.LogViewer.Log, "Options...") },
		})

	return es
}

func (e *ExploreScene) Update(state *GameState) error {
	e.LogViewer.Update()
	e.exloreMenu.Update()

	if e.nextAction == "fight" {
		e.nextAction = ""
		state.SceneManager.GoTo(NewCombatScene(state.World.Player.Character, actor.NewRandomNPC()))
	}

	if e.nextAction == "rest" {
		e.nextAction = ""
		e.LogViewer.Log = append(e.LogViewer.Log, "Resting...")
		e.LogViewer.Log = append(e.LogViewer.Log, "You feel rested. HP restored.")
		state.World.Player.Character.HP = state.World.Player.Character.MaxHP
	}

	return nil
}

func (e *ExploreScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	outputText := component.JoinOutputter(screen, []component.StringOutputter{e.LogViewer, e.exloreMenu}, "\n")
	component.DrawFromBottom(screen, outputText, 10, 30)
}

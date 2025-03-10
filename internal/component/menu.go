package component

import (
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
)

type MenuActions = map[int]func()

type Menu struct {
	Optionsviewer *OptionsViewer
	Typewriter    *Typewriter
	Actions       MenuActions
	backAction    func()
}

func NewMenu(options []string, actions MenuActions) *Menu {
	return &Menu{
		Optionsviewer: &OptionsViewer{Options: options},
		Typewriter:    NewTypewriter(),
		Actions:       actions,
	}
}

func NewMenuWithBack(options []string, actions MenuActions, backAction func()) *Menu {
	return &Menu{
		Optionsviewer: &OptionsViewer{Options: options, CanBack: true},
		Typewriter:    NewTypewriter(),
		Actions:       actions,
		backAction:    backAction,
	}
}

func (m *Menu) Update() error {
	m.Optionsviewer.Update()
	m.Typewriter.Update()

	if ebiten.IsKeyPressed(ebiten.KeyEnter) && m.Typewriter.Text != "" {
		choice, err := strconv.Atoi(m.Typewriter.Text)
		if err != nil {
			return err
		}

		if choice == 0 && m.Optionsviewer.CanBack {
			m.backAction()
		}

		action, ok := m.Actions[choice]
		if ok {
			action()
		}
		m.Typewriter.Clear()
	}

	return nil
}

func (m *Menu) StringOutput(screen *ebiten.Image) string {
	return JoinOutputter(screen, []StringOutputter{m.Optionsviewer, m.Typewriter}, "\n")
}

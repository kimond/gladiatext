package internal

import (
	"kimond/gladiatext/internal/world"

	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Update(state *GameState) error
	Draw(screen *ebiten.Image)
}

type SceneManager struct {
	current Scene
	next    Scene
}

func (s *SceneManager) Update() error {
	if s.next != nil {
		s.current = s.next
		s.next = nil
	} else {
		world := world.NewWorld()
		return s.current.Update(&GameState{SceneManager: s, World: world})
	}

	return nil
}

func (s *SceneManager) Draw(screen *ebiten.Image) {
	s.current.Draw(screen)
}

func (s *SceneManager) GoTo(next Scene) {
	if s.current == nil {
		s.current = next
	} else {
		s.next = next
	}
}

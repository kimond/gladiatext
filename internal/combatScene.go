package internal

import (
	"fmt"
	"image/color"
	"kimond/gladiatext/internal/actor"
	"kimond/gladiatext/internal/animation"
	"kimond/gladiatext/internal/combat"
	"kimond/gladiatext/internal/component"
	"kimond/gladiatext/internal/controller"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/basicfont"
)

type CombatScene struct {
	Player           *actor.Character
	Enemy            *actor.NPC
	CombatController *controller.CombatController
	Turn             int
	EnemyLog         *component.LogViewer
	LogViewer        *component.LogViewer
	AttackMenuOpen   bool
	combatMenu       *component.Menu
	attackMenu       *component.Menu
	asciiAnimation   *component.ASCIIAnimation
	lastHintTime     time.Time
}

func NewCombatScene(player *actor.Character, enemy *actor.NPC) *CombatScene {
	cs := &CombatScene{
		Player:           player,
		Enemy:            enemy,
		CombatController: controller.NewCombatController(player, enemy),
		Turn:             0,
		EnemyLog:         &component.LogViewer{Title: enemy.Name},
		LogViewer:        &component.LogViewer{Title: "Combat Log", Log: []string{"Begin Combat"}},
		asciiAnimation:   component.NewASCIIAnimation(),
		lastHintTime:     time.Now(),
	}

	cs.asciiAnimation.SetFramesAndSpeed(animation.CombatIdleFrames, 500*time.Millisecond)

	combatMenu := component.NewMenu(
		[]string{"Attack", "Dodge", "Parry", "Breath"},
		component.MenuActions{
			1: func() { cs.AttackMenuOpen = true },
			2: func() { cs.ProcessTurn(combat.ActionDodge) },
			3: func() { cs.ProcessTurn(combat.ActionParry) },
			4: func() { cs.ProcessTurn(combat.ActionBreath) },
		},
	)
	cs.combatMenu = combatMenu

	attackMenu := component.NewMenuWithBack([]string{"Light"},
		component.MenuActions{1: func() { cs.ProcessTurn(combat.ActionAttack) }},
		func() { cs.AttackMenuOpen = false },
	)
	cs.attackMenu = attackMenu
	cs.GiveHint()

	return cs
}

func (s *CombatScene) ProcessTurn(action combat.Action) {
	roundLogs := s.CombatController.ResolveRound(action)

	s.LogViewer.SetLogs(roundLogs)
	s.AttackMenuOpen = false
}

func (s *CombatScene) CheckIfBattleIsOver(sceneManager *SceneManager) {
	if s.Enemy.Character.HP <= 0 && s.Player.HP <= 0 {
		sceneManager.GoTo(NewCombatResultScene("Draw"))
	}
	if s.Player.HP <= 0 {
		sceneManager.GoTo(NewCombatResultScene("Defeat"))
	}
	if s.Enemy.Character.HP <= 0 {
		sceneManager.GoTo(NewCombatResultScene("Victory"))
	}
}

func (s *CombatScene) GiveHint() {
	hint := s.Enemy.GetHint()
	s.EnemyLog.SetLog(hint)
}

func (s *CombatScene) SetEnemyLog(text string) {
	s.EnemyLog.SetLog(text)
}

func (s *CombatScene) Update(state *GameState) error {
	s.asciiAnimation.Update()
	s.EnemyLog.Update()
	s.LogViewer.Update()
	if s.AttackMenuOpen {
		s.attackMenu.Update()
	} else {
		s.combatMenu.Update()
	}

	if time.Since(s.lastHintTime) > 10*time.Second {
		s.GiveHint()
		s.lastHintTime = time.Now()
	}

	s.CheckIfBattleIsOver(state.SceneManager)

	return nil
}

func (s *CombatScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	eOp := &text.DrawOptions{}
	eOp.GeoM.Translate(10, 100)
	eOp.LineSpacing = 20
	text.Draw(screen, s.EnemyLog.StringOutput(screen), text.NewGoXFace(basicfont.Face7x13), eOp)

	var activeMenu *component.Menu
	if s.AttackMenuOpen {
		activeMenu = s.attackMenu
	} else {
		activeMenu = s.combatMenu
	}
	outputText := component.JoinOutputter(screen, []component.StringOutputter{s.LogViewer, activeMenu}, "\n")
	component.DrawFromBottom(screen, outputText, 10, 30)

	// Render HP & Stamina
	hpOp := &text.DrawOptions{}
	hpOp.GeoM.Translate(600, 40)
	text.Draw(screen, fmt.Sprintf("HP: %d/%d", s.Player.HP, s.Player.MaxHP), text.NewGoXFace(basicfont.Face7x13), hpOp)
	staminaOp := &text.DrawOptions{}
	staminaOp.GeoM.Translate(600, 60)
	text.Draw(screen, fmt.Sprintf("Stamina: %d/%d", s.Player.Stamina, s.Player.MaxStamina), text.NewGoXFace(basicfont.Face7x13), staminaOp)
}

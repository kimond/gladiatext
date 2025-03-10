package actor

import (
	"kimond/gladiatext/internal/combat"
	"math/rand"
)

// Personality types define combat tendencies
type Personality string

const (
	Aggressive Personality = "Aggressive"
	Defensive  Personality = "Defensive"
	Cautious   Personality = "Cautious"
	Reckless   Personality = "Reckless"
)

// Combat states provide hints to the player
type CombatMood string

const (
	ConfidentMood CombatMood = "Confident"
	HesitantMood  CombatMood = "Hesitant"
	ExhaustedMood CombatMood = "Exhausted"
	EnragedMood   CombatMood = "Enraged"
	NeutralMood   CombatMood = "Neutral"
)

type NPC struct {
	Name        string
	Character   *Character
	Personality Personality
	Mood        CombatMood
}

func NewRandomNPC() *NPC {
	personalities := []Personality{Aggressive, Defensive, Cautious, Reckless}
	name := GenerateRandomName()

	return &NPC{
		Name:        name,
		Character:   NewCharacter(name, 10, 10, 10, 10),
		Personality: personalities[rand.Intn(len(personalities))],
		Mood:        NeutralMood,
	}
}

// Update NPC combat state based on stamina and health
func (n *NPC) UpdateMood() {
	if n.Character.Stamina < n.Character.MaxStamina/3 {
		n.Mood = ExhaustedMood
	} else if n.Character.HP < n.Character.MaxHP/3 {
		n.Mood = HesitantMood
	} else if n.Personality == Aggressive && n.Character.HP > n.Character.MaxHP/2 {
		n.Mood = EnragedMood
	} else {
		n.Mood = NeutralMood
	}
}

func (n *NPC) GetHint() string {
	return getRandomHintForMood(n.Mood)
}

func getRandomHintForMood(mood CombatMood) string {
	hints := hintsPerMood[mood]
	index := rand.Intn(len(hints))

	return hints[index]
}

var hintsPerMood = map[CombatMood][]string{
	ConfidentMood: {
		"stands tall, unfazed.",
		"smirk as they size you up.",
		"You sense no hesitation in their movements.",
	},
	HesitantMood: {
		"their's grip falters slightly.",
		"You notice a flicker of doubt in their eyes.",
		"Their stance shifts uneasily.",
	},
	EnragedMood: {
		"tightens their grip, ready to strike.",
		"A fire burns in their eyes as they lunge forward.",
		"They breathe heavily, pushing forward with force.",
	},
	ExhaustedMood: {
		"Their movements slow down, exhaustion setting in.",
		"You hear them panting, struggling to keep up.",
		"their's arms tremble slightly.",
	},
	NeutralMood: {
		"watching you closely.",
		"seems to be waiting for your next move.",
	},
}

func (n *NPC) ChooseAction() combat.Action {
	switch n.Personality {
	case Aggressive:
		return weightedChoice(map[combat.Action]int{combat.ActionAttack: 70, combat.ActionParry: 20, combat.ActionDodge: 10})
	case Defensive:
		return weightedChoice(map[combat.Action]int{combat.ActionAttack: 20, combat.ActionParry: 50, combat.ActionDodge: 30})
	case Cautious:
		return weightedChoice(map[combat.Action]int{combat.ActionAttack: 10, combat.ActionParry: 30, combat.ActionDodge: 60})
	case Reckless:
		return weightedChoice(map[combat.Action]int{combat.ActionAttack: 80, combat.ActionParry: 10, combat.ActionDodge: 10})
	default:
		return weightedChoice(map[combat.Action]int{combat.ActionAttack: 33, combat.ActionParry: 33, combat.ActionDodge: 33})
	}
}

func weightedChoice(options map[combat.Action]int) combat.Action {
	total := 0
	for _, weight := range options {
		total += weight
	}

	randVal := rand.Intn(total)
	for choice, weight := range options {
		if randVal < weight {
			return choice
		}
		randVal -= weight
	}
	return combat.ActionBreath
}

package actor

import (
	"kimond/gladiatext/internal/stats"
)

// Character struct holding player/enemy data
type Character struct {
	Name       string
	HP         int
	MaxHP      int
	Stamina    int
	MaxStamina int
	Momentum   float64
	Stats      stats.Stats
}

// NewCharacter creates a new character with calculated HP & Stamina
func NewCharacter(name string, str, dex, end, mot int) *Character {
	maxHP := 25 + (end * 2)      // Endurance increases HP slightly
	maxStamina := 25 + (end * 2) // Endurance heavily influences stamina

	return &Character{
		Name:       name,
		HP:         maxHP,
		MaxHP:      maxHP,
		Stamina:    maxStamina,
		MaxStamina: maxStamina,
		Momentum:   0,
		Stats:      stats.Stats{Strength: str, Dexterity: dex, Endurance: end, Motivation: mot},
	}
}

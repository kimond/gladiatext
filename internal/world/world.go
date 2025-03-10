package world

import "kimond/gladiatext/internal/actor"

type World struct {
	Player *actor.Player
}

func NewWorld() *World {
	player := actor.NewPlayer("Hero")
	return &World{
		Player: player,
	}
}

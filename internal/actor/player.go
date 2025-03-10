package actor

type Player struct {
	Name      string
	Character *Character
	Fame      int
}

func NewPlayer(name string) *Player {
	return &Player{
		Name:      name,
		Character: NewCharacter(name, 20, 20, 20, 5),
		Fame:      0,
	}
}

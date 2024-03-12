package player

func NewPlayer(name string) *Player {
	return &Player{
		Name:    name,
		Kills:   0,
		Victims: map[string]int{},
	}
}

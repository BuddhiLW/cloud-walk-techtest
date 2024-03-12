package player

func NewPlayer(name string) *Player {
	return &Player{
		Name:    name,
		Kills:   0,
		Victims: map[string]int{},
	}
}

func (p *Player) AddKill() {
	p.Kills++
}

func (p *Player) AddVictim(victim string) {
	// p.Victims = append(p.Victims, victim)
	// p.AddKill()
}

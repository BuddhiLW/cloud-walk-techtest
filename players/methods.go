package player

func NewPlayer(name string) *Player {
	return &Player{
		Name:    name,
		Kills:   0,
		Victims: map[string]Way{},
	}
}

func (p *Player) hasVictimized(victim string) bool {
	_, ok := p.Victims[victim]
	return ok
}

func (p *Player) AddKill(victim, action string) {
	p.Kills++
	if p.hasVictimized(victim) {
		p.Victims[victim][action]++
	} else {
		p.Victims[victim] = Way{action: 1}
	}
}

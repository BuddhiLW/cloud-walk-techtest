package player

func NewPlayer(name string) *Player {
	return &Player{
		Name:    name,
		Kills:   0,
		Victims: map[string]Ways{},
	}
}
func World() *Player {
	return NewPlayer("<world>")
}

func NewPlayers() *Players {
	world := World()
	return &Players{world.Name: world}
}

func (p *Players) AddPlayer(player *Player) {
	(*p)[player.Name] = player
}

func (p *Player) hasVictimized(victim string) bool {
	_, ok := p.Victims[victim]
	return ok
}

func (p *Player) AddKill(victim, action string) {
	// p.Kills++
	if p.hasVictimized(victim) {
		p.Victims[victim][action]++
	} else {
		p.Victims[victim] = Ways{action: 1}
	}
}

func (w Ways) CountKills() int {
	kills := 0
	for _, v := range w {
		kills += v
	}
	return kills
}

func (v *Victims) CountKills() int {
	kills := 0
	for _, w := range *v {
		kills += w.CountKills()
	}
	return kills
}

func (ws *Ways) HasMode(mode string) bool {
	_, ok := (*ws)[mode]
	return ok
}

func NewWays() *Ways {
	return &Ways{}
}

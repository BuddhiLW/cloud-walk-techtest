package game

import (
	player "github.com/BuddhiLW/cloud-walk-techtest/player"
)

func (g *Game) addPlayer(p Player) {
	g.Players[p.Name] = p
}

func (g *Game) countKills(name string) int {
	// Get the kills from the player
	kills := g.Players[name].Kills

	// Remove the kills from the world
	killWorld := g.Players["<world>"].Victims[name]
	return kills - killWorld
}

func NewGame() *Game {
	var world = player.NewPlayer("<world>")
	var game = Game{Players: map[string]Player{}}
	game.addPlayer(*world)

	return &game
}

func (g Game) NewEncodeGame() *EncodeGame {
	totalKills := 0
	killMap := make(map[string]int)
	for _, p := range g.Players {
		totalKills += p.Kills
		killMap[p.Name] = p.Kills
	}

	return &EncodeGame{
		TotalKills: totalKills,
		Players:    make([]string, len(g.Players)),
		Kills:      killMap,
	}
}

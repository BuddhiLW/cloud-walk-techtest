package game

import (
	p "github.com/BuddhiLW/cloud-walk-techtest/players"
)

func NewGame(ps *p.Players) *Game {
	var game = Game{Players: *ps}
	return &game
}

// Boundary rules for the game
// 1. If the killer is "<world>", the kill is subtracted from the victim
// 2. All kills by "<world>" are added to the total kills, normally
// 3. "<world>" must not be listed as a player

func (g *Game) NewEncodeGame() *EncodeGame {
	totalKills := 0
	killMap := make(map[string]int)
	playerNames := make([]string, 0, len(g.Players)-1)

	for name, p := range g.Players {
		if name == "<world>" {
			// Conform to the counting rules 2 and 3.
			// -------------------------------------------
			totalKills += p.Kills
		} else {
			// Add player to the list of players
			playerNames = append(playerNames, name)

			// Do not consider <world> kills into the calculation
			// of the total kills
			totalKills += p.Kills

			// Conform to the counting rule 1.
			// -------------------------------------------
			// Consider <world> kills into the calculation
			// of the player's kills (score)
			killMap[name] = g.countKills(name)
		}
	}

	return &EncodeGame{
		TotalKills: totalKills,
		Players:    playerNames,
		Kills:      killMap,
	}
}

func (g *Game) AddPlayer(p *p.Player) {
	// Check if player already exists
	if _, ok := g.Players[p.Name]; ok {
		return
	}

	// Add player to the game -- case if it doesn't exist
	g.Players[p.Name] = p
}

func (g *Game) countKills(name string) int {
	// Get the kills from the player
	kills := g.Players[name].Kills

	// Remove the kills from the world
	killWorld := g.Players["<world>"].Victims[name].CountKills()

	return kills - killWorld
}

// Enconding logic

// game := Game{
//     // initialize the Game type..
// }

// encodeGame := EncodeGame{
//     TotalKills: game.TotalKills,
//     Players:    make([]string, len(game.Players)),
//     Kills:      game.Kills,
// }

// for i, player := range game.Players {
//     encodeGame.Players[i] = player.Name
// }

// jsonStr, err := json.Marshal(encodeGame)
// if err != nil {
//     // handle error
// }
// fmt.Println(string(jsonStr))

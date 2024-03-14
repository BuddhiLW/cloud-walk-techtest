package game

import (
	"sort"

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
	killMap := Kills{}
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

func (g *Game) countKills(name string) Kill {
	// Get the kills from the player
	kills := g.Players[name].Kills

	// Remove the kills from the world
	killWorld := g.Players["<world>"].Victims[name].CountKills()

	// Calculate the kills

	return Kill(kills - killWorld)
}

func NewEncodeGames() *EncodeGames {
	return &EncodeGames{}
}

func NewRank() *Rank {
	return &Rank{}
}

func NewGameRank() *GameRank {
	return &GameRank{}
}

func (g *EncodeGame) RankPlayersByKills() (*GameRank, []*Rank) {
	gameRank := NewGameRank()
	scoreList := make([]*Rank, 0, len(g.Kills))

	for playerName, kill := range g.Kills {
		rank := NewRank()
		*rank = Rank{Name: Name(playerName), Kills: int(kill)}
		scoreList = append(scoreList, rank)
	}

	sort.Slice(scoreList, func(i, j int) bool {
		return scoreList[i].Kills > scoreList[j].Kills // Descending order
	})

	for i, rank := range scoreList {
		rank.Position = i + 1
		(*gameRank)[rank.Name] = rank
	}

	return gameRank, scoreList
}

func (g *Game) CountTypeKills() p.Ways {
	ways := p.NewWays()

	// Count the kills for each kill-mode, in the game
	for _, p := range g.Players {
		for _, way := range p.Victims {
			for mode, kills := range way {
				if ways.HasMode(mode) {
					(*ways)[mode] += kills
				} else {
					(*ways)[mode] = kills
				}
			}
		}
	}
	return *ways
}

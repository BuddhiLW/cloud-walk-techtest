package game

import p "github.com/BuddhiLW/cloud-walk-techtest/players"

func NewGame() *Game {
	var world = p.NewPlayer("<world>")
	var game = Game{Players: map[string]p.Player{}}
	game.AddPlayer(*world)

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

func (g *Game) AddPlayer(p p.Player) {
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
	killWorld := g.Players["<world>"].Victims[name]
	return kills - killWorld
}

func (g *Game) AddTotalKills() {
}

func (g *Game) AddKill(killer string, killed string) {
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

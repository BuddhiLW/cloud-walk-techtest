package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// Data can be found in the following URL:
// gist: 	https://gist.github.com/cloudwalk-tests/be1b636e58abff14088c8b5309f575d8
// raw: 	https://gist.githubusercontent.com/cloudwalk-tests/be1b636e58abff14088c8b5309f575d8/raw/df6ef4a9c0b326ce3760233ef24ae8bfa8e33940/qgames.log

func ReadGist(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	return string(body)
}

func (g *Game) AddPlayer(p Player) {
	// Check if player already exists
	if _, ok := g.Players[p.Name]; ok {
		return
	}

	// Add player to the game -- case if it doesn't exist
	g.Players[p.Name] = p
}

func (p *Player) AddKill() {
	p.Kills++
}

func (p *Player) AddVictim(victim string) {
	p.Victims = append(p.Victims, victim)
	p.AddKill()
}

func (g *Game) AddKill(killer string, killed string) {

	// if killer == "<world>" {
	// 	g.Kills[killed]--
	// } else {
	// 	g.Kills[killer]++
	// }
}

// func (g *Game) AddKill(killer string, killed string) {
// 	// Check if killer is already in the game
// 	if _, ok := g.Players[killer]; !ok {
// 		g.AddPlayer(Player{Name: killer})
// 	}

// 	// if killer == "<world>" {
// 	// 	g.Kills[killed]--
// 	// } else {
// 	// 	g.Kills[killer]++
// 	// }
// }

func (g *Game) AddTotalKills() {
	g.TotalKills++
}

func main() {
	rawData := ReadGist("https://gist.githubusercontent.com/cloudwalk-tests/be1b636e58abff14088c8b5309f575d8/raw/df6ef4a9c0b326ce3760233ef24ae8bfa8e33940/qgames.log")
	dataBySplit := strings.Split(rawData, "------------------------------------------------------------")

	gameCount := 0
	for _, v := range dataBySplit {
		// has regex =InitGame:= or not -- determine if it's a game
		if strings.Contains(v, "InitGame:") {
			gameData := strings.Split(v, "\n")
			for n, v := range gameData {
				fmt.Println(n, string(v))
			}

			// Count the number of games
			gameCount++
		}
	}
	fmt.Println("Game Count: ", gameCount)
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

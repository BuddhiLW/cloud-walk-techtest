package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/BuddhiLW/cloud-walk-techtest/data"
	g "github.com/BuddhiLW/cloud-walk-techtest/game"
	p "github.com/BuddhiLW/cloud-walk-techtest/players"
)

// Data can be found in the following URL:
// gist: 	https://gist.github.com/cloudwalk-tests/be1b636e58abff14088c8b5309f575d8
// raw: 	https://gist.githubusercontent.com/cloudwalk-tests/be1b636e58abff14088c8b5309f575d8/raw/df6ef4a9c0b326ce3760233ef24ae8bfa8e33940/qgames.log

func main() {
	var gist data.Gist = "https://gist.githubusercontent.com/cloudwalk-tests/be1b636e58abff14088c8b5309f575d8/raw/df6ef4a9c0b326ce3760233ef24ae8bfa8e33940/qgames.log"
	rawData := gist.ReadGist()
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

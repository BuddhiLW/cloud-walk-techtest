package main

import (
	"encoding/json"
	"fmt"

	data "github.com/BuddhiLW/cloud-walk-techtest/data"
	g "github.com/BuddhiLW/cloud-walk-techtest/game"
	// p "github.com/BuddhiLW/cloud-walk-techtest/players"
)

// Data can be found in the following URL:
// gist: 	https://gist.github.com/cloudwalk-tests/be1b636e58abff14088c8b5309f575d8
// raw: 	https://gist.githubusercontent.com/cloudwalk-tests/be1b636e58abff14088c8b5309f575d8/raw/df6ef4a9c0b326ce3760233ef24ae8bfa8e33940/qgames.log

func main() {
	// Fetch data
	var gist data.Gist = "https://gist.githubusercontent.com/cloudwalk-tests/be1b636e58abff14088c8b5309f575d8/raw/df6ef4a9c0b326ce3760233ef24ae8bfa8e33940/qgames.log"
	rawData := gist.ReadGist()
	rawDataByGame := rawData.ToGames()

	// Process data, for each game
	for _, game := range rawDataByGame {
		lines := game.ToLines()

		// Parse and extract data
		var gl data.GameLines = lines
		players := gl.Players()
		gl.Kills(players)

		// Encode data in Game format
		game := g.NewGame(players)
		encodebleGame := game.NewEncodeGame()
		json, _ := json.Marshal(encodebleGame)

		// Print the result of the game
		fmt.Println(string(json))
	}
}

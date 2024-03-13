package game

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/BuddhiLW/cloud-walk-techtest/data"
)

func GameOutput(n ...int) string {
	var gist data.Gist = "https://gist.githubusercontent.com/cloudwalk-tests/be1b636e58abff14088c8b5309f575d8/raw/df6ef4a9c0b326ce3760233ef24ae8bfa8e33940/qgames.log"
	rawData := gist.ReadGist()
	rawDataByGame := rawData.ToGames()

	if len(n) == 0 {
		allEncodedGames := NewEncodeGames()
		log.Println("All games (N): ", len(rawDataByGame)+1)
		log.Println("---------------------------------\n")

		for i, game := range rawDataByGame {
			lines := game.ToLines()
			// Parse and extract data
			var gl data.GameLines = lines
			players := gl.Players()
			gl.Kills(players)

			// Encode data in Game format
			game := NewGame(players)
			encodeableGame := game.NewEncodeGame()

			// Add the game to the list of games (Encoded format)
			index := strconv.Itoa(i + 1)
			index = "game_" + index
			(*allEncodedGames)[index] = encodeableGame
		}

		json, _ := json.Marshal(allEncodedGames)
		printableJSON, _ := PrettyString(string(json))
		fmt.Println(printableJSON)
		return string(json)
	}

	// Check if the game exists in log
	if n[0]-1 > len(rawDataByGame) {
		log.Fatalf("Game %d not found; the log only contains up to game-number %d", n[0], len(rawDataByGame)+1)
	}

	// Process data, for game n
	lines := rawDataByGame[n[0]-1].ToLines()

	// Parse and extract data
	var gl data.GameLines = lines
	players := gl.Players()
	gl.Kills(players)

	// Encode data in Game format
	game := NewGame(players)
	encodebleGame := game.NewEncodeGame()
	json, _ := json.Marshal(encodebleGame)

	// Print the result of the game
	printableJSON, _ := PrettyString(string(json))
	fmt.Println(printableJSON)

	return string(json)
}

func PrettyString(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}

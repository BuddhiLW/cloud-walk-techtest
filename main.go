package main

import (
	"fmt"

	data "github.com/BuddhiLW/cloud-walk-techtest/data"
	// g "github.com/BuddhiLW/cloud-walk-techtest/game"
	// p "github.com/BuddhiLW/cloud-walk-techtest/players"
)

// Data can be found in the following URL:
// gist: 	https://gist.github.com/cloudwalk-tests/be1b636e58abff14088c8b5309f575d8
// raw: 	https://gist.githubusercontent.com/cloudwalk-tests/be1b636e58abff14088c8b5309f575d8/raw/df6ef4a9c0b326ce3760233ef24ae8bfa8e33940/qgames.log

func main() {
	var gist data.Gist = "https://gist.githubusercontent.com/cloudwalk-tests/be1b636e58abff14088c8b5309f575d8/raw/df6ef4a9c0b326ce3760233ef24ae8bfa8e33940/qgames.log"
	rawData := gist.ReadGist()
	rawDataByGame := rawData.ToGames()
	lines := rawDataByGame[1].ToLines()
	var gl data.GameLines = lines
	// for _, v := range gl {
	// 	fmt.Println(v)
	// }
	players := gl.Players()
	fmt.Println(players)
}

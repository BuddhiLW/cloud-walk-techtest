package data

import (
	"strings"
)

func (r RawData) ToGames() RawGames {
	dataBySplit := RawGames{}
	var gameCount int = 0

	splits := strings.Split(r.String(), "------------------------------------------------------------")

	for _, v := range splits {
		// has regex =InitGame:= or not -- determine if it's a game
		if strings.Contains(v, "InitGame:") {
			dataBySplit[gameCount] = RawGame(v)
			gameCount++
		}
	}
	return dataBySplit
}

func (rg RawGame) ToLines() []Line {
	var lines []Line
	for _, v := range strings.Split(rg.String(), "\n") {
		line := Line(v)
		lines = append(lines, line)
	}
	return lines
}

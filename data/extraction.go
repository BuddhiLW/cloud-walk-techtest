package data

import (
	"errors"
	"regexp"
	"strings"

	p "github.com/BuddhiLW/cloud-walk-techtest/players"
)

func (r RawData) String() string {
	return string(r)
}

func (r RawData) SplitByGame() RawGames {
	var dataBySplit RawGames
	var gameCount int = 0

	splits := strings.Split(r.String(), "------------------------------------------------------------")

	for _, v := range splits {
		// has regex =InitGame:= or not -- determine if it's a game
		if strings.Contains(v, "InitGame:") {
			gameCount++
			rawGame := RawGame(v)
			dataBySplit[gameCount] = rawGame
		}
	}
	return dataBySplit
}

func (r RawGame) String() string {
	return string(r)
}

func (l Line) String() string {
	return string(l)
}

func (l Line) HasPlayer() bool {
	return strings.Contains(l.String(), "ClientUserinfoChanged")
}

func (l Line) ExtractPlayer() string {
	reg := regexp.MustCompile(`n\\(.*?)\\t`)
	return reg.FindString(l.String())
}

func (l Line) Player() (p.Player, error) {
	if l.HasPlayer() {
		player := p.NewPlayer(l.ExtractPlayer())
		return *player, nil
	} else {
		return p.Player{}, errors.New("No player found")
	}
}

func containPlayer(playersMap map[string]p.Player, name string) bool {
	_, ok := playersMap[name]
	return ok
}

func (lines GameLines) Players() p.Players {
	var players map[string]p.Player

	// Extract players from the lines
	for _, v := range lines {
		if v.HasPlayer() {
			playerName := v.ExtractPlayer()
			if !containPlayer(players, playerName) {
				player := p.NewPlayer(playerName)
				players[playerName] = *player
			}
		}
	}

	return players
}

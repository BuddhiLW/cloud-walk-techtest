package data

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	p "github.com/BuddhiLW/cloud-walk-techtest/players"
)

func (r RawData) String() string {
	return string(r)
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

func (l Line) HasKill() bool {
	return strings.Contains(l.String(), "Kill")
}

func (l Line) ExtractPlayer() string {
	reg := regexp.MustCompile(`n\\(.*?)\\t`)
	s := reg.FindString(l.String())
	// Remove the first and last 2 characters from the string
	// in order to get the player name
	return s[2 : len(s)-2]
}

func (l Line) ExtractAction() (Killer, Victim, Mode) {
	reg := regexp.MustCompile(`([^:]*?) killed (.*?) by (MOD_.*)`)
	s := reg.FindStringSubmatch(l.String())
	killer, killed, mode := Killer(strings.TrimSpace(s[1])), Victim(s[2]), Mode(s[3])
	return killer, killed, mode
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
	var players p.Players = map[string]p.Player{}

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

func (lines GameLines) Kills(players *p.Players) {
	for _, v := range lines {
		if v.HasKill() {
			_, _, _ = v.ExtractAction() // killer, killed, action

			// player := (*players)[killer]
			// player.AddKill(killed, action)
		}
	}
}

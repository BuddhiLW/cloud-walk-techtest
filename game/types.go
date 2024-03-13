package game

import player "github.com/BuddhiLW/cloud-walk-techtest/players"

type Game struct {
	Players player.Players
}

type EncodeGame struct {
	TotalKills int            `json:"total_kills"`
	Players    []string       `json:"players"`
	Kills      map[string]int `json:"kills"`
}

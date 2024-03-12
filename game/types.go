package game

import player "github.com/BuddhiLW/cloud-walk-techtest/players"

type Game struct {
	// TotalKills int
	// Kills      map[string]int
	Players map[string]player.Player
}

type EncodeGame struct {
	TotalKills int            `json:"total_kills"`
	Players    []string       `json:"players"`
	Kills      map[string]int `json:"kills"`
}

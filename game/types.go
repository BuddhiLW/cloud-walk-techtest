package game

import player "github.com/BuddhiLW/cloud-walk-techtest/players"

type Game struct {
	Players player.Players
}

type Kill int
type Kills map[string]Kill

type EncodeGame struct {
	TotalKills int      `json:"total_kills"`
	Players    []string `json:"players"`
	Kills      Kills    `json:"kills"`
}

type EncodeGames map[string]*EncodeGame

type Name string

type Rank struct {
	Name     Name `json:"name"`
	Kills    int  `json:"kills"`
	Position int  `json:"position"`
}

type GameRank map[Name]*Rank

// type EncodeGameRank map[string]*GameRank

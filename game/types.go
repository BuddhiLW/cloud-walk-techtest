package game

type Game struct {
	// TotalKills int
	// Kills      map[string]int
	Players map[string]Player
}

type EncodeGame struct {
	TotalKills int            `json:"total_kills"`
	Players    []string       `json:"players"`
	Kills      map[string]int `json:"kills"`
}

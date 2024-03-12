package player

type Player struct {
	Name    string         `json:"name"`
	Kills   int            `json:"kills"`
	Victims map[string]int `json:"victims"`
}

type Players map[string]Player

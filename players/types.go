package player

type Player struct {
	Name    string  `json:"name"`
	Kills   int     `json:"kills"`
	Victims Victims `json:"victims"`
}

type Players map[string]*Player
type Victims map[string]Ways
type Ways map[string]int

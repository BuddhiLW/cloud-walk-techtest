package player

type Player struct {
	Name    string         `json:"name"`
	Kills   int            `json:"kills"`
	Victims map[string]Way `json:"victims"`
}

type Players map[string]Player
type Way map[string]int

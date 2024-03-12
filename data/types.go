package data

type Gist string
type RawData string
type RawGame string
type RawGames map[int]RawGame
type Line string
type GameLines []Line

type Killer string
type Victim string
type Mode string

func (k Killer) String() string {
	return string(k)
}

func (v Victim) String() string {
	return string(v)
}

func (m Mode) String() string {
	return string(m)
}

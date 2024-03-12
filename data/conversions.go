package data

import "strings"

func (gl GameLines) Append(line Line) {
	gl[len(gl)] = line
}

func (rg RawGame) ToLines() []Line {
	var lines []Line
	for _, v := range strings.Split(rg.String(), "\n") {
		line := Line(v)
		lines = append(lines, line)
	}
	return lines
}

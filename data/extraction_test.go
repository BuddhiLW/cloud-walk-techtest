package data

import (
	p "github.com/BuddhiLW/cloud-walk-techtest/players"
	"testing"
)

func TestHasPlayer(t *testing.T) {
	line := Line("ClientUserinfoChanged: 2 n\\Isgalamido\\t\\0\\model\\uriel/zael\\hmodel\\uriel/zael\\g_redteam\\\\g_blueteam\\\\c1\\5\\c2\\5\\hc\\100\\w\\0\\l\\0\\tt\\0\\tl\\0")
	expected := true
	result := line.HasPlayer()
	if result != expected {
		t.Errorf("Expected: %t, Got: %t", expected, result)
	}
}

func TestExtractPlayer(t *testing.T) {
	line := Line("ClientUserinfoChanged: 2 n\\Isgalamido\\t\\0\\model\\uriel/zael\\hmodel\\uriel/zael\\g_redteam\\\\g_blueteam\\\\c1\\5\\c2\\5\\hc\\100\\w\\0\\l\\0\\tt\\0\\tl\\0")
	expected := "Isgalamido"
	result := line.ExtractPlayer()
	if result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}
}

func TestPlayers(t *testing.T) {
	var gist Gist = "https://gist.githubusercontent.com/cloudwalk-tests/be1b636e58abff14088c8b5309f575d8/raw/df6ef4a9c0b326ce3760233ef24ae8bfa8e33940/qgames.log"
	rawData := gist.ReadGist()
	rawDataByGame := rawData.ToGames()
	lines := rawDataByGame[1].ToLines()
	var gl GameLines = lines

	players := gl.Players()
	emptyMap := map[string]int{}
	expected := map[string]p.Player{
		"Dono da Bola": p.Player{"Dono da Bola", 0, emptyMap},
		"Isgalamido":   p.Player{"Isgalamido", 0, emptyMap},
		"Mocinha":      p.Player{"Mocinha", 0, emptyMap},
	}

	for _, v := range players {
		name := v.Name

		// expected contains the same player entries?
		if _, ok := expected[name]; !ok {
			t.Errorf("Expected: %v, Got: %v", expected[name], v)
		}
	}
}

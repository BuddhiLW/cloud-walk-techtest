package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// https://gist.github.com/cloudwalk-tests/be1b636e58abff14088c8b5309f575d8

func ReadGist(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(body)
}

func main() {
	rawData := ReadGist("https://gist.githubusercontent.com/cloudwalk-tests/be1b636e58abff14088c8b5309f575d8/raw/df6ef4a9c0b326ce3760233ef24ae8bfa8e33940/qgames.log")
	dataBySplit := strings.Split(rawData, "------------------------------------------------------------")

	gameCount := 0
	for _, v := range dataBySplit {
		// has regex =InitGame:= or not -- determine if it's a game
		if strings.Contains(v, "InitGame:") {
			gameData := strings.Split(v, "\n")
			for n, v := range gameData {
				fmt.Println(n, string(v))
			}

			// Count the number of games
			gameCount++
		}
	}
	fmt.Println("Game Count: ", gameCount)
}

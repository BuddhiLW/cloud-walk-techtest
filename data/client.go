package data

import (
	"io"
	"net/http"
	"os"
)

func (g Gist) url() string {
	return string(g)
}

func (g Gist) ReadGist() RawData {
	resp, err := http.Get(g.url())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var rawData RawData = RawData(body)
	return rawData
}

type Path string

func (path Path) String() string {
	return string(path)
}

func (path Path) ReadFile() RawData {
	content, err := os.ReadFile(path.String())
	if err != nil {
		panic(err)
	}

	return RawData(content)
}

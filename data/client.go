package data

import (
	"io"
	"net/http"
)

func (g Gist) url() string {
	return string(g)
}

func (g Gist) ReadGist() string {
	resp, err := http.Get(g.url())
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

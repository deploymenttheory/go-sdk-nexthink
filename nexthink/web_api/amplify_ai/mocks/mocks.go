package mocks

import (
	"embed"
	"github.com/jarcoal/httpmock"
	"net/http"
)

//go:embed fixtures/*.json
var fixtures embed.FS

func Fixture(name string) []byte {
	b, err := fixtures.ReadFile("fixtures/" + name + ".json")
	if err != nil {
		panic(err)
	}
	return b
}
func Responder(status int, name string) httpmock.Responder {
	return func(_ *http.Request) (*http.Response, error) {
		r := httpmock.NewBytesResponse(status, Fixture(name))
		r.Header.Set("Content-Type", "application/json")
		return r, nil
	}
}

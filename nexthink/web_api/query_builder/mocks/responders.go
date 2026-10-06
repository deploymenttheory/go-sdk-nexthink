package mocks

import (
	"embed"
	"github.com/jarcoal/httpmock"
	"net/http"
)

//go:embed *.json
var fixtures embed.FS

func Fixture(name string) []byte {
	b, err := fixtures.ReadFile(name + ".json")
	if err != nil {
		panic(err)
	}
	return b
}
func Responder(status int, name string) httpmock.Responder {
	return func(r *http.Request) (*http.Response, error) {
		v := httpmock.NewBytesResponse(status, Fixture(name))
		v.Header.Set("Content-Type", "application/json")
		v.Header.Set("X-Request-ID", "fixture-request")
		return v, nil
	}
}

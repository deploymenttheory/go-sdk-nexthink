// Package mocks contains synthetic JSON fixtures; these are not lab captures.
package mocks

import (
	"embed"
	"net/http"

	"github.com/jarcoal/httpmock"
)

//go:embed *.json
var fixtures embed.FS

func Fixture(name string) []byte {
	data, err := fixtures.ReadFile(name + ".json")
	if err != nil {
		panic(err)
	}
	return data
}

func Responder(status int, name string) httpmock.Responder {
	return func(_ *http.Request) (*http.Response, error) {
		var data []byte
		if name != "" {
			data = Fixture(name)
		}
		response := httpmock.NewBytesResponse(status, data)
		response.Header.Set("Content-Type", "application/json")
		response.Header.Set("X-Request-ID", "fixture-request")
		return response, nil
	}
}

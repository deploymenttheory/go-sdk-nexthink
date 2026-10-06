package mocks

import "embed"

//go:embed *.json
var files embed.FS

func Fixture(n string) []byte {
	b, e := files.ReadFile(n + ".json")
	if e != nil {
		panic(e)
	}
	return b
}

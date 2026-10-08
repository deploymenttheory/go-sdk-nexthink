package mocks

import "embed"

//go:embed *.json
var files embed.FS

func Fixture(name string) []byte {
	b, e := files.ReadFile(name + ".json")
	if e != nil {
		panic(e)
	}
	return b
}

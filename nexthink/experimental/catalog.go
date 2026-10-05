package experimental

import (
	_ "embed"
	"encoding/json"
)

// Operation records an observed HTTP contract. Authentication records live
// evidence, not a guarantee for every tenant or role.
type Operation struct {
	ID             string `json:"id"`
	Method         string `json:"method"`
	Path           string `json:"path"`
	Evidence       string `json:"evidence"`
	Authentication string `json:"authentication"`
}

//go:embed operations.json
var catalogJSON []byte

var catalog = func() []Operation {
	var result []Operation
	if err := json.Unmarshal(catalogJSON, &result); err != nil {
		panic(err)
	}
	return result
}()

// Operations returns a copy of the implemented experimental operation catalog.
func Operations() []Operation { return append([]Operation(nil), catalog...) }
func operationByID(id string) (Operation, bool) {
	for _, op := range catalog {
		if op.ID == id {
			return op, true
		}
	}
	return Operation{}, false
}

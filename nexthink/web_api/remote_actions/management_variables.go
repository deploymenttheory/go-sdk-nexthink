package remote_actions

import "encoding/json"

func managementVariables(request any) (map[string]any, error) {
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	var variables map[string]any
	err = json.Unmarshal(b, &variables)
	return variables, err
}

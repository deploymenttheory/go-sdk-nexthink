package query_builder

import "encoding/json"

func (t *Transform) UnmarshalJSON(data []byte) error {
	type known Transform
	var value known
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "objectURI")
	delete(fields, "timeSeriesURI")
	delete(fields, "metric")
	value.AdditionalFields = fields
	*t = Transform(value)
	return nil
}
func (t Transform) MarshalJSON() ([]byte, error) {
	type known Transform
	data, err := json.Marshal(known(t))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err = json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	for k, v := range t.AdditionalFields {
		if k == "objectURI" || k == "timeSeriesURI" || k == "metric" {
			continue
		}
		fields[k] = v
	}
	return json.Marshal(fields)
}

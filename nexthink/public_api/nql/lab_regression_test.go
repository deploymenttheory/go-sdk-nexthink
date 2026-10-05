package nql

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestBuilderKeepsFiltersInTheirEventScope(t *testing.T) {
	query := NewQueryBuilder().FromDevices().With("execution.crashes during past 7d").WhereEquals("binary.name", `a"b.exe`).ComputeSum("crashes", "number_of_crashes").WhereGreater("crashes", "1").Build()
	filter := strings.Index(query, `binary.name == "a\"b.exe"`)
	compute := strings.Index(query, "| compute")
	if filter < 0 || filter > compute || strings.Index(query, "crashes > 1") < compute {
		t.Fatalf("filter scope or literal escaping changed: %s", query)
	}
}

func TestCSVToJSONPreservesTextAndRejectsAmbiguousHeaders(t *testing.T) {
	b, err := csvToJSON([]byte("id,label\n001,\"comma,quoted\"\n"))
	if err != nil {
		t.Fatal(err)
	}
	var rows []map[string]string
	if err = json.Unmarshal(b, &rows); err != nil {
		t.Fatal(err)
	}
	if rows[0]["id"] != "001" || rows[0]["label"] != "comma,quoted" {
		t.Fatal("CSV conversion changed values")
	}
	if _, err = csvToJSON([]byte("id,id\n1,2\n")); err == nil {
		t.Fatal("duplicate header silently overwrote a column")
	}
}

func FuzzRaggedV1Results(f *testing.F) {
	f.Add([]byte(`{"headers":["a","b"],"data":[[1],[]]}`), 0, 1)
	f.Add([]byte(`{"headers":[],"data":[]}`), -1, 0)
	f.Fuzz(func(t *testing.T, raw []byte, row, col int) {
		var response ExecuteNQLV1Response
		if json.Unmarshal(raw, &response) != nil {
			return
		}
		rs := NewV1ResultSet(&response)
		_, err := rs.Get(row, col)
		valid := row >= 0 && row < len(response.Data) && col >= 0 && col < len(response.Headers) && col < len(response.Data[row])
		if valid != (err == nil) {
			t.Fatal("bounds check disagrees with available cells")
		}
		_ = rs.ToV2Format()
	})
}

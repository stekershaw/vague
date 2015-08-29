package vague_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stekershaw/vague"
)

var intTests = []struct {
	json     []byte    // input
	expected vague.Int // expected result
}{
	{[]byte(`1`), 1},
	{[]byte(`"1"`), 1},
	{[]byte(`""`), 0},
}

func TestVagueIntUnmarshal(t *testing.T) {
	for _, tt := range intTests {
		var actual vague.Int
		if err := json.Unmarshal(tt.json, &actual); err != nil {
			t.Errorf("json.Unmarshal(\"%+v\") failed: %s", tt.json, err)
		}
		if actual != tt.expected {
			t.Errorf("json.Unmarshal(%d): expected %d, actual %d", tt.json, tt.expected, actual)
		}
	}
}

func TestVagueIntDecoder(t *testing.T) {
	for _, tt := range intTests {
		var actual vague.Int
		buf := bytes.NewBuffer(tt.json)
		decoder := json.NewDecoder(buf)
		if err := decoder.Decode(&actual); err != nil {
			t.Errorf("json.Decoder.Decode(\"%+v\") failed: %s", tt.json, err)
		}
		if actual != tt.expected {
			t.Errorf("json.Decoder.Decode(%d): expected %d, actual %d", tt.json, tt.expected, actual)
		}
	}
}

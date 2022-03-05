package jatomic

import (
	"encoding/json"
	"testing"
)

type jastruct struct {
	Bool   Bool   `json:"bool"`
	Int32  Int32  `json:"int32"`
	Int64  Int64  `json:"int64"`
	Uint32 Uint32 `json:"uint32"`
	Uint64 Uint64 `json:"uint64"`
}

func TestJatomic(t *testing.T) {
	var j jastruct
	j.Bool.Store(true)
	j.Int32.Store(1)
	j.Int64.Store(1)
	j.Uint32.Store(1)
	j.Uint64.Store(1)

	jb, err := json.Marshal(j)
	if err != nil {
		t.Fatalf("unexpected json.Marshal error: %s", err)
	}

	const expected = `{"bool":true,"int32":1,"int64":1,"uint32":1,"uint64":1}`
	if string(jb) != expected {
		t.Fatalf("marshalled output mismatch:\nexpected: %s\ngot: %s\n", expected, string(jb))
	}

	j = jastruct{}
	if err := json.Unmarshal([]byte(expected), &j); err != nil {
		t.Fatalf("unexpected json.Unmarshal error: %s", err)
	}

	j.Int32.Add(1)
	j.Int64.Add(1)
	j.Uint32.Add(1)
	j.Uint64.Add(1)

	jb, err = json.Marshal(j)
	if err != nil {
		t.Fatalf("unexpected second json.Marshal error: %s", err)
	}

	const expected2 = `{"bool":true,"int32":2,"int64":2,"uint32":2,"uint64":2}`
	if string(jb) != expected2 {
		t.Fatalf("unmarshalled output mismatch:\nexpected: %s\ngot: %s\n", expected2, string(jb))
	}
}

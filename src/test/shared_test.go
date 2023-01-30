package test

import (
	"bytes"
	"encoding/json"
	"io"
)

func toReader(v interface{}) io.Reader {
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(v)
	return &buf
}

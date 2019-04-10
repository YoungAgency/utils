package utils

import (
	"encoding/json"
	"errors"
	"fmt"
)

// returns a slice of bytes that rapresent a key:value json
// string(GetJSONBytes("foo", "bar") = ""foo":"bar""
// string(GetJSONBytes("foo", 5) = ""foo":5"
func GetJSONBytes(key string, val interface{}) (b []byte, err error) {
	b, err = json.Marshal(val)
	if err != nil {
		return
	}
	b = []byte(fmt.Sprintf("\"%s\":%s", key, b))
	return
}

func CloseJSON(b []byte) error {
	if b == nil {
		return errors.New("b param is nil")
	}
	l := len(b)
	if l > 0 {
		if b[l-1] == byte(',') {
			b = b[:l-1]
		}
	}
	b = append(b, byte('}'))
	return nil
}

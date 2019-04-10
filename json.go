package utils

import (
	"encoding/json"
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

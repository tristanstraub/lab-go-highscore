package json

import (
	"encoding/json"
)

func Decode(src []byte, factory func() interface{}) interface{} {
	dst := factory()
	json.Unmarshal(src, dst)
	return dst
}

func Encode(src interface{}) ([]byte, error) {
	return json.Marshal(src)
}

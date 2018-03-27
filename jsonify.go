package utils

import "encoding/json"

// Jsonify the object
func Jsonify(v interface{}) string {
	data, _ := json.Marshal(v)
	return string(data)
}

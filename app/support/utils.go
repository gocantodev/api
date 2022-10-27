package support

import "encoding/json"

func ToJson(receiver interface{}) string {
	value, err := json.Marshal(receiver)

	if err != nil {
		return ""
	}

	return string(value[:])
}

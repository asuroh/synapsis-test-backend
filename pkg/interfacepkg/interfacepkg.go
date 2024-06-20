package interfacepkg

import (
	"encoding/json"
)

// Marshall convert interface json to string
func Marshall(data interface{}) (res string) {
	name, err := json.Marshal(data)
	if err != nil {
		return res
	}
	res = string(name)

	return res
}

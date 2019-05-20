package strings

import "encoding/json"

//ToString format object to string
func ToString(obj interface{}) string {
	var str string
	if obj == nil {
		return str
	}
	byteArray, err := json.Marshal(obj)
	if err != nil {
		return err.Error()
	}
	return string(byteArray)
}


package arr

import "reflect"

// Contain 判断slice中是否包含某个item
func Contain(value interface{}, arr interface{}) bool {
	switch reflect.TypeOf(arr).Kind() {
	case reflect.Slice:
		a := reflect.ValueOf(arr)
		for i := 0; i < a.Len(); i++ {
			if reflect.DeepEqual(value, a.Index(i).Interface()) {
				return true
			}
		}
	}
	return false
}

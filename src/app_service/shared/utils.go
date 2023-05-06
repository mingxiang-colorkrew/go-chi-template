package shared

import "reflect"

func hasNotNilField(s interface{}) bool {
	val := reflect.ValueOf(s).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if !field.IsNil() {
			return true
		}
	}

	return false
}

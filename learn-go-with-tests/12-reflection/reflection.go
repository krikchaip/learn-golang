package reflection

import "reflect"

func Walk(x any, fn func(string)) {
	val := reflect.ValueOf(x)
	field := val.Field(0)
	fn(field.String())
}

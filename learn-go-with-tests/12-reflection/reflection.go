package reflection

import "reflect"

func Walk(x any, fn func(string)) {
	// v1(x, fn)
	v2(x, fn)
}

func v1(x any, fn func(string)) {
	// ?? get the value part of x (there are type, value, ...)
	val := reflect.ValueOf(x)

	// ?? assumes that `val` is a struct, get the first struct field value
	field := val.Field(0)

	// ?? assumes that `field` is a string
	fn(field.String())
}

func v2(x any, fn func(string)) {
	val := reflect.ValueOf(x)

	// ?? assumes that x is a struct and iterate through each of its field
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		// ?? making sure that this particular field is a string
		if field.Kind() == reflect.String {
			fn(field.String())
		}
	}
}

package reflection

import "reflect"

// takes a struct x and calls fn for all strings fields
// found inside recursively.
func Walk(x any, fn func(string)) {
	// v1(x, fn)
	// v2(x, fn)
	v3(x, fn)
}

func v1(x any, fn func(string)) {
	// ?? get the value part of x (there are reflect.Type and reflect.Value)
	val := reflect.ValueOf(x)

	// ?? assumes that struct x has atleast one field
	field := val.Field(0)

	// ?? assumes that the `field` has a type of string
	fn(field.String())
}

func v2(x any, fn func(string)) {
	val := reflect.ValueOf(x)

	// ?? iterate through each of the struct fields
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		// ?? making sure that this particular field
		// ?? has one of the following types
		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			// ?? convert from reflect.Value to interface{}
			v2(field.Interface(), fn)
		}
	}
}

func v3(x any, fn func(string)) {
	val := reflect.ValueOf(x)

	// ?? dereference the underlying value of a pointer
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			v3(field.Interface(), fn)
		}
	}
}

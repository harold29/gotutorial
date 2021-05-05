package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	// fn("I still can't believe South Korea beat Germany 2-0 to put them last in their group")
	val := reflect.ValueOf(x)
	field := val.Field(0)
	fn(field.String())
}
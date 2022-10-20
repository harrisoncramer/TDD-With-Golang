package reflection

import "reflect"

/*
Reflection can be used to interact with data
where the types of the data is not known at compile time.
*/
func walk(x interface{}, f func(s string)) {

	val := getValue(x)

	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			/* The interface method returns the underlying value */
			walk(val.Index(i).Interface(), f)
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), f)
		}
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walk(val.Field(i).Interface(), f)
		}
	case reflect.String:
		f(val.String())
	}

}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Pointer {
		return val.Elem()
	}
	return val
}

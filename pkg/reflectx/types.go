package reflectx

import "reflect"

// IsNumber reports whether v is a numeric value or can be converted to an int.
//
// Example:
//
//	reflectx.IsNumber(42)   // true
//	reflectx.IsNumber("42") // true (convertible)
func IsNumber(v any) bool {
	switch v.(type) {
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64:
		return true
	}

	vval := reflect.ValueOf(v)
	return vval.CanConvert(reflect.TypeOf(int(0)))
}

// IsString reports whether v is a string or string-like type.
func IsString(v any) bool {
	switch v.(type) {
	case string:
		return true
	}
	if reflect.TypeOf(v).Kind() == reflect.String {
		return true
	}
	tstr := reflect.TypeOf("")
	vval := reflect.ValueOf(v)
	if !IsNumber(v) && vval.CanConvert(tstr) {
		return true
	}

	return false
}

// IsBool reports whether v is a boolean value.
func IsBool(v any) bool {
	switch v.(type) {
	case bool:
		return true
	}

	return false
}

// IsPrimitive reports whether v is a primitive value (bool, number, or string).
func IsPrimitive(v any) bool {
	return IsBool(v) || IsNumber(v) || IsString(v)
}

// IsSlice reports whether v is a slice or pointer to a slice.
//
// Example:
//
//	reflectx.IsSlice([]int{1, 2, 3})    // true
//	reflectx.IsSlice(&[]string{"a","b"}) // true
func IsSlice(v any) bool {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	switch val.Kind() {
	case reflect.Slice:
		return true
	}

	return false
}

// IsStruct reports whether v is a struct or pointer to a struct.
func IsStruct(v any) bool {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val.Kind() == reflect.Struct
}

func TypeIsElementWrapper(t reflect.Type) bool {
	kind := t.Kind()
	return kind == reflect.Array ||
		kind == reflect.Chan ||
		kind == reflect.Map ||
		kind == reflect.Ptr ||
		kind == reflect.Slice
}

func ValueIsElementWrapper(v reflect.Value) bool {
	kind := v.Kind()
	return kind == reflect.Ptr ||
		kind == reflect.Interface
}

// ValueOfElement drills down on the input v to get
// the reflect.Value of the fundumental element;
// e.g. if v is a Ptr, it will get the reflect.Value of the
// dereferenced type of v instead.
func ValueOfElement(v any) reflect.Value {
	value := reflect.ValueOf(v)
	for value.Kind() == reflect.Ptr || value.Kind() == reflect.Interface {
		if value.Kind() == reflect.Ptr {
			value = value.Elem()
		}
		if value.Kind() == reflect.Interface {
			value = reflect.ValueOf(value.Interface())
		}
	}
	return value
}

// TypeOfElement drills down on the input v to get
// the reflect.Type of the fundumental element;
// e.g. if v is a Ptr, it will get the reflect.Type of the
// dereferenced type of v instead.
func TypeOfElement(v any) reflect.Type {
	typeOf := reflect.TypeOf(v)
	for TypeIsElementWrapper(typeOf) {
		typeOf = typeOf.Elem()
	}
	return typeOf
}

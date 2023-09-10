package reflectx

import (
	"reflect"
)

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

func IsBool(v any) bool {

	switch v.(type) {
	case bool:
		return true
	}

	return false
}

func IsPrimitive(v any) bool {

	return IsBool(v) || IsNumber(v) || IsString(v)
}

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

func IsStruct(v any) bool {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val.Kind() == reflect.Struct
}

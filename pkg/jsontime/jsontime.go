package jsontime

import (
	"encoding/json"
	"reflect"
	"strings"
	"time"

	"github.com/tartale/go/pkg/structs"
)

var DefaultFormat = time.RFC3339

/*
Time allows marshal/unmarshal of a time field using a custom format.
Inspired by this SO post:

	https://stackoverflow.com/a/20510912/1258206

The default time format is the RFC3339 format that is currently expected
by the default json marshaler.

Example:

	type MyStruct struct {
		// The 'format' tag specifies the time layout expected in a JSON document
		MyTime formattedtime.Time `json:"myTime" format:"2006-01-02"`
	}
	myStruct := MyStruct{}
	myJson := `{"foo":"foo","bar":"1976-07-31"}`
	err := jsontime.UnmarshalJSON([]byte(myJson), &myStruct)
*/
type Time struct {
	time.Time `json:"-"`
	Raw       string `json:"-"`
}

func New(t time.Time) *Time {
	return &Time{Time: t}
}

func Now() *Time {
	return &Time{Time: time.Now()}
}

func (t *Time) MarshalJSON() ([]byte, error) {

	unquoted := strings.Trim(t.Raw, `"`)
	return []byte(`"` + unquoted + `"`), nil
}

func (t *Time) UnmarshalJSON(data []byte) error {

	t.Raw = string(data)
	return nil
}

func MarshalJSON(v any) ([]byte, error) {

	marshalTime(v)
	return json.Marshal(v)
}

func UnmarshalJSON(data []byte, v any) error {

	err := json.Unmarshal(data, v)
	if err != nil {
		return nil
	}
	unmarshalTime(v)

	return nil
}

func marshalTime(v any) {

	walkFn := func(field reflect.StructField, value reflect.Value) error {

		switch val := value.Interface().(type) {
		case Time:
			format := getFormat(field)
			newRawValue := val.Time.Format(format)
			raw := value.FieldByName("Raw")
			raw.SetString(newRawValue)
		}

		return nil
	}

	structs.Walk(v, walkFn)
}

func unmarshalTime(v any) {

	walkFn := func(field reflect.StructField, value reflect.Value) error {

		switch val := value.Interface().(type) {
		case Time:
			format := getFormat(field)
			unquoted := strings.Trim(val.Raw, `"`)
			newTime, err := time.Parse(format, unquoted)
			if err != nil {
				return err
			}
			t := value.FieldByName("Time")
			t.Set(reflect.ValueOf(newTime))
		}

		return nil
	}

	structs.Walk(v, walkFn)
}

func getFormat(f reflect.StructField) string {
	var format string
	if tag := f.Tag.Get("format"); tag != "" {
		format = tag
	} else {
		format = DefaultFormat
	}

	return format
}

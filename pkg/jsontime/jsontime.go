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
	time.Time
	Layout string `json:"-"`
}

func New(t time.Time) *Time {
	return &Time{Time: t}
}

func Now() *Time {
	return &Time{Time: time.Now()}
}

func (f *Time) Format() string {

	return f.Time.Format(f.Layout)
}

func (f *Time) ParseFrom(value string) error {

	newTime, err := time.Parse(f.Layout, value)
	if err != nil {
		return err
	}
	f.Time = newTime

	return nil
}

func (t *Time) MarshalJSON() ([]byte, error) {

	return []byte(`"` + t.Format() + `"`), nil
}

func (f *Time) UnmarshalJSON(data []byte) error {

	unquoted := strings.Trim(string(data), `"`)
	err := f.ParseFrom(unquoted)
	if err != nil {
		return err
	}

	return nil
}

func MarshalJSON(v any) ([]byte, error) {

	addTimeMarshaling(v)
	return json.Marshal(v)
}

func UnmarshalJSON(data []byte, v any) error {

	addTimeMarshaling(v)
	return json.Unmarshal(data, v)
}

func addTimeMarshaling(v any) {

	setLayout := func(field reflect.StructField, value reflect.Value) error {

		switch value.Interface().(type) {
		case Time:
			layout := value.FieldByName("Layout")
			if tag := field.Tag.Get("format"); tag != "" {
				layout.SetString(tag)
			} else {
				layout.SetString(DefaultFormat)
			}
		}

		return nil
	}

	structs.Walk(v, setLayout)
}

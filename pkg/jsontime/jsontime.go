package jsontime

import (
	"encoding/json"
	"reflect"
	"strings"
	"time"

	"github.com/tartale/go/pkg/logz"
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
		MyTime jsontime.Time `json:"myTime" format:"2006-01-02"`
	}
	myStruct := MyStruct{}
	myJson := `{"foo":"foo","bar":"1976-07-31"}`
	err := jsontime.UnmarshalJSON([]byte(myJson), &myStruct)
*/
type Time struct {
	time.Time `json:"-"`
	Raw       string `json:"-"`
}

// New constructs a jsontime.Time from the given time.Time.
//
// Example:
//
//	t := time.Now()
//	jsTime := jsontime.New(t)
//	_ = jsTime
func New(t time.Time) *Time {
	return &Time{Time: t}
}

// Now returns a jsontime.Time set to the current instant.
func Now() *Time {
	return &Time{Time: time.Now()}
}

// MarshalJSON marshals the Time using the raw string value.
func (t *Time) MarshalJSON() ([]byte, error) {
	unquoted := strings.Trim(t.Raw, `"`)
	return []byte(`"` + unquoted + `"`), nil
}

// UnmarshalJSON stores the raw JSON string representation into t.Raw.
func (t *Time) UnmarshalJSON(data []byte) error {
	unquoted := strings.Trim(string(data), `"`)
	t.Raw = unquoted
	return nil
}

// MarshalJSONIndent marshals v to JSON with indentation after first
// rewriting any jsontime.Time fields into their string representations.
//
// Example:
//
//	type Payload struct {
//		When jsontime.Time `json:"when" format:"2006-01-02"`
//	}
//	p := Payload{When: *jsontime.New(time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC))}
//	b, err := jsontime.MarshalJSONIndent(&p, "", "  ")
//	_ = b
//	_ = err
func MarshalJSONIndent(v any, prefix, indent string) ([]byte, error) {
	MarshalTime(v)
	return json.MarshalIndent(v, prefix, indent)
}

// MarshalJSON marshals v to JSON after rewriting any jsontime.Time fields
// into their string representations.
func MarshalJSON(v any) ([]byte, error) {
	MarshalTime(v)
	return json.Marshal(v)
}

// UnmarshalJSON unmarshals data into v and then parses any jsontime.Time
// fields from their raw string values using the configured format tags.
func UnmarshalJSON(data []byte, v any) error {
	err := json.Unmarshal(data, v)
	if err != nil {
		return err
	}
	UnmarshalTime(v)

	return nil
}

// MarshalTime walks v and updates any jsontime.Time fields' Raw values
// based on the configured format tags or DefaultFormat.
func MarshalTime(v any) {
	walkFn := func(field reflect.StructField, value reflect.Value) error {
		switch val := value.Interface().(type) {
		case Time:
			format := getFormat(field)
			newRawValue := val.Time.Format(format)
			raw := value.FieldByName("Raw")
			raw.SetString(newRawValue)
			logz.Logger().Debugf("marshaled json time field; name: %s, rawValue: %s\n", field.Name, newRawValue)
		}

		return nil
	}

	structs.Walk(v, walkFn)
}

// UnmarshalTime walks v and parses any jsontime.Time fields from their
// Raw string representations using the configured format tags or DefaultFormat.
func UnmarshalTime(v any) {
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
			logz.Logger().Debugf("unmarshaled json time field; name: %s, newTime: %s\n", field.Name, newTime)
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

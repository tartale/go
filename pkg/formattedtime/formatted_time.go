package formattedtime

import (
	"strings"
	"time"
)

// Time allows marshal/unmarshal of a time field using a custom format.
// Inspired by this SO post:
//
//	https://stackoverflow.com/a/20510912/1258206
type Time struct {
	time.Time
	Layout string
}

func New(t time.Time, layout string) *Time {

	return &Time{t, layout}
}

func Zero(layout string) *Time {

	return &Time{time.Time{}, layout}
}

func Now(layout string) *Time {

	return &Time{time.Now(), layout}
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

func (f *Time) MarshalText() ([]byte, error) {

	return []byte(f.Format()), nil
}

func (f *Time) UnmarshalText(text []byte) error {

	return f.ParseFrom(string(text))
}

func (f *Time) MarshalJSON() ([]byte, error) {

	return []byte(`"` + f.Format() + `"`), nil
}

func (f *Time) UnmarshalJSON(data []byte) error {

	unquoted := strings.Trim(string(data), `"`)
	err := f.ParseFrom(unquoted)
	if err != nil {
		return err
	}

	return nil
}

package formattedtime

import (
	"strings"
	"time"
)

// FormattedTime allows marshal/unmarshal of a time field using a custom format.
// Inspired by this SO post:
//
//	https://stackoverflow.com/a/20510912/1258206
type FormattedTime struct {
	time.Time
	Layout string
}

func New(t time.Time, layout string) *FormattedTime {

	return &FormattedTime{t, layout}
}

func Zero(layout string) *FormattedTime {

	return &FormattedTime{time.Time{}, layout}
}

func Now(layout string) *FormattedTime {

	return &FormattedTime{time.Now(), layout}
}

func (f *FormattedTime) Format() string {

	return f.Time.Format(f.Layout)
}

func (f *FormattedTime) ParseFrom(value string) error {

	newTime, err := time.Parse(f.Layout, value)
	if err != nil {
		return err
	}
	f.Time = newTime

	return nil
}

func (f *FormattedTime) MarshalText() ([]byte, error) {

	return []byte(f.Format()), nil
}

func (f *FormattedTime) UnmarshalText(text []byte) error {

	return f.ParseFrom(string(text))
}

func (f *FormattedTime) MarshalJSON() ([]byte, error) {

	return []byte(`"` + f.Format() + `"`), nil
}

func (f *FormattedTime) UnmarshalJSON(data []byte) error {

	unquoted := strings.Trim(string(data), `"`)
	err := f.ParseFrom(unquoted)
	if err != nil {
		return err
	}

	return nil
}

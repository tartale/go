package formattedtime

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Foo string        `json:"foo,omitempty"`
	Bar FormattedTime `json:"bar,omitempty"`
}

func TestFormattedTime_JSONMarshal(t *testing.T) {

	myTime, _ := time.Parse(time.RFC3339, "1976-07-31T14:30:00Z")
	formattedTime := FormattedTime{myTime, time.Kitchen}

	testStruct := TestStruct{Foo: "foo", Bar: formattedTime}
	testJson, err := json.Marshal(testStruct)
	assert.Nil(t, err)
	assert.Equal(t, `{"foo":"foo","bar":"2:30PM"}`, string(testJson))

	testStruct.Bar.Layout = time.DateOnly
	testJson, err = json.Marshal(testStruct)
	assert.Nil(t, err)
	assert.Equal(t, `{"foo":"foo","bar":"1976-07-31"}`, string(testJson))

}

func TestJsonTime_JSONUnmarshal(t *testing.T) {

	formattedTime := Now(time.RFC3339)
	testJson := `{"foo":"foo","bar":"1976-07-31T14:30:00Z"}`
	testStruct := TestStruct{Foo: "", Bar: *formattedTime}

	err := json.Unmarshal([]byte(testJson), &testStruct)
	assert.Nil(t, err)
	assert.Equal(t, time.July, formattedTime.Month())

	formattedTime.Layout = time.Kitchen
	testJson = `{"foo":"foo","bar":"2:30PM"}`
	testStruct.Bar = *formattedTime

	err = json.Unmarshal([]byte(testJson), &testStruct)
	assert.Nil(t, err)
	assert.Equal(t, 14, testStruct.Bar.Hour())
	assert.Equal(t, 30, testStruct.Bar.Minute())

}

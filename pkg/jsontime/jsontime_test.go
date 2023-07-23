package jsontime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Foo string `json:"foo,omitempty"`
	Bar Time   `json:"bar,omitempty" format:"2006-01-02"`
}

func TestAddTimeMarshalingl(t *testing.T) {
	testStruct := TestStruct{}
	addTimeMarshaling[TestStruct](&testStruct)

	assert.Equal(t, "2006-01-02", testStruct.Bar.Layout)
}

func TestJSONTime_JSONMarshal(t *testing.T) {

	myTime, _ := time.Parse(time.RFC3339, "1976-07-31T14:30:00Z")
	formattedTime := Time{Time: myTime}

	testStruct := TestStruct{Foo: "foo", Bar: formattedTime}
	testJson, err := MarshalJSON[TestStruct](&testStruct)

	assert.Nil(t, err)
	assert.Equal(t, `{"foo":"foo","bar":"1976-07-31"}`, string(testJson))

}

func TestJSONTime_JSONUnmarshal(t *testing.T) {

	testJson := `{"foo":"foo","bar":"1976-07-31"}`
	testStruct := TestStruct{}

	err := UnmarshalJSON[TestStruct]([]byte(testJson), &testStruct)
	assert.Nil(t, err)
	assert.Equal(t, 1976, testStruct.Bar.Year())
	assert.Equal(t, time.July, testStruct.Bar.Month())
	assert.Equal(t, 31, testStruct.Bar.Day())
}

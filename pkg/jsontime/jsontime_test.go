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

type TestNestedStruct struct {
	Foo    string     `json:"foo,omitempty"`
	Nested TestStruct `json:"nested,omitempty"`
}

type TestStructWithAnyField struct {
	Foo    any        `json:"foo,omitempty"`
	Nested TestStruct `json:"nested,omitempty"`
}

type TestStructWithNestedAnyField struct {
	Foo    string `json:"foo,omitempty"`
	Nested any    `json:"nested,omitempty"`
}

func TestAddTimeMarshaling(t *testing.T) {

	testStruct := TestStruct{}
	addTimeMarshaling(&testStruct)
	assert.Equal(t, "2006-01-02", testStruct.Bar.Layout)

}

func TestAddTimeMarshaling_NestedStruct(t *testing.T) {

	testStruct := TestNestedStruct{}
	addTimeMarshaling(&testStruct)
	assert.Equal(t, "2006-01-02", testStruct.Nested.Bar.Layout)

}

func TestAddTimeMarshaling_AnyField(t *testing.T) {

	testStruct := TestStructWithAnyField{}
	addTimeMarshaling(&testStruct)
	assert.Equal(t, "2006-01-02", testStruct.Nested.Bar.Layout)

}

func TestAddTimeMarshaling_NestedAnyField(t *testing.T) {

	nestedStruct := TestStruct{}
	testStruct := TestStructWithNestedAnyField{Nested: &nestedStruct}
	addTimeMarshaling(&testStruct)
	assert.Equal(t, "2006-01-02", testStruct.Nested.(*TestStruct).Bar.Layout)

}

func TestJSONTime_JSONMarshal(t *testing.T) {

	myTime, _ := time.Parse(time.RFC3339, "1976-07-31T14:30:00Z")
	jsonTime := Time{Time: myTime}

	testStruct := TestStruct{Foo: "foo", Bar: jsonTime}
	testJson, err := MarshalJSON[TestStruct](&testStruct)

	assert.Nil(t, err)
	assert.Equal(t, `{"foo":"foo","bar":"1976-07-31"}`, string(testJson))

}

func TestJSONTime_JSONMarshal_AnyField(t *testing.T) {

	myTime, _ := time.Parse(time.RFC3339, "1976-07-31T14:30:00Z")
	jsonTime := Time{Time: myTime}

	testStruct := TestStruct{Foo: "foo", Bar: jsonTime}
	testStructWithAnyField := TestStructWithAnyField{"foo", testStruct}
	testJson, err := MarshalJSON[TestStructWithAnyField](&testStructWithAnyField)

	assert.Nil(t, err)
	assert.Equal(t, `{"foo":"foo","nested":{"foo":"foo","bar":"1976-07-31"}}`, string(testJson))

}

func TestJSONTime_JSONMarshal_NestedAnyField(t *testing.T) {

	myTime, _ := time.Parse(time.RFC3339, "1976-07-31T14:30:00Z")
	jsonTime := Time{Time: myTime}

	testStruct := TestStruct{Foo: "foo", Bar: jsonTime}
	testStructWithNestedAnyField := TestStructWithNestedAnyField{"foo", &testStruct}
	testJson, err := MarshalJSON[TestStructWithNestedAnyField](&testStructWithNestedAnyField)

	assert.Nil(t, err)
	assert.Equal(t, `{"foo":"foo","nested":{"foo":"foo","bar":"1976-07-31"}}`, string(testJson))

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

func TestJSONTime_JSONUnmarshal_AnyField(t *testing.T) {

	testJson := `{"foo":"foo","nested":{"foo":"foo","bar":"1976-07-31"}}`
	testStructWithAnyField := TestStructWithAnyField{}

	err := UnmarshalJSON[TestStructWithAnyField]([]byte(testJson), &testStructWithAnyField)
	testStruct := testStructWithAnyField.Nested
	assert.Nil(t, err)
	assert.Equal(t, 1976, testStruct.Bar.Year())
	assert.Equal(t, time.July, testStruct.Bar.Month())
	assert.Equal(t, 31, testStruct.Bar.Day())

}

func TestJSONTime_JSONUnmarshal_NestedAnyField(t *testing.T) {

	testJson := `{"foo":"foo","nested":{"foo":"foo","bar":"1976-07-31"}}`
	testStructWithNestedAnyField := TestStructWithNestedAnyField{Nested: &TestStruct{}}

	err := UnmarshalJSON[TestStructWithNestedAnyField]([]byte(testJson), &testStructWithNestedAnyField)
	testStruct := testStructWithNestedAnyField.Nested.(*TestStruct)
	assert.Nil(t, err)
	assert.Equal(t, 1976, testStruct.Bar.Year())
	assert.Equal(t, time.July, testStruct.Bar.Month())
	assert.Equal(t, 31, testStruct.Bar.Day())

}

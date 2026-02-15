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

type TestStructWithSliceField struct {
	Foo   string       `json:"foo,omitempty"`
	Slice []TestStruct `json:"slice,omitempty"`
}

func TestMarshalTime(t *testing.T) {

	myTime, _ := time.Parse(time.RFC3339, "1976-07-31T14:30:00Z")
	jsonTime := New(myTime)
	testStruct := TestStruct{}
	testStruct.Bar = *jsonTime

	MarshalTime(&testStruct)

	assert.Equal(t, "1976-07-31", testStruct.Bar.Raw)
}

func TestMarshalTime_NestedStruct(t *testing.T) {

	myTime, _ := time.Parse(time.RFC3339, "1976-07-31T14:30:00Z")
	jsonTime := New(myTime)
	testStruct := TestNestedStruct{}
	testStruct.Nested.Bar = *jsonTime

	MarshalTime(&testStruct)

	assert.Equal(t, "1976-07-31", testStruct.Nested.Bar.Raw)
}

func TestMarshalTime_AnyField(t *testing.T) {

	myTime, _ := time.Parse(time.RFC3339, "1976-07-31T14:30:00Z")
	jsonTime := New(myTime)
	testStruct := TestStructWithAnyField{}
	testStruct.Nested.Bar = *jsonTime

	MarshalTime(&testStruct)
	assert.Equal(t, "1976-07-31", testStruct.Nested.Bar.Raw)
}

func TestMarshalTime_NestedAnyField(t *testing.T) {

	myTime, _ := time.Parse(time.RFC3339, "1976-07-31T14:30:00Z")
	jsonTime := New(myTime)

	nestedStruct := TestStruct{}
	nestedStruct.Bar = *jsonTime
	testStruct := TestStructWithNestedAnyField{Nested: &nestedStruct}

	MarshalTime(&testStruct)

	assert.Equal(t, "1976-07-31", testStruct.Nested.(*TestStruct).Bar.Raw)
}

func TestMarshalTime_SliceField(t *testing.T) {

	myTime, _ := time.Parse(time.RFC3339, "1976-07-31T14:30:00Z")
	jsonTime := New(myTime)
	item := TestStruct{}
	item.Bar = *jsonTime
	testStruct := TestStructWithSliceField{}
	testStruct.Slice = append(testStruct.Slice, item)

	MarshalTime(&testStruct)

	assert.Len(t, testStruct.Slice, 1)
	assert.Equal(t, "1976-07-31", testStruct.Slice[0].Bar.Raw)
}

func TestUnmarshalTime(t *testing.T) {

	testStruct := TestStruct{}
	testStruct.Bar.Raw = "1976-07-31"

	UnmarshalTime(&testStruct)

	assert.Equal(t, 1976, testStruct.Bar.Time.Year())
	assert.Equal(t, time.Month(7), testStruct.Bar.Time.Month())
	assert.Equal(t, 31, testStruct.Bar.Time.Day())
}

func TestUnmarshalTime_NestedStruct(t *testing.T) {

	testStruct := TestNestedStruct{}
	testStruct.Nested.Bar.Raw = "1976-07-31"

	UnmarshalTime(&testStruct)

	assert.Equal(t, 1976, testStruct.Nested.Bar.Time.Year())
	assert.Equal(t, time.Month(7), testStruct.Nested.Bar.Time.Month())
	assert.Equal(t, 31, testStruct.Nested.Bar.Time.Day())
}

func TestUnmarshalTime_AnyField(t *testing.T) {

	testStruct := TestStructWithAnyField{}
	testStruct.Nested.Bar.Raw = "1976-07-31"

	UnmarshalTime(&testStruct)

	assert.Equal(t, 1976, testStruct.Nested.Bar.Time.Year())
	assert.Equal(t, time.Month(7), testStruct.Nested.Bar.Time.Month())
	assert.Equal(t, 31, testStruct.Nested.Bar.Time.Day())
}

func TestUnmarshalTime_NestedAnyField(t *testing.T) {

	testStruct := TestStructWithNestedAnyField{}
	nestedStruct := TestStruct{}
	nestedStruct.Bar.Raw = "1976-07-31"
	testStruct.Nested = &nestedStruct

	UnmarshalTime(&testStruct)

	newNestedStruct := testStruct.Nested.(*TestStruct)
	assert.Equal(t, 1976, newNestedStruct.Bar.Time.Year())
	assert.Equal(t, time.Month(7), newNestedStruct.Bar.Time.Month())
	assert.Equal(t, 31, newNestedStruct.Bar.Time.Day())
}

func TestJSONTime_JSONMarshal(t *testing.T) {

	myTime, _ := time.Parse(time.RFC3339, "1976-07-31T14:30:00Z")
	jsonTime := Time{Time: myTime}
	testStruct := TestStruct{Foo: "foo", Bar: jsonTime}

	testJson, err := MarshalJSON(&testStruct)

	assert.Nil(t, err)
	assert.Equal(t, `{"foo":"foo","bar":"1976-07-31"}`, string(testJson))
}

func TestJSONTime_JSONMarshal_AnyField(t *testing.T) {

	myTime, _ := time.Parse(time.RFC3339, "1976-07-31T14:30:00Z")
	jsonTime := Time{Time: myTime}
	testStruct := TestStruct{Foo: "foo", Bar: jsonTime}
	testStructWithAnyField := TestStructWithAnyField{"foo", testStruct}

	testJson, err := MarshalJSON(&testStructWithAnyField)

	assert.Nil(t, err)
	assert.Equal(t, `{"foo":"foo","nested":{"foo":"foo","bar":"1976-07-31"}}`, string(testJson))
}

func TestJSONTime_JSONMarshal_NestedAnyField(t *testing.T) {

	myTime, _ := time.Parse(time.RFC3339, "1976-07-31T14:30:00Z")
	jsonTime := Time{Time: myTime}
	testStruct := TestStruct{Foo: "foo", Bar: jsonTime}
	testStructWithNestedAnyField := TestStructWithNestedAnyField{"foo", &testStruct}

	testJson, err := MarshalJSON(&testStructWithNestedAnyField)

	assert.Nil(t, err)
	assert.Equal(t, `{"foo":"foo","nested":{"foo":"foo","bar":"1976-07-31"}}`, string(testJson))
}

func TestJSONTime_JSONUnmarshal(t *testing.T) {

	testJson := `{"foo":"foo","bar":"1976-07-31"}`
	testStruct := TestStruct{}

	err := UnmarshalJSON([]byte(testJson), &testStruct)
	assert.Nil(t, err)
	assert.Equal(t, 1976, testStruct.Bar.Year())
	assert.Equal(t, time.July, testStruct.Bar.Month())
	assert.Equal(t, 31, testStruct.Bar.Day())
}

func TestJSONTime_JSONUnmarshal_AnyField(t *testing.T) {

	testJson := `{"foo":"foo","nested":{"foo":"foo","bar":"1976-07-31"}}`
	testStructWithAnyField := TestStructWithAnyField{}

	err := UnmarshalJSON([]byte(testJson), &testStructWithAnyField)

	testStruct := testStructWithAnyField.Nested
	assert.Nil(t, err)
	assert.Equal(t, 1976, testStruct.Bar.Year())
	assert.Equal(t, time.July, testStruct.Bar.Month())
	assert.Equal(t, 31, testStruct.Bar.Day())
}

func TestJSONTime_JSONUnmarshal_NestedAnyField(t *testing.T) {

	testJson := `{"foo":"foo","nested":{"foo":"foo","bar":"1976-07-31"}}`
	testStructWithNestedAnyField := TestStructWithNestedAnyField{Nested: &TestStruct{}}

	err := UnmarshalJSON([]byte(testJson), &testStructWithNestedAnyField)

	testStruct := testStructWithNestedAnyField.Nested.(*TestStruct)
	assert.Nil(t, err)
	assert.Equal(t, 1976, testStruct.Bar.Year())
	assert.Equal(t, time.July, testStruct.Bar.Month())
	assert.Equal(t, 31, testStruct.Bar.Day())
}

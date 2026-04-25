package structs

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	type Nested struct {
		S string
	}
	s := struct {
		A string
		B int
		C bool
		D Nested
		E int
	}{
		A: "a-value",
		B: 2,
		C: true,
		D: Nested{
			S: "s-value",
		},
		E: 999,
	}
	expected := `
A string a-value
B int 2
C bool true
D structs.Nested {s-value}
S string s-value
E int 999`

	var actual string
	err := Walk(s, func(field reflect.StructField, value reflect.Value) error {
		actual = fmt.Sprintf("%s\n%s %s %v", actual, field.Name, field.Type, value.Interface())
		return nil
	})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if actual != expected {
		t.Errorf("Unexpected answer: %v", actual)
	}
}

func TestWalk_NotAStruct(t *testing.T) {
	foo := []string{"foo"}

	defer func() {
		err := recover()
		if err == nil {
			t.Error("Passing a non struct into Walk should panic")
		}
	}()

	// this should panic. We are going to recover and and test it
	_ = Walk(foo, func(field reflect.StructField, value reflect.Value) error { return nil })
}

func TestWalk_Omitnested(t *testing.T) {
	type Nested struct {
		S string
	}
	s := struct {
		A string
		B int
		C bool
		D Nested `structs:",omitnested"`
		E int
	}{
		A: "a-value",
		B: 2,
		C: true,
		D: Nested{
			S: "s-value",
		},
		E: 999,
	}
	expected := `
A string a-value
B int 2
C bool true
D structs.Nested {s-value}
E int 999`

	var actual string
	err := Walk(s, func(field reflect.StructField, value reflect.Value) error {
		actual = fmt.Sprintf("%s\n%s %s %v", actual, field.Name, field.Type, value.Interface())
		return nil
	})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if actual != expected {
		t.Errorf("Unexpected answer: %v", actual)
	}
}

func TestWalk_Flatten(t *testing.T) {
	type Nested struct {
		S string
	}
	s := struct {
		A string
		B int
		C bool
		D Nested `structs:",flatten"`
		E int
	}{
		A: "a-value",
		B: 2,
		C: true,
		D: Nested{
			S: "s-value",
		},
		E: 999,
	}
	expected := `
A string a-value
B int 2
C bool true
S string s-value
E int 999`

	var actual string
	err := Walk(s, func(field reflect.StructField, value reflect.Value) error {
		actual = fmt.Sprintf("%s\n%s %s %v", actual, field.Name, field.Type, value.Interface())
		return nil
	})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if actual != expected {
		t.Errorf("Unexpected answer: %v", actual)
	}
}

func TestWalk_StructWithAnyField(t *testing.T) {
	type TestStruct struct {
		Foo string
	}
	type TestStructWithAnyField struct {
		Foo    any
		Nested TestStruct
	}
	var s TestStructWithAnyField

	var actual string
	err := Walk(s, func(field reflect.StructField, value reflect.Value) error {
		actual = fmt.Sprintf("%s\n%s %s %v", actual, field.Name, field.Type, value.Interface())
		return nil
	})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := `
Foo interface {} <nil>
Nested structs.TestStruct {}
Foo string `
	if actual != expected {
		t.Errorf("Unexpected answer: %v", actual)
	}
}

func TestWalk_StructWithNestedAnyField(t *testing.T) {
	type TestStructWithNestedAnyField struct {
		Foo    string
		Nested any
	}
	var s TestStructWithNestedAnyField

	var actual string
	err := Walk(s, func(field reflect.StructField, value reflect.Value) error {
		actual = fmt.Sprintf("%s\n%s %s %v", actual, field.Name, field.Type, value.Interface())
		return nil
	})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := `
Foo string 
Nested interface {} <nil>`
	if actual != expected {
		t.Errorf("Unexpected answer: %v", actual)
	}
}

func TestWalk_StructWithSliceField(t *testing.T) {
}

package jsonx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testJson = `{"foo":1,"bar":"baz","buz":true,"cuz":{"muz":"fuz"},"fud":["mud",true,10.0]}`

func TestQueryToType(t *testing.T) {

	intVal, err := QueryToType[int]("foo", testJson)
	assert.Nil(t, err)
	assert.Equal(t, 1, *intVal)

	strVal, err := QueryToType[string]("bar", testJson)
	assert.Nil(t, err)
	assert.Equal(t, "baz", *strVal)

	boolVal, err := QueryToType[bool]("buz", testJson)
	assert.Nil(t, err)
	assert.Equal(t, true, *boolVal)

	mapVal, err := QueryToType[map[string]any]("cuz", testJson)
	expectedMapVal := map[string]any{"muz": "fuz"}
	assert.Nil(t, err)
	assert.Equal(t, expectedMapVal, *mapVal)

	type Cuz struct {
		Muz string `json:"cuz"`
	}
	cuz, err := QueryToType[Cuz]("cuz", testJson)
	expectedCuz := Cuz{Muz: "fuz"}
	assert.Nil(t, err)
	assert.Equal(t, expectedCuz, *cuz)

	type Fud []any
	fud, err := QueryToType[Fud]("fud", testJson)
	expectedFud := Fud{"mud", true, 10.0}
	assert.Nil(t, err)
	assert.Equal(t, expectedFud, *fud)
}

func TestQueryToJson(t *testing.T) {

	jsonVal, err := QueryToJson("foo", testJson)
	assert.Nil(t, err)
	assert.Equal(t, `1`, jsonVal)

	jsonVal, err = QueryToJson("bar", testJson)
	assert.Nil(t, err)
	assert.Equal(t, `"baz"`, jsonVal)

	jsonVal, err = QueryToJson("buz", testJson)
	assert.Nil(t, err)
	assert.Equal(t, `true`, jsonVal)

	jsonVal, err = QueryToJson("cuz", testJson)
	assert.Nil(t, err)
	assert.Equal(t, `{"muz":"fuz"}`, jsonVal)

	jsonVal, err = QueryToJson("fud", testJson)
	assert.Nil(t, err)
	assert.Equal(t, `["mud",true,10]`, jsonVal)
}

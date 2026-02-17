package jsonx

import (
	"encoding/json"

	"github.com/elgs/gojq"
	"github.com/tartale/go/pkg/generics"
)

// QueryToType runs a gojq query path against inputJson and casts the result to *T.
//
// Example:
//
//	var name *string
//	name, err := jsonx.QueryToType[string](".name", `{"name":"alice"}`)
//	_ = name
//	_ = err
func QueryToType[T any](path string, inputJson string) (*T, error) {
	parser, err := gojq.NewStringQuery(inputJson)
	if err != nil {
		return nil, err
	}
	obj, err := parser.Query(path)
	if err != nil {
		return nil, err
	}
	val, err := generics.CastTo[T](obj)
	if err != nil {
		return nil, err
	}

	return val, nil
}

// QueryObjToType marshals input to JSON and then runs a gojq query path,
// returning the result cast to *T.
func QueryObjToType[T any](path string, input any) (*T, error) {
	inputJson, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	return QueryToType[T](path, string(inputJson))
}

// QueryToJson runs a gojq query path against inputJson and returns the
// resulting value marshaled back to JSON.
//
// Example:
//
//	out, err := jsonx.QueryToJson(".age", `{"age":30}`)
//	_ = out
//	_ = err
func QueryToJson(path string, inputJson string) (string, error) {
	parser, err := gojq.NewStringQuery(inputJson)
	if err != nil {
		return "", err
	}
	obj, err := parser.Query(path)
	if err != nil {
		return "", err
	}
	objBytes, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}

	return string(objBytes), nil
}

// QueryObjToJson marshals input to JSON and runs a gojq query path, returning
// the resulting value marshaled as JSON.
func QueryObjToJson(path string, input any) (string, error) {
	inputJson, err := json.Marshal(input)
	if err != nil {
		return "", err
	}

	return QueryToJson(path, string(inputJson))
}

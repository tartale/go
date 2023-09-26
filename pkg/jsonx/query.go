package jsonx

import (
	"encoding/json"

	"github.com/elgs/gojq"
	"github.com/tartale/go/pkg/generics"
)

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

func QueryObjToType[T any](path string, input any) (*T, error) {

	inputJson, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	return QueryToType[T](path, string(inputJson))
}

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

func QueryObjToJson(path string, input any) (string, error) {

	inputJson, err := json.Marshal(input)
	if err != nil {
		return "", err
	}

	return QueryToJson(path, string(inputJson))
}

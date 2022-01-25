package pkg

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRenderFromJSON_Simple(t *testing.T) {
	inputTemplate := `{{ .field | default "defaultValue" }}`
	inputJSON := `{ "field": "foo" }`
	output := bytes.NewBufferString("")
	DefaultOutputWriter = output

	err := RenderTextFromJSON(inputTemplate, inputJSON, "")
	assert.Nil(t, err)
	assert.Equal(t, "foo", output.String())
}

func TestRenderFromJSON_UseDefault(t *testing.T) {
	inputTemplate := `{{ .field | default "defaultValue" }}`
	inputJSON := `{ "foo": "bar" }`
	output := bytes.NewBufferString("")
	DefaultOutputWriter = output

	err := RenderTextFromJSON(inputTemplate, inputJSON, "")
	assert.Nil(t, err)
	assert.Equal(t, "defaultValue", output.String())
}

func TestRenderFromJSON_NoDefault(t *testing.T) {
	inputTemplate := `{{ .field }}`
	inputJSON := `{ "foo": "bar" }`
	output := bytes.NewBufferString("")
	DefaultOutputWriter = output

	err := RenderTextFromJSON(inputTemplate, inputJSON, "")
	assert.Nil(t, err)
	assert.Equal(t, "<no value>", output.String())
}

func TestRenderFromJSON_IllegalTemplate(t *testing.T) {
	inputTemplate := `{{ .field }`
	inputJSON := `{ "foo": "bar" }`
	output := bytes.NewBufferString("")
	DefaultOutputWriter = output

	err := RenderTextFromJSON(inputTemplate, inputJSON, "")
	assert.NotNil(t, err)
}

func TestRenderFromJSON_IllegalJSON(t *testing.T) {
	inputTemplate := `{{ .field }}`
	inputJSON := `{ "foo: "bar" }`
	output := bytes.NewBufferString("")
	DefaultOutputWriter = output

	err := RenderTextFromJSON(inputTemplate, inputJSON, "")
	assert.NotNil(t, err)
}

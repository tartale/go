package pkg

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
)

var DefaultOutputWriter io.Writer = os.Stdout

func RenderTextFromJSON(inputTemplate, inputData, output string) error {

	inputTemplate, err := getInputTemplate(inputTemplate)
	if err != nil {
		return err
	}
	jsonData, err := getJSON(inputData)
	if err != nil {
		return err
	}
	outputWriter, err := getOutputWriter(output)
	if err != nil {
		return err
	}

	return renderText(inputTemplate, jsonData, outputWriter)
}

func getInputTemplate(input string) (string, error) {
	inputFile, err := os.Open(input)
	if err != nil {
		// assume the input is a string
		return input, nil
	}
	defer inputFile.Close()
	inputFileContents, err := ioutil.ReadAll(inputFile)
	if err != nil {
		return "", err
	}

	return string(inputFileContents), nil
}

func getJSON(input string) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	if strings.HasSuffix(input, ".json") {
		// assume the input is a file, and read it
		inputFile, err := os.Open(input)
		if err != nil {
			return nil, err
		}
		defer inputFile.Close()
		inputFileContents, err := ioutil.ReadAll(inputFile)
		if err != nil {
			return nil, err
		}
		input = string(inputFileContents)
	}
	err := json.Unmarshal([]byte(input), &result)

	return result, err
}

func getOutputWriter(output string) (io.Writer, error) {
	var (
		outputWriter io.Writer
		err          error
	)

	if output != "" {
		outputWriter, err = os.Open(output)
		if err != nil {
			return nil, err
		}
	} else {
		outputWriter = DefaultOutputWriter
	}

	return outputWriter, nil
}

func renderText(inputTemplate string, inputData map[string]interface{}, outputWriter io.Writer) error {
	var templateFuncs = make(template.FuncMap)
	for k, v := range sprig.TxtFuncMap() {
		templateFuncs[k] = v
	}
	tmpl, err := template.New(inputTemplate).Funcs(templateFuncs).Parse(inputTemplate)
	if err != nil {
		return err
	}
	err = tmpl.Execute(outputWriter, inputData)
	if err != nil {
		return err
	}

	return nil
}

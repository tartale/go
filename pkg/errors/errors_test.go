package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrors_Combine_NoErrors(t *testing.T) {
	var errs Errors

	assert.Nil(t, errs.Combine("", "; "))
}

func TestErrors_Combine_OneError(t *testing.T) {
	var errs Errors

	errs = append(errs, errors.New("WTF"))

	assert.Equal(t, errors.New("WTF"), errs.Combine("", "; "))
}

func TestErrors_Combine_MultipleErrors(t *testing.T) {
	var errs Errors

	errs = append(errs, errors.New("WTF"))
	errs = append(errs, errors.New("How did this happen?"))
	errs = append(errs, errors.New("Worked on my machine"))

	assert.Equal(t, errors.New("WTF; How did this happen?; Worked on my machine"), errs.Combine("", "; "))
}

func TestErrors_Combine_PrefixMessage(t *testing.T) {
	var errs Errors

	errs = append(errs, errors.New("WTF"))
	errs = append(errs, errors.New("How did this happen?"))
	errs = append(errs, errors.New("Worked on my machine"))

	assert.Equal(t, errors.New("error while doing stuff: WTF; How did this happen?; Worked on my machine"), errs.Combine("error while doing stuff", "; "))
}

func TestErrors_Combine_LineSeparator(t *testing.T) {
	var errs Errors

	errs = append(errs, errors.New("WTF"))
	errs = append(errs, errors.New("How did this happen?"))
	errs = append(errs, errors.New("Worked on my machine"))

	assert.Equal(t, errors.New("WTF\nHow did this happen?\nWorked on my machine"), errs.Combine("", "\n"))
}

func TestErrors_Error(t *testing.T) {
	var errs Errors

	errs = append(errs, errors.New("WTF"))
	errs = append(errs, errors.New("How did this happen?"))
	errs = append(errs, errors.New("Worked on my machine"))

	assert.Equal(t, "WTF; How did this happen?; Worked on my machine", errs.Error())
}

func FunctionThatSucceeds(a int, b int) (int, error) {
	return a + b, nil
}

func FunctionThatErrors(a int, b int) (int, error) {
	return 0, errors.New("WTF")
}

func TestErrors_Try_SuccessfulFunction(t *testing.T) {
	var errs Errors

	a := 10
	b := 20
	errs.Try(func() error {
		_, err := FunctionThatSucceeds(a, b)

		return err
	})

	assert.Equal(t, 0, len(errs))
}

func TestErrors_Try_ErrorFunction(t *testing.T) {
	var errs Errors

	a := 10
	b := 20
	errs.Try(func() error {
		_, err := FunctionThatErrors(a, b)

		return err
	})

	assert.Equal(t, 1, len(errs))
}

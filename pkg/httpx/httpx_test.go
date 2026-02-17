package httpx

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tartale/go/pkg/errorz"
)

func TestGetResponseError_IncludesStatusAndBody(t *testing.T) {
	resp := &http.Response{
		Status:     "500 Internal Server Error",
		StatusCode: 500,
		Body:       io.NopCloser(strings.NewReader("some error body")),
	}

	err := GetResponseError(resp)

	if assert.Error(t, err) {
		assert.ErrorIs(t, err, errorz.ErrResponse)
		assert.Contains(t, err.Error(), "500 Internal Server Error")
		assert.Contains(t, err.Error(), "some error body")
	}
}


package httpx

import (
	"bufio"
	"fmt"
	"net/http"

	"github.com/tartale/go/pkg/errorz"
)

// GetResponseError builds a wrapped ErrResponse from the HTTP status and body
// of resp, suitable for returning from HTTP client helpers.
//
// Example:
//
//	resp, _ := http.DefaultClient.Do(req)
//	if resp.StatusCode >= 400 {
//		return httpx.GetResponseError(resp)
//	}
func GetResponseError(resp *http.Response) error {

	status := resp.Status
	r := bufio.NewReader(resp.Body)
	errorString := status
	bodyString, err := r.ReadString(byte(0))
	if err != nil {
		errorString = fmt.Sprintf("%s; %s", errorString, bodyString)
	}

	return fmt.Errorf("%w: %s", errorz.ErrResponse, errorString)
}

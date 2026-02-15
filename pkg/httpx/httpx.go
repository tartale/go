package httpx

import (
	"bufio"
	"fmt"
	"net/http"

	"github.com/tartale/go/pkg/errorz"
)

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

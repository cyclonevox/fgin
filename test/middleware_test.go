package test

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestMiddleware(t *testing.T) {
	setUpTestServer()

	t.Run("get", func(t *testing.T) {

		resp, err := http.Get(testUrl + "/v3/hello")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Println(string(body))
	})
}

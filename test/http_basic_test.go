package test

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestHTTPBasic(t *testing.T) {
	setUpTestServer()

	t.Run("get1", func(t *testing.T) {

		resp, err := http.Get(testUrl)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Println(string(body))
	})

	t.Run("get2", func(t *testing.T) {

		resp, err := http.Get(testUrl + "/hello")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Println(string(body))
	})

	t.Run("get3", func(t *testing.T) {

		resp, err := http.Get(testUrl + "/v2")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Println(string(body))
	})

	t.Run("get4", func(t *testing.T) {

		resp, err := http.Get(testUrl + "/v1/hello")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Println(string(body))
	})
}

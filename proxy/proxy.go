package proxy

import (
	"fmt"
	"io"
	"net/http"
)

func HttpCall(url string) error {

	response, err := http.Get(url)
	if err != nil {
		return err
	}
	byteBody, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(byteBody))
	return nil
}

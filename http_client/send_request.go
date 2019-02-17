package http_client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	ContentTypeJson = `application/json`
)

func SendPostRequest(url, data string) (err error) {
	resp, err := http.Post(url, ContentTypeJson, strings.NewReader(data))
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	return
}

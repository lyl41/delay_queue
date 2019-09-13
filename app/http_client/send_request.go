package http_client

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	ContentTypeJson   = `application/json`
	PostResultSuccess = `SUCCESS`
)

func SendPostRequest(url, data string) (err error) {
	resp, err := http.Post(url, ContentTypeJson, strings.NewReader(data))
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		fmt.Println("post err", err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if strings.ToUpper(string(body)) != PostResultSuccess {
		err = errors.New("post result is not SUCCESS, try to rePost")
		fmt.Println("post result is not SUCCESS, ", string(body))
	}
	return
}

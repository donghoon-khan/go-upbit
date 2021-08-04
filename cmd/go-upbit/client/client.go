package client

import (
	"io/ioutil"
	"net/http"
)

const (
	apiUrl     = "https://api.upbit.com"
	apiVersion = "v1"
)

func getResponse(client *http.Client, req *http.Request) ([]byte, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

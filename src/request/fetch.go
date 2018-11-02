package request

import (
	"net/http"
	"encoding/json"
	"bytes"
	"errors"
)

func FetchJson(req *http.Request, res interface{}) (*http.Response, error) {
	client := &http.Client{}

	resp, err := client.Do(req); if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if res != nil {
		err = json.NewDecoder(resp.Body).Decode(res); if err != nil {
			return nil, err
		}
	}
	return resp, err
}

func CreateFetchJson(url string, method string, token string, body interface{}) (*http.Request, error) {
	jsonData := []byte("")
	err := errors.New("")

	if body != nil {
		jsonData, err = json.Marshal(body); if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData)); if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", token)
	req.Header.Set("Accept", "application/json")
	if method != "GET" {
		req.Header.Set("Content-Type", "application/json")
	}
	return req, nil
}

func ExpectJson(url string, method string, token string, body interface{}, res interface{}) (*http.Response, error) {
	req, err := CreateFetchJson(url, method, token, body); if err != nil {
		return nil, err
	}
	return FetchJson(req, res)
}

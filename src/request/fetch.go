package request

import (
	"net/http"
	"encoding/json"
	"bytes"
	"errors"
	"github.com/amstee/easy-cut/src/logger"
)


func IsValid(status int) bool {
	return status >= 200 && status < 300
}

func FetchJson(req *http.Request, res interface{}) (*http.Response, error) {
	client := &http.Client{}

	DisplayRequest(req)
	resp, err := client.Do(req); if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	DisplayResponse(resp)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		if res != nil {
			err = json.NewDecoder(resp.Body).Decode(res); if err != nil {
				return nil, err
			}
		}
	} else {
		logger.Info.Printf("Received an invalid status %d code for %s", resp.StatusCode, req.URL)
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

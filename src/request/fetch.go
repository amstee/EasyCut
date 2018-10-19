package request

import (
	"net/http"
	"encoding/json"
	"bytes"
	"errors"
)

func ExpectJson(url string, method string, token string, body interface{}, res interface{}) (*http.Response, error) {
	jsonData := []byte("")
	err := errors.New("")
	client := &http.Client{}

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
	resp, err := client.Do(req); if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(res); if err != nil {
		return nil, err
	}
	return resp, err
}
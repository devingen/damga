package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func sendRequest(apiKey, method, path string, body interface{}, responseBody interface{}) (interface{}, error) {

	requestBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	apiURL := "https://api.damga.devingen.io"
	request, err := http.NewRequest(method, apiURL+path, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("api-key", apiKey)

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&responseBody)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

func post(apiKey, path string, body interface{}, responseBody interface{}) (interface{}, error) {
	return sendRequest(apiKey, http.MethodPost, path, body, responseBody)
}

package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

func sendRequest(method, path string, body interface{}, responseBody interface{}) (interface{}, error) {

	requestBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	apiURL := os.Getenv("DAS_DAMGA_API_URL")
	if apiURL == "" {
		apiURL = "https://api.damga.devingen.io"
	}

	request, err := http.NewRequest(method, apiURL+path, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json")

	apiKey := os.Getenv("DAMGA_API_KEY")
	if apiKey != "" {
		request.Header.Add("api-key", apiKey)
	}

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

func post(path string, body interface{}, responseBody interface{}) (interface{}, error) {
	return sendRequest(http.MethodPost, path, body, responseBody)
}

func put(path string, body interface{}, responseBody interface{}) (interface{}, error) {
	return sendRequest(http.MethodPut, path, body, responseBody)
}

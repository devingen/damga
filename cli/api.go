package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func sendRequest(apiAddress, apiKey, method, path string, body interface{}, responseBody interface{}) (interface{}, error) {

	requestBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	apiURL := "https://api.damga.das.devingen.io"
	if apiAddress != "" {
		apiURL = apiAddress
	}
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

func post(apiAddress, apiKey, path string, body interface{}, responseBody interface{}) (interface{}, error) {
	return sendRequest(apiAddress, apiKey, http.MethodPost, path, body, responseBody)
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func QueryFlowise(query string) (string, error) {
	flowiseApi := os.Getenv("FLOWISE_API")
	flowiseKey := os.Getenv("FLOWISE_KEY")

	type FlowiseQuery struct {
		Question string `json:"question"`
	}

	body := &FlowiseQuery{
		Question: query,
	}

	payloadBuf := new(bytes.Buffer)
	err := json.NewEncoder(payloadBuf).Encode(body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	req, err := http.NewRequest("POST", flowiseApi, bytes.NewBuffer(payloadBuf.Bytes()))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", flowiseKey)

	resp, err := HTTPClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	respStr := strings.Trim(string(respBody), "\"")

	fmt.Println(respStr)

	return respStr, nil
}

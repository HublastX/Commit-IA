package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RequestPayload struct {
	Message string `json:"message"`
}

type ResponsePayload struct {
	Response string `json:"response"`
}

type CommitAnalyzerRequest struct {
	CodeChanges string `json:"code_changes"`
	Description string `json:"description"`
	Tag         string `json:"tag"`
	Language    string `json:"language"`
}

func SendCommitAnalysisRequest(url string, codeChanges, description, tag, language string) (*ResponsePayload, error) {
	payload := CommitAnalyzerRequest{
		CodeChanges: codeChanges,
		Description: description,
		Tag:         tag,
		Language:    language,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error serializing commit analysis payload: %v", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error making POST request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API response error: %s, body: %s", resp.Status, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	var responsePayload ResponsePayload
	err = json.Unmarshal(body, &responsePayload)
	if err != nil {
		return nil, fmt.Errorf("error deserializing response: %v", err)
	}

	return &responsePayload, nil
}

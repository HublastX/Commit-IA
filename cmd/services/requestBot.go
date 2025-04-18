package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	schemas "github.com/HublastX/Commit-IA/schema"
)

func SendCommitAnalysisRequest(url string, codeChanges, description, tag, language string) (*schemas.ResponsePayload, error) {
	payload := schemas.CommitAnalyzerRequest{
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

	var responsePayload schemas.ResponsePayload
	err = json.Unmarshal(body, &responsePayload)
	if err != nil {
		return nil, fmt.Errorf("error deserializing response: %v", err)
	}

	return &responsePayload, nil
}

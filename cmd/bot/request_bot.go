package bot

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

func SendMessageToBot(url string, message string) (*ResponsePayload, error) {

	payload := RequestPayload{
		Message: message,
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("erro ao serializar o payload: %v", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer a requisição POST: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("erro na resposta da API: %s, body: %s", resp.Status, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler a resposta: %v", err)
	}

	var responsePayload ResponsePayload
	err = json.Unmarshal(body, &responsePayload)
	if err != nil {
		return nil, fmt.Errorf("erro ao desserializar a resposta: %v", err)
	}

	return &responsePayload, nil
}

package ai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"go-echo-starter/internal/domain"
)

type Service struct {
	apiKey string
	client *http.Client
}

func NewService(apiKey string) *Service {
	return &Service{
		apiKey: apiKey,
		client: &http.Client{},
	}
}

func (s *Service) ParseMedicalText(ctx context.Context, input string) (*domain.AIResponse, error) {

	prompt := fmt.Sprintf(`
		Extract structured data from the following medical text.

		Return ONLY valid JSON.
		Do NOT include markdown formatting.
		Do NOT include backticks.
		Do NOT include explanations.

		Format:
		{
		"drugs": [{"name": "", "dosage": "", "frequency": "", "duration": ""}],
		"lab_tests": [],
		"notes": ""
		}

		Text:
		"%s"
		`, input)

	reqBody := map[string]interface{}{
		"model": "openrouter/free", // free model
		"messages": []map[string]string{
			{"role": "user", "content": prompt},
		},
	}

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		"https://openrouter.ai/api/v1/chat/completions",
		bytes.NewBuffer(bodyBytes),
	)
	if err != nil {
		return nil, err
	}

	// Required headers
	req.Header.Set("Authorization", "Bearer "+s.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("HTTP-Referer", "http://localhost")
	req.Header.Set("X-Title", "medical-parser-app")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Read raw response (important for debugging)
	body, _ := io.ReadAll(resp.Body)

	var raw map[string]interface{}
	if err := json.Unmarshal(body, &raw); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w\nRaw: %s", err, string(body))
	}

	// Handle API errors safely
	if errObj, ok := raw["error"]; ok {
		return nil, fmt.Errorf("openrouter error: %v", errObj)
	}

	// Safe extraction
	choices, ok := raw["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		return nil, fmt.Errorf("invalid response: %s", string(body))
	}

	choice, ok := choices[0].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid choice format")
	}

	message, ok := choice["message"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid message format")
	}

	content, ok := message["content"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid content format")
	}

	// Parse AI JSON
	var result domain.AIResponse
	if err := json.Unmarshal([]byte(content), &result); err != nil {
		return nil, fmt.Errorf("failed to parse AI JSON: %w\nRaw content: %s", err, content)
	}

	return &result, nil
}

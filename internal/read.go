package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func ReadFromFile(filename string) (ApiResponse, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return ApiResponse{}, fmt.Errorf("failed to read file: %w", err)
	}
	var g ApiResponse
	if err := json.Unmarshal(data, &g); err != nil {
		return ApiResponse{}, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	return g, nil
}

func FetchData() (ApiResponse, error) {
	url := "https://harita.iski.gov.tr/data/mahallelerKesinti.geojson"
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return ApiResponse{}, fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ApiResponse{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return ApiResponse{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var g ApiResponse
	if err := json.Unmarshal(data, &g); err != nil {
		return ApiResponse{}, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return g, nil
}

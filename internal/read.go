package internal

import (
	"encoding/json"
	"fmt"
	"os"
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

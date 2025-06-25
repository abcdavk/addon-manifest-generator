package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type MinecraftServerLatest struct {
	Version string `json:"version"` // Harus huruf besar & tag JSON benar
}

func conn(moduleName string) (version string) {
	url := fmt.Sprintf("https://registry.npmjs.org/%s/latest", moduleName)

	resp, err := http.Get(url)
	if err != nil {
		return "2.0.0"
	}
	defer resp.Body.Close()

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("could not read response body: %s\n", err)
	}

	var result MinecraftServerLatest
	if err := json.Unmarshal(resBody, &result); err != nil {
		fmt.Printf("failed to decode JSON: %v\n", err)
	}

	return result.Version
}

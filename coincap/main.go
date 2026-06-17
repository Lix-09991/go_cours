package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
)

type Asset struct {
	Name     string `json:"name"`
	PriceUsd string `json:"priceUsd"`
}

type assetsResponse struct {
	Data []Asset `json:"data"`
}

func main() {
	apiKey := "f4ba20a6b1d84b6e1d0dc2af80781db8dd600e7e883cd4e0893a41994dea2c09"

	req, err := http.NewRequest("GET", "https://api.coincap.io/v2/assets", nil)
	if err != nil {
		fmt.Println("error creating request:", err)
		os.Exit(1)
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error sending request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading response:", err)
		os.Exit(1)
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("api error: status=%d body=%s\n", resp.StatusCode, string(body))
		os.Exit(1)
	}

	var parsed assetsResponse
	if err := json.Unmarshal(body, &parsed); err != nil {
		fmt.Println("error parsing json:", err)
		os.Exit(1)
	}

	type outAsset struct {
		Name     string `json:"name"`
		PriceUsd string `json:"priceUsd"`
	}

	target := map[string]bool{
		"Bitcoin":  true,
		"Ethereum": true,
	}

	var result []outAsset
	for _, a := range parsed.Data {
		if target[a.Name] {
			result = append(result, outAsset{
				Name:     a.Name,
				PriceUsd: a.PriceUsd,
			})
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})
}

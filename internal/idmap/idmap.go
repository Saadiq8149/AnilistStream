package idmap

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const baseURL = "https://api.ids.moe"

type IDMapService struct{}

func NewIDMapService() *IDMapService {
	return &IDMapService{}
}

func (s *IDMapService) GetIDMap(id string, provider string) (map[string]string, error) {
	url := fmt.Sprintf("%s/ids/%s?platform=%s", baseURL, id, provider)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("IDS_MOE_API_KEY"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ids.moe returned status %d", resp.StatusCode)
	}

	var raw map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return nil, err
	}

	result := make(map[string]string)

	for k, v := range raw {
		if v == nil {
			continue
		}

		switch val := v.(type) {
		case string:
			result[k] = val
		case float64:
			result[k] = fmt.Sprintf("%.0f", val)
		}
	}

	return result, nil
}

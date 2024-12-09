package gobreach

import (
	"io"
	"fmt"
	"net/http"
	"encoding/json"
)

type BreachEntry struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Sha1     string `json:"sha1"`
	Hash     string `json:"hash"`
	Sources  string `json:"sources"`
}

type BreachDirectoryResponse struct {
	Found  int           `json:"found"`
	Result []BreachEntry `json:"result"`
}

type BreachDirectoryClient struct {
	APIKey string
}

func NewBreachDirectoryClient(apiKey string) (*BreachDirectoryClient, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("API key cannot be empty")
	}

	return &BreachDirectoryClient{
		APIKey: apiKey,
	}, nil
}

func (client *BreachDirectoryClient) SearchEmail(email string) (*BreachDirectoryResponse, error) {
	if email == "" {
		return nil, fmt.Errorf("email cannot be empty")
	}

	url := "https://breachdirectory.p.rapidapi.com/?func=auto&term=" + email

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Add("x-rapidapi-key", client.APIKey)
	req.Header.Add("x-rapidapi-host", "breachdirectory.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error executing request: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var response BreachDirectoryResponse

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON response: %v", err)
	}

	return &response, nil
}

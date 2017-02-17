package now

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Client is the basic http client for making http calls and processing responses
type Client struct {
	BaseURL string
	Token   string
}

// ProcessResponse parses the response body into specified result
func (c *Client) ProcessResponse(resp *http.Response, result interface{}) error {
	if resp.StatusCode != 200 {
		errorResponse := ErrorResponse{}
		err := json.NewDecoder(resp.Body).Decode(&errorResponse)
		if err != nil {
			log.Error(err)
			return err
		}
		return errors.New(errorResponse.Err.Message)
	}
	if result != nil {
		return json.NewDecoder(resp.Body).Decode(&result)
	}
	return nil
}

// Call makes a request with speficied method, path, data and result interface
func (c *Client) Call(method, path string, data, result interface{}) error {
	b := new(bytes.Buffer)
	if data != nil {
		json.NewEncoder(b).Encode(data)
	}
	req, err := http.NewRequest(method, c.BaseURL+path, b)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Token)
	if err != nil {
		log.Error(err)
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Error(err)
		return err
	}
	return c.ProcessResponse(resp, result)
}

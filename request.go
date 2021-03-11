package picqer

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// NewRequest creates a new HTTP request with an optional JSON body
func (c *Client) NewRequest(method, path string, body interface{}) (*http.Response, error) {
	var j []byte = nil

	if body != nil {
		var err error

		j, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	client := &http.Client{}

	req, err := http.NewRequest(method, c.BaseURL+"/"+path, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.Token, "")

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "go-picqer (github.com/anthonycook/go-picqer - hi@acook.me)")

	res, _ := client.Do(req)

	return res, nil
}

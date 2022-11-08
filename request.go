package picqer

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
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

	req, err := http.NewRequest(method, "https://"+c.BaseURL+"/api/v1/"+path, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.Token, "")

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "go-picqer (github.com/anthonycook/picqer - hi@acook.me)")

	res, _ := client.Do(req)

	// Check rate limit
	rateLimit := res.Header.Get("X-RateLimit-Limit")
	rateLimitRemaining := res.Header.Get("X-RateLimit-Remaining")

	if rateLimitRemaining == rateLimit {
		time.Sleep(60 * time.Second)
	}

	return res, nil
}

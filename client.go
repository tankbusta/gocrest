package gocrest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Global Client
var c Client = Client{&http.Client{}}

type Client struct {
	*http.Client
}

type ClientError struct {
	Message       string `json:"message"`
	IsLocalized   bool   `json:"isLocalized"`
	Key           string `json:"noSuchAlliance"`
	ExceptionType string `json:"ForbiddenError"`
}

func (s *Client) doGetRequest(ep string, unmarshal_to interface{}) error {
	fmt.Sprintf("EP: %s\n", ep)

	req, err := http.NewRequest("GET", ep, nil)
	if err != nil {
		return err
	}

	req.Header.Add("User-Agent", UA_STRING)

	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	// Set up the decoder to read the response
	dec := json.NewDecoder(resp.Body)

	// If the status was not successful...
	if resp.StatusCode != http.StatusOK {
		e := ClientError{}
		if err = dec.Decode(e); err != nil {
			return err
		}
		unmarshal_to = &e
		return err
	}

	if err = dec.Decode(unmarshal_to); err != nil {
		return err
	}
	return nil
}

package automower

import (
	"encoding/json"
	"net/http"
)

type Client struct {
	client      *http.Client
	accessToken string
	provider    string
}

func NewClientWithUserAndPassword(user, password string) (*Client, error) {
	c := Client{client: http.DefaultClient}
	err := c.authenticate(user, password)
	return &c, err
}

func (c *Client) do(req *http.Request) (*http.Response, error) {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.accessToken)
	req.Header.Set("Authorization-Provider", c.provider)
	return c.client.Do(req)
}

func (c *Client) doJSON(req *http.Request, v interface{}) error {
	res, err := c.do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(v)

	if err != nil {
		return err
	}
	return nil
}

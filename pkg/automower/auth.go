package automower

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func (c *Client) authenticate(user, password string) error {

	u := loginRequest{Data: loginRequestBody{Attributes: loginRequestAttributes{Username: user, Password: password}, Type: "token"}}
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(u)
	if err != nil {
		return err
	}

	res, err := http.Post(authUrlToken, "application/json", b)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return errors.New("Authentication failed: " + strconv.Itoa(res.StatusCode))
	}

	var m message
	err = json.NewDecoder(res.Body).Decode(&m)
	if err != nil {
		return err
	}

	c.provider = m.Data.Attributes["provider"].(string)
	c.accessToken = m.Data.ID
	return nil
}

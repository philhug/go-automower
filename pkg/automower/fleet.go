package automower

import (
	"io/ioutil"
	"log"
	"net/http"
)

func (c *Client) getAssets() error {
	req, err := http.NewRequest("GET", fleetUrlAssets+"?query=&sort=ALPHABETICALLY&desc=true&offset=0&limit=20", nil)
	if err != nil {
		return err
	}
	res, err := c.do(req)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	log.Println(string(body))
	return nil
}

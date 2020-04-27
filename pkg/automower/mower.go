package automower

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func (c *Client) getMowers() ([]Mower, error) {
	var mowers []Mower

	req, err := http.NewRequest("GET", trackUrlMowers, nil)
	if err != nil {
		return nil, err
	}

	err = c.doJSON(req, &mowers)
	if err != nil {
		return nil, err
	}
	return mowers, nil
}

func (c *Client) getMowerStatus(id string) (*MowerStatus, error) {
	mow := MowerStatus{}

	req, err := http.NewRequest("GET", trackUrlMowers+"/"+id+"/status", nil)
	if err != nil {
		return nil, err
	}
	err = c.doJSON(req, &mow)
	if err != nil {
		return nil, err
	}
	return &mow, nil
}

func (c *Client) controlMower(id string, action, duration string) error {
	ar := actionRequest{action, duration}

	b, err := json.Marshal(ar)
	if err != nil {
		return err
	}
	log.Println(string(b))
	rb := bytes.NewBuffer(b)

	req, err := http.NewRequest("POST", trackUrlMowers+"/"+id+"/control", rb)
	if err != nil {
		return err
	}
	req.Header.Set("Content-type", "application/json")

	res, err := c.do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	log.Println(string(body))
	return nil
}

func (c *Client) getMowerGeoFenceText(id string) error {
	req, err := http.NewRequest("GET", trackUrlMowers+"/"+id+"/geofence", nil)
	if err != nil {
		return err
	}
	res, err := c.do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	log.Println(string(body))
	return nil
}

func (c *Client) getMowerGeoFence(id string) (*mowerGeoFence, error) {
	geofence := mowerGeoFence{}

	req, err := http.NewRequest("GET", trackUrlMowers+"/"+id+"/geofence", nil)
	if err != nil {
		return nil, err
	}
	err = c.doJSON(req, &geofence)
	if err != nil {
		return nil, err
	}
	return &geofence, nil
}

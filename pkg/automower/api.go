package automower

import (
	"strconv"
	"time"
)

func (c *Client) Status(m *Mower) (*MowerStatus, error) {
	return c.getMowerStatus(m.ID)
}

func (c *Client) StartWithTimer(m *Mower) error {
	return c.controlMower(m.ID, mowerActionStart, "timer")
}

func (c *Client) StartWithDuration(m *Mower, d time.Duration) error {
	s := strconv.Itoa(int(d.Minutes()))
	return c.controlMower(m.ID, mowerActionStart, s)
}

func (c *Client) StopWithTimer(m *Mower) error {
	return c.controlMower(m.ID, mowerActionStop, "timer")
}

func (c *Client) Mowers() ([]Mower, error) {
	return c.getMowers()
}

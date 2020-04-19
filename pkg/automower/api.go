package automower

func (c *Client) Status(m *Mower) (*MowerStatus, error) {
	return c.getMowerStatus(m.ID)
}

func (c *Client) StartWithTimer(m *Mower) error {
	return c.controlMower(m.ID, mowerActionStart)
}

func (c *Client) StopWithTimer(m *Mower) error {
	return c.controlMower(m.ID, mowerActionStop)
}

func (c *Client) Mowers() ([]Mower, error) {
	return c.getMowers()
}

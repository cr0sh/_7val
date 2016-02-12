package _7val

import "net"

// Client is a struct for emulating MCPE client.
type Client struct {
	Username string
	RHost    string
	conn     net.Conn
	raknet   *RaknetSession
}

// NewClient creates MCPE client, initializes, and returns it.
func NewClient(username string, rhost string) (c *Client, err error) {
	c = new(Client)
	c.Username, c.RHost = username, rhost
	err = c.Init()
	return
}

// Init sets valid initial value for struct.
func (c *Client) Init() (err error) {
	c.raknet = new(RaknetSession)
	c.raknet.Init()
	c.conn, err = net.Dial("udp4", c.RHost)
	return
}

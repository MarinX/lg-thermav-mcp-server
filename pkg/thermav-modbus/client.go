package thermavmodbus

import (
	"fmt"

	"github.com/simonvetter/modbus"
)

type Client struct {
	modbus.ClientConfiguration
	cli *modbus.ModbusClient
}

// New creates a new Modbus client with the given endpoint and options.
func New(endpoint string, opts ...Option) *Client {
	client := &Client{}
	client.URL = endpoint

	for i := range opts {
		opts[i](client)
	}
	return client
}

// Open initializes and opens the Modbus client connection.
func (c *Client) Open() error {
	cl, err := modbus.NewClient(&modbus.ClientConfiguration{
		URL:      c.URL,
		Timeout:  c.Timeout,
		StopBits: c.StopBits,
		Parity:   c.Parity,
		Speed:    c.Speed,
		Logger:   c.Logger,
	})
	if err != nil {
		return err
	}
	c.cli = cl
	return c.cli.Open()
}

// Close terminates the Modbus client connection.
func (c *Client) Close() error {
	if c.cli == nil {
		return fmt.Errorf("modbus client not initialized")
	}
	return c.cli.Close()
}

package thermavmodbus

import (
	"log"
	"time"
)

type Option func(c *Client)

func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.Timeout = timeout
	}
}

func WithParity(parity uint) Option {
	return func(c *Client) {
		c.Parity = parity
	}
}

func WithStopBits(bits uint) Option {
	return func(c *Client) {
		c.StopBits = bits
	}
}

func WithSpeed(speed uint) Option {
	return func(c *Client) {
		c.Speed = speed
	}
}

func WithLogger(log *log.Logger) Option {
	return func(c *Client) {
		c.Logger = log
	}
}

package emailc

import (
	"gopkg.in/gomail.v2"
)

type Client struct {
	config Config
	kernel *gomail.Dialer
}

func NewClient(options ...Option) (*Client, error) {
	return NewClientWithConfig(newConfig(options...))
}

func NewClientWithConfig(config Config) (*Client, error) {
	dialer := gomail.NewDialer(config.Host, config.Port, config.Username, config.Password)
	closer, err := dialer.Dial()
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = closer.Close()
	}()
	return &Client{kernel: dialer, config: config}, nil
}

func (client *Client) Kernel() *gomail.Dialer {
	return client.kernel
}

func (client *Client) Send(message *Message) error {
	if message.Sender == "" {
		message.Sender = client.config.Username
	}
	msg := message.toMessage()
	return client.kernel.DialAndSend(msg)
}

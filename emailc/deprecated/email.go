package deprecated

import (
	"fmt"
	"net/smtp"
)

const (
	HostQQMail = "smtp.qq.com"
	PortQQMail = "25"
)

type Client interface {
	Send(msg *Message) error
}

type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
	// default QQMail host
	Host string `json:"host"`
	// default QQMail port
	Port string `json:"port"`
	// usually an empty string
	Identity string `json:"identity"`
}

func NewClient(config Config) Client {
	handleConfig(&config)
	auth := smtp.PlainAuth(config.Identity, config.Username, config.Password, config.Host)
	return &client{
		config: config,
		auth:   auth,
	}
}

type client struct {
	config Config
	auth   smtp.Auth
}

func (c *client) Send(message *Message) error {
	msg, err := handleMessage(c, message)
	if err != nil {
		return err
	}
	addr := fmt.Sprintf("%s:%s", c.config.Host, c.config.Port)
	return smtp.SendMail(addr, c.auth, c.config.Username, message.To, msg)
}

func handleConfig(config *Config) {
	if config.Host == "" {
		config.Host = HostQQMail
	}
	if config.Port == "" {
		config.Port = PortQQMail
	}
}

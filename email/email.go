package email

import (
	"fmt"
	"github.com/whereabouts/sdk-go/utils/mapper"
	"github.com/whereabouts/sdk-go/utils/timer"
	"net/smtp"
	"strings"
)

const (
	contentTypeDefault = "text/plain;charset=UTF-8"
	contentTypeHtml    = "text/html;charset=UTF-8"
	HostQQMail         = "smtp.qq.com"
	PortQQMail         = "25"
)

type Mail struct {
	username string
	password string
	host     string
	port     string
	auth     smtp.Auth
	message  *Message
}

type Message struct {
	From string `json:"From"`
	To   string `json:"To"`
	// 抄送
	CC string `json:"Cc"`
	// 暗抄送
	BCC         string `json:"Bcc"`
	Subject     string `json:"Subject"`
	ContentType string `json:"Content-Type"`
	Date        string `json:"Date"`
}

// Auth authenticates the identity
// In particular, the password is the authorization code, not your email password
func Auth(username, password, host, port string) *Mail {
	return &Mail{
		username: username,
		password: password,
		host:     host,
		port:     port,
		auth:     smtp.PlainAuth("", username, password, host),
		message:  &Message{From: username},
	}
}

func (mail *Mail) SetFrom(from string) *Mail {
	mail.message.From = from
	return mail
}

func (mail *Mail) SetCC(CC []string) *Mail {
	mail.message.CC = strings.Join(CC, ";")
	return mail
}

func (mail *Mail) SetBCC(BCC []string) *Mail {
	mail.message.BCC = strings.Join(BCC, ";")
	return mail
}

func (mail *Mail) SetSubject(subject string) *Mail {
	mail.message.Subject = subject
	return mail
}

// Plain send email in a plain format
func (mail *Mail) Plain(to []string, msg string) error {
	mail.message.ContentType = contentTypeDefault
	return mail.send(to, msg)
}

// Html send email in a html format
func (mail *Mail) Html(to []string, msg string) error {
	mail.message.ContentType = contentTypeHtml
	return mail.send(to, msg)
}

func (mail *Mail) send(to []string, msg string) error {
	out, err := mail.handleMessage(to, msg)
	if err != nil {
		return err
	}
	err = smtp.SendMail(fmt.Sprintf("%s:%s", mail.host, mail.port), mail.auth, mail.username, to, []byte(out))
	if err != nil {
		return err
	}
	return nil
}

func (mail *Mail) handleMessage(to []string, msg string) (s string, err error) {
	mail.message.To = strings.Join(to, ";")
	mail.message.Date = timer.Now()
	msgMap := make(map[string]interface{})
	msgMap, err = mapper.Struct2Map(mail.message)
	if err != nil {
		return "", err
	}
	for key, value := range msgMap {
		s = fmt.Sprintf("%s%s:%s\r\n", s, key, value)
	}
	s = fmt.Sprintf("%s\r\n%s", s, msg)
	return s, nil
}

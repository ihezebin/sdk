package email

import (
	"errors"
	"fmt"
	"github.com/whereabouts/sdk-go/utils/mapper"
	"github.com/whereabouts/sdk-go/utils/timer"
	"strings"
)

const (
	ContentTypeText = "text/plain;charset=UTF-8"
	ContentTypeHtml = "text/html;charset=UTF-8"
)

type Message struct {
	// From sender, default "whereabouts.icu"
	From string `json:"From"`
	// To receivers, can not be nil
	To []string `json:"-"`
	// CC send a duplicate to
	CC []string `json:"-"`
	// BCC blind carbon copy
	BCC []string `json:"-"`
	// Subject email title, default "test email"
	Subject string `json:"Subject"`
	// ContentType email type：text or html
	ContentType string `json:"Content-Type"`
	// Date send time, default now
	Date string `json:"Date"`
	// Body email content, default "hello world！this is a test mail！"
	Body string `json:"-"`
}

var defaultMessage = Message{
	From:        "whereabouts.icu",
	ContentType: ContentTypeHtml,
	Date:        timer.Now(),
	Subject:     "test email",
	Body:        "hello world！this is a test mail！",
}

func NewMessage() *Message {
	return &defaultMessage
}

func handleMessage(client *client, message *Message) ([]byte, error) {
	if len(message.To) == 0 {
		return nil, errors.New("receiver has to be at least one person, use SetReceiver")
	}

	message.From = fmt.Sprintf("%s<%s>", message.From, client.config.Username)

	msgM, err := mapper.Struct2Map(message)
	if err != nil {
		return nil, err
	}
	msgM["To"] = strings.Join(message.To, ";")
	msgM["Cc"] = strings.Join(message.CC, ";")
	msgM["Bcc"] = strings.Join(message.BCC, ";")

	var msg string
	for key, value := range msgM {
		msg = fmt.Sprintf("%s%s:%s\r\n", msg, key, value)
	}
	msg = fmt.Sprintf("%s\r\n%s", msg, message.Body)

	return []byte(msg), nil
}

func (message *Message) SetDate(date string) *Message {
	message.Date = date
	return message
}

func (message *Message) SetBody(body string) *Message {
	message.Body = body
	return message
}

// Deprecated: Use SetBody to instead.
func (message *Message) SetContent(content string) *Message {
	message.Body = content
	return message
}

func (message *Message) SetContentType(contentType string) *Message {
	message.ContentType = contentType
	return message
}

func (message *Message) SetSubject(subject string) *Message {
	message.Subject = subject
	return message
}

func (message *Message) SetReceiver(to ...string) *Message {
	message.To = to
	return message
}

func (message *Message) SetSender(from string) *Message {
	message.From = from
	return message
}

func (message *Message) SetCC(cc ...string) *Message {
	message.CC = cc
	return message
}

func (message *Message) SetBCC(bcc ...string) *Message {
	message.BCC = bcc
	return message
}

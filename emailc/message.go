package emailc

import (
	"gopkg.in/gomail.v2"
	"io"
	"time"
)

type Message struct {
	// Sender 发件者邮箱地址
	Sender string
	// Receiver 收件者邮箱地址
	Receiver []string
	// Title 邮件标题
	Title string
	// Text 邮件文本内容, 设置后 Html 失效
	Text string
	// 邮件HTML富文本内容, 设置后 Text 失效
	Html string
	// Date 发件时间
	Date time.Time
	// CC 抄送人邮箱地址
	CC []string
	// BCC 密送人邮箱地址
	BCC    []string
	Attach []Attach
}

type Attach struct {
	// Name 附件文件名称
	Name string
	// File 附件文件Reader
	File io.Reader
}

func NewAttach(name string, file io.Reader) Attach {
	return Attach{Name: name, File: file}
}

func NewMessage() *Message {
	return &Message{
		Date: time.Now(),
	}
}

func (msg *Message) WithSender(sender string) *Message {
	msg.Sender = sender
	return msg
}

func (msg *Message) WithReceiver(receiver ...string) *Message {
	msg.Receiver = receiver
	return msg
}

func (msg *Message) WithTitle(title string) *Message {
	msg.Title = title
	return msg
}

func (msg *Message) WithText(text string) *Message {
	msg.Text = text
	msg.Html = ""
	return msg
}

func (msg *Message) WithHtml(html string) *Message {
	msg.Html = html
	msg.Text = ""
	return msg
}

func (msg *Message) WithDate(date time.Time) *Message {
	msg.Date = date
	return msg
}

func (msg *Message) WithCC(cc ...string) *Message {
	msg.CC = cc
	return msg
}

func (msg *Message) WithBCC(bcc ...string) *Message {
	msg.BCC = bcc
	return msg
}

func (msg *Message) WithAttach(attach ...Attach) *Message {
	msg.Attach = attach
	return msg
}

func (msg *Message) toMessage() *gomail.Message {
	message := gomail.NewMessage()
	//message.Attach()
	message.SetHeader("From", msg.Sender)
	message.SetHeader("To", msg.Receiver...)
	message.SetHeader("Subject", msg.Title)
	message.SetHeader("Date", message.FormatDate(msg.Date))
	if len(msg.CC) > 0 {
		message.SetHeader("CC", msg.CC...)
	}
	if len(msg.BCC) > 0 {
		message.SetHeader("BCC", msg.BCC...)
	}
	if msg.Text != "" {
		message.SetBody("text/plain;charset=UTF-8", msg.Text)
	} else {
		message.SetBody("text/html;charset=UTF-8", msg.Html)
	}

	if len(msg.Attach) > 0 {
		for _, attach := range msg.Attach {
			message.Attach(attach.Name, gomail.SetCopyFunc(func(writer io.Writer) error {
				_, err := io.Copy(writer, attach.File)
				return err
			}))
		}
	}

	return message
}

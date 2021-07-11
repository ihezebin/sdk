package email

import (
	"fmt"
	"log"
	"testing"
)

//q1a2
func TestEmail(t *testing.T) {
	emailClient := NewClient(Config{
		Username: "86744316@qq.com",
		Password: "hsbpzk***gsjbec**c",
		Host:     HostQQMail,
		Port:     PortQQMail,
	})
	//msg := NewMessage().SetReceiver("378129361@qq.com").SetBody("Hello World！This is a test mail！")
	msg := NewMessage().SetReceiver("378129361@qq.com").
		SetBody(`
			<html>
			<body>
				<h3 style="color:red">
				"Hello World！This is a test mail！"
				</h3>
			</body>
			</html>
		`)
	err := emailClient.Send(msg)
	if err != nil {
		log.Println("send mail err:", err)
	}
	fmt.Println("send mail successfully")
}

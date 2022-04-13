package emailc

import (
	"fmt"
	"github.com/whereabouts/sdk/logger"
	"testing"
	"time"
)

func TestInitEmail(t *testing.T) {
	_, err := NewClientWithConfig(Config{
		Username: "system@whereabouts.icu",
		Password: "*****",
		Host:     HostExmail,
		Port:     PortExmail,
	})
	if err != nil {
		logger.Fatal(err)
	}
}

func TestSendEmail(t *testing.T) {
	client, err := NewClientWithConfig(Config{
		Username: "system@whereabouts.icu",
		Password: "****",
		Host:     HostExmail,
		Port:     PortExmail,
	})
	if err != nil {
		logger.Fatal(err)
	}
	//now, err := timer.Parse("2006-01-02 15:04:05")
	msg := NewMessage().
		WithTitle("test").
		WithReceiver("86744316@qq.com").
		WithDate(time.Now()).
		WithSender("WBTS").
		WithHtml(`
			<html>
			<body>
				<h3 style="color:white;background-color:skyblue">
				"Hello World！This is a test mail！"
				</h3>
			</body>
			</html>
		`)
		//WithAttach(NewAttach("test.txt", strings.NewReader("dsadsad")) )
	err = client.Send(msg)
	if err != nil {
		logger.Fatal("send mail err:", err)
	}
	fmt.Println("send mail successfully")
}

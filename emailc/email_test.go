package emailc

import (
	"fmt"
	"github.com/whereabouts/sdk/logger"
	"github.com/whereabouts/sdk/utils/timer"
	"strings"
	"testing"
)

//q1a2
func TestEmail(t *testing.T) {
	client, err := NewClientWithConfig(Config{
		Username: "heds@whereabouts.icu",
		Password: "dsdsd.",
		Host:     HostExmail,
		Port:     PortExmail,
	})
	if err != nil {
		logger.Fatal(err)
	}
	now, err := timer.Parse("2006-01-02 15:04:05")
	msg := NewMessage().
		WithTitle("test").
		WithReceiver("asdas@qq.com").
		WithDate(now).
		WithText(`
			<html>
			<body>
				<h3 style="color:white;background-color:skyblue">
				"Hello World！This is a test mail！"
				</h3>
			</body>
			</html>
		`).
		WithAttach(NewAttach("test.txt", strings.NewReader("dsadsad")))
	err = client.Send(msg)
	if err != nil {
		logger.Fatal("send mail err:", err)
	}
	fmt.Println("send mail successfully")
}

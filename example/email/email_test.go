package email

import (
	"fmt"
	"github.com/whereabouts/utils/mail"
	"testing"
)

func TestEmail(t *testing.T) {
	auth := mail.Auth("86744316@qq.com", "[your authorization code]", mail.HostQQMail, mail.PortQQMail)
	sender := auth.SetFrom("whereabouts.icu")
	err := sender.SetSubject("Hello World Plain").Plain([]string{"378129361@qq.com"}, "Hello World！This is a test mail！")
	if err != nil {
		fmt.Println("send mail err:", err)
	} else {
		fmt.Println("send mail successfully")
	}
	err = sender.SetSubject("Hello World HTML").Html([]string{"378129361@qq.com"}, `
			<html>
			<body>
				<h3 style="color:red">
				"Hello World！This is a test mail！"
				</h3>
			</body>
			</html>
		`)
	if err != nil {
		fmt.Println("send mail err:", err)
	} else {
		fmt.Println("send mail successfully")
	}
}

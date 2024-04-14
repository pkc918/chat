package utils

import (
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	"github.com/spf13/viper"
	"net/smtp"
	"net/textproto"
)

func SendCodeByQQEmail(toEmail string) error {
	var host = viper.GetString("email.host")
	var port = viper.GetString("email.port")
	var from = viper.GetString("email.from")
	var subject = viper.GetString("email.subject")
	var password = viper.GetString("email.password")
	var serverName = viper.GetString("email.serverName")

	html := fmt.Sprintf("<p>当前验证码：%v</p>", 123456)
	e := &email.Email{
		To:      []string{toEmail},
		From:    fmt.Sprintf("WeChat-Plus <%v>", from),
		Subject: subject,
		HTML:    []byte(html),
		Headers: textproto.MIMEHeader{},
	}
	err := e.SendWithTLS(fmt.Sprintf("%v:%v", host, port), smtp.PlainAuth("", from, password, host), &tls.Config{
		ServerName:         serverName,
		InsecureSkipVerify: true,
	})
	return err
}

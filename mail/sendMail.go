package sendMail

import (
	"crypto/tls"
	"fmt"
	"os"
	"strconv"

	gomail "gopkg.in/gomail.v2"
)

func SendMail(body *string, commentID int) error {
	url := os.Getenv("DEALABS_URL")
	mailHostName := os.Getenv("DEALABS_HOSTNAME")
	mailHostNamePort := os.Getenv("DEALABS_HOSTNAME_PORT")
	mailSender := os.Getenv("DEALABS_MAIL_SENDER")
	mailSenderPassword := os.Getenv("DEALABS_MAIL_SENDER_PASSWORD")

	port, _ := strconv.Atoi(mailHostNamePort)
	d := gomail.NewDialer(mailHostName, port, mailSender, mailSenderPassword)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := gomail.NewMessage()
	m.SetHeader("From", mailSender)
	//TODO put mailingList.txt
	m.SetHeader("To", "exempleReciever@gmail.com")
	m.SetHeader("Subject", fmt.Sprintf("Comment no: %d", commentID))
	m.SetBody("text/html", fmt.Sprintf("%s\n%s", url, *body))

	err := d.DialAndSend(m)
	return err
}

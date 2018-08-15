package sendMail

import (
	"crypto/tls"
	"fmt"
	"os"
	"strconv"
	"strings"
	"io/ioutil"

	gomail "gopkg.in/gomail.v2"
)

func SendMail(body *string, commentID int, email, password string) error {
	url := os.Getenv("DEALABS_URL")
	mailHostName := os.Getenv("DEALABS_HOSTNAME")
	mailHostNamePort := os.Getenv("DEALABS_HOSTNAME_PORT")
	mailingListPath := os.Getenv("DEALABS_MAILINGLIST_PATH") 

	port, _ := strconv.Atoi(mailHostNamePort)
	d := gomail.NewPlainDialer(mailHostName, port, email, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	dat, err := ioutil.ReadFile(mailingListPath)
    if err != nil {
		return err
	}
	file := string(dat)
	file = strings.Trim(file, "\n")

	fmt.Println(">", file, "<")

	mails := strings.Split(file, "\n")

	m := gomail.NewMessage()
	m.SetHeader("From", email)
	m.SetHeader("To", mails...)
	m.SetHeader("Subject", fmt.Sprintf("Comment no: %d", commentID))
	m.SetBody("text/html", fmt.Sprintf("%s\n%s", url, *body))

	return d.DialAndSend(m)
}

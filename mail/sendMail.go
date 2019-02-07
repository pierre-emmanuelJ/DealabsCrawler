package sendMail

import (
	"crypto/tls"
	"fmt"
	"os"
	"strconv"
	"strings"
	"io/ioutil"
	"log"

	gomail "gopkg.in/gomail.v2"
)

func SendMail(body *string, commentID int, email, password string) error {
	url := os.Getenv("DEALABS_URL")
	if url == "" {
		return fmt.Errorf("env DEALABS_URL is missing")
	}
	mailHostName := os.Getenv("DEALABS_HOSTNAME")
	if mailHostName == "" {
		return fmt.Errorf("env DEALABS_HOSTNAME is missing")
	}
	mailHostNamePort := os.Getenv("DEALABS_HOSTNAME_PORT")
	if mailHostNamePort == "" {
		return fmt.Errorf("env DEALABS_HOSTNAME_PORT is missing")
	}
	mailingListFilename := os.Getenv("DEALABS_MAILINGLIST_FILENAME")
	if mailingListFilename == "" {
		return fmt.Errorf("env DEALABS_MAILINGLIST_FILENAME is missing")
	}

	port, _ := strconv.Atoi(mailHostNamePort)
	d := gomail.NewPlainDialer(mailHostName, port, email, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	dat, err := ioutil.ReadFile("./mailinglist/"+mailingListFilename)
	if err != nil {
		return err
	}
	file := string(dat)
	file = strings.Trim(file, "\n")

	mails := strings.Split(file, "\n")

	m := gomail.NewMessage()
	m.SetHeader("From", email)
	m.SetHeader("Bcc", mails...)
	m.SetHeader("Subject", fmt.Sprintf("Comment no: %d", commentID))
	m.SetBody("text/html", fmt.Sprintf("%s\n%s", url, *body))

	return d.DialAndSend(m)
}

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	mail "github.com/pierre-emmanuelJ/DealabsCrawler/mail"

	"github.com/spf13/pflag"

	dealcrawler "github.com/pierre-emmanuelJ/DealabsCrawler/dealabsCrawler"
)

func main() {

	var flagError pflag.ErrorHandling
	docCmd := pflag.NewFlagSet("", flagError)
	var email = docCmd.StringP("sender-mail", "m", "", "sender mail")
	var password = docCmd.StringP("sender-mail-password", "p", "", "Sender password mail")
	var help = docCmd.BoolP("help", "h", false, "Help about any command")

	if err := docCmd.Parse(os.Args); err != nil {
		log.Fatal(err)
	}

	if *help {
		_, err := fmt.Fprintf(os.Stderr, "Usage of %s:\n\n%s", os.Args[0], docCmd.FlagUsages())
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(1)
	}

	if *email == "" || *password == "" {
		log.Fatalf("error: flag %q and %q must be set\n", "sender-mail", "sender-mail-password")
	}

	for {
		comment, err := dealcrawler.Crawler()
		if err != nil {
			log.Fatal(err)
		}
		if comment != nil {
			if err := mail.SendMail(comment.Body, comment.StrID, *email, *password); err != nil {
				log.Fatalf("Failed to send mail : %v", err)
			}
		}

		time.Sleep(30 * time.Second)
	}
}

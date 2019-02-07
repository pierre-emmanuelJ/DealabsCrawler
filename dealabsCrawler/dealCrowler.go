package crawler

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	mail "github.com/pierre-emmanuelJ/DealabsCrawler/mail"
	"golang.org/x/net/html"
)

type Comment struct {
	body  string
	strID string
}

var AllComment []Comment

func renderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return buf.String()
}

func getAllComments(doc *html.Node) error {
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "article" {
			for _, a := range n.Attr {
				if a.Key == "id" {
					body := renderNode(n)
					AllComment = append(AllComment, Comment{body: body, strID: a.Val})
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return nil
}

func Crawler(email, password string) {
	nbComment := 0
	url := os.Getenv("DEALABS_URL")
	if url == "" {
		log.Println("env var DEALABS_URL is missing")
		return
	}

	value := os.Getenv("COMMENTID")
	if value == "" {
		os.Setenv("COMMENTID", "0")
		nbComment = 0
	} else {
		nbComment, _ = strconv.Atoi(value)
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Failed to load page : %v", err)
		return
	}
	bytes, _ := ioutil.ReadAll(resp.Body)

	doc, _ := html.Parse(strings.NewReader(string(bytes)))
	if err := getAllComments(doc); err != nil {
		log.Println("failed to Getallcomment", err)
		return
	}

	commentID := 0
	last := len(AllComment) - 1
	if last < 0 {
		log.Println("No comments found in page")
		return
	}

	if _, err := fmt.Sscanf(AllComment[last].strID, "comment-%d", &commentID); err != nil {
		fmt.Println(err)
		return
	}

	if nbComment < commentID {
		os.Setenv("COMMENTID", strconv.FormatInt(int64(commentID), 10))
		if err := mail.SendMail(&AllComment[last].body, commentID, email, password); err != nil {
			log.Println("Failed to send mail : %v\n", err)
			return
		}
		fmt.Println("Success COMMENTID:", commentID, "sent")
	}
	resp.Body.Close()
	AllComment = nil
}

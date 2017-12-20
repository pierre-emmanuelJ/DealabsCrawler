package crawler

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
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

//TODO remove this global
var AllComment []Comment

//TODO check error
func renderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return buf.String()
}

//TODO check error
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

//TODO return error code to determine if program have to exit or continue !!!!!!!!
func Crawler() {
	nbComment := 0
	url := os.Getenv("DEALABS_URL")
	if url == "" {
		fmt.Errorf("URL link is empty")
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
		fmt.Printf("Failed to load page : %v", err)
		return
	}
	bytes, _ := ioutil.ReadAll(resp.Body)

	doc, _ := html.Parse(strings.NewReader(string(bytes)))
	if err := getAllComments(doc); err != nil {
		fmt.Errorf("failed to Getallcomment", err)
		return
	}

	commentID := 0
	last := len(AllComment) - 1
	if last < 0 {
		fmt.Errorf("No comments found in page")
		return
	}

	if _, err := fmt.Sscanf(AllComment[last].strID, "comment-%d", &commentID); err != nil {
		fmt.Println(err)
		return
	}

	if nbComment < commentID {
		os.Setenv("COMMENTID", strconv.FormatInt(int64(commentID), 10))
		if err := mail.SendMail(&AllComment[last].body, commentID); err != nil {
			fmt.Printf("Failed to send mail : %v", err)
			return
		}
	}
	resp.Body.Close()
	AllComment = nil
}

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

	"golang.org/x/net/html"
)

//Comment represent a Dealabs page comment
type Comment struct {
	Body  string
	StrID string
}

func renderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return buf.String()
}

func getAllComments(doc *html.Node) []Comment {
	var allComment []Comment
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "article" {
			for _, a := range n.Attr {
				if a.Key == "id" {
					body := renderNode(n)
					allComment = append(allComment, Comment{Body: body, StrID: a.Val})
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return allComment
}

//Crawler crawl DEALABS_URL page
func Crawler() (*Comment, error) {
	nbComment := 0
	url := os.Getenv("DEALABS_URL")
	if url == "" {
		return nil, fmt.Errorf("env var DEALABS_URL is missing")
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
		return nil, fmt.Errorf("Failed to load page: %v", err)
	}
	bytes, _ := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()

	doc, _ := html.Parse(strings.NewReader(string(bytes)))
	allComment := getAllComments(doc)
	if len(allComment) == 0 {
		return nil, fmt.Errorf("No comments found in page")
	}

	commentID := 0
	last := len(allComment) - 1
	if _, err := fmt.Sscanf(allComment[last].StrID, "comment-%d", &commentID); err != nil {
		return nil, err
	}

	if nbComment < commentID {
		strID := strconv.FormatInt(int64(commentID), 10)
		os.Setenv("COMMENTID", strID)
		return &Comment{Body: allComment[last].Body, StrID: strID}, nil
	}
	return nil, nil
}

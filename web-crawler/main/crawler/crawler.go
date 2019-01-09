package crawler

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

// GetLinks takes a url string and returns a Page struct with the original url and a string of urls on the page
func GetLinks(url string) Page {
	rsp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making http request for ", url)
		return Page{}
	}
	z := html.NewTokenizer(rsp.Body)
	pageLinks := []string{}
	for {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			return Page{url, pageLinks}
		case tt == html.StartTagToken:
			t := z.Token()
			if t.Data != "a" {
				continue
			}
			ok, href := getHref(t)
			if !ok {
				continue
			}
			pageLinks = append(pageLinks, "https://monzo.com"+href)
		}
	}
}

func getHref(t html.Token) (ok bool, href string) {
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
			if href == "" || href[0] != '/' || href == "/-play-store-redirect" {
				ok = false
			} else if len(href) > 9 && href[0:10] == "/cdn-cgi/l" {
				ok = false
			} else {
				ok = true
			}
		}
	}
	return
}

package crawler

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

// GetURLs takes a url string and returns a map with the original url and a []string of urls on the page
func GetURLs(url string) []string {
	rsp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making http request for ", url)
		return []string{}
	}
	z := html.NewTokenizer(rsp.Body)
	pageLinks := []string{}
	for {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			return pageLinks
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

// Find urls on page and filter out any links to external sites
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

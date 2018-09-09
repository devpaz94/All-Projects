package main

import (
	"golang.org/x/net/html"
	"net/http"
	"fmt"
)


type page struct {
	originalUrl string
	urlList []string
}

func getHref(t html.Token)(ok bool, href string) {
	for _,a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
			if href == "" || href[0] != '/' || href == "/-play-store-redirect"{
				ok = false
			} else if  len(href) > 9 && href[0:10] == "/cdn-cgi/l" {
				ok = false
			} else {
				ok = true
			}
		}
	}
	return
}

func getLinks(url string) page {
	rsp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making http request for ", url)
		return page{}
	}
	z := html.NewTokenizer(rsp.Body)
	pageLinks := []string{}
	for {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			return page{url, pageLinks}
		case tt == html.StartTagToken:
			t := z.Token()
			if t.Data != "a" {
				continue
			}
			ok, href := getHref(t)
			if !ok {
				continue
			}
			pageLinks = append(pageLinks, "https://monzo.com" + href)
		}
	}
}

func worker(jobs <-chan string, results chan<- page)  {
	for n := range jobs {
		results <- getLinks(n)
	}
}

func main() {
	seed := "https://monzo.com/"

	unvisitedUrls := map[string]struct{}{seed: struct{}{}}
	visitedUrls := map[string][]string{}

	for len(unvisitedUrls) > 0 {
		l := len(unvisitedUrls)
		jobs := make(chan string, len(unvisitedUrls))
		results := make(chan page,len(unvisitedUrls))

		for i := 0; i < l; i++ {
			go worker(jobs, results)
		}

		for url := range unvisitedUrls {
			jobs <- url
			visitedUrls[url] = []string{}
			delete(unvisitedUrls, url)
		}

		for i := 0; i < l; i++ {
			links := <-results
			visitedUrls[links.originalUrl] = links.urlList
			for _, url := range links.urlList {
				_, inUnvisited := unvisitedUrls[url]
				_, inVisited := visitedUrls[url]
				if inUnvisited || inVisited {
					continue
				}
				unvisitedUrls[url] = struct{}{}
			}
		}
	}

	for k, v := range visitedUrls {
		fmt.Println(k)
		fmt.Println(v)
		fmt.Println(" ")
	}
	fmt.Println(len(visitedUrls))
}

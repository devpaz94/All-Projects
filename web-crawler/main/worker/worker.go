package worker

import (
	"github.com/All-Projects/web-crawler/main/crawler"
)

// InitialiseWorkers does a thing
func InitialiseWorkers(jobs <-chan string, numberOfThreads int, urlMaps *URLMaps, wg *WaitGroup) {
	for i := 0; i < numberOfThreads; i++ {
		go func() {
			for url := range jobs {
				urls := crawler.GetLinks(url)
				checkLists(urls, urlMaps, wg)
			}
		}()
	}
}

func checkLists(urls crawler.Page, urlMaps *URLMaps, wg *WaitGroup) {
	for _, url := range urls.URLList {
		_, inUnchecked := urlMaps.UnCheckedURLs[url]
		_, inChecked := urlMaps.UnCheckedURLs[url]
		if inUnchecked || inChecked {
			continue
		}
		urlMaps.UnCheckedURLs[url] = struct{}{}

	}
	wg.Worker.Done()
}

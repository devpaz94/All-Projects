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
		uncheckedUrls := ReadMap(&urlMaps.UnCheckedURLs, wg)
		checkedUrls := ReadMap(&urlMaps.CheckedURLs, wg)
		_, inUnchecked := uncheckedUrls[url]
		_, inChecked := checkedUrls[url]
		if inUnchecked || inChecked {
			continue
		}
		WriteMap(&urlMaps.UnCheckedURLs, wg, url, struct{}{})
	}
	wg.Worker.Done()
}

package worker

import (
	"github.com/All-Projects/web-crawler/main/crawler"
)

// InitialiseWorkers does a thing
func InitialiseWorkers(jobs <-chan string, numberOfThreads int, urlMaps *URLMaps) {
	for i := 0; i < numberOfThreads; i++ {
		go func() {
			for url := range jobs {
				urls := crawler.GetLinks(url)
				checkLists(urls, urlMaps)
			}
		}()
	}
}

func checkLists(urls crawler.Page, urlMaps *URLMaps) {
	for _, url := range urls.URLList {
		checkedUrls, uncheckedUrls := ReadMaps(urlMaps)
		_, inUnchecked := uncheckedUrls[url]
		_, inChecked := checkedUrls[url]
		if inUnchecked || inChecked {
			continue
		}
		WriteMap(urlMaps, url, struct{}{})
	}
	Worker.Done()
}

func WorkersNotDone(urlMaps *URLMaps) bool {
	_, unchecked := ReadMaps(urlMaps)
	if len(unchecked) == 0 {
		Worker.Wait()
		_, unchecked := ReadMaps(urlMaps)
		if len(unchecked) == 0 {
			return false
		}
	}
	return true
}

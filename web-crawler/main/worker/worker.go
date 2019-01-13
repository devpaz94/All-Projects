package worker

import (
	"sync"

	"github.com/All-Projects/web-crawler/main/crawler"
)

// InitialiseWorkers does a thing
func InitialiseWorkers(jobs <-chan string, numberOfThreads int, urlMaps *URLMaps, wg *sync.WaitGroup) {
	for i := 0; i < numberOfThreads; i++ {
		go worker(jobs, urlMaps, wg)
	}
}

func worker(jobs <-chan string, urlMaps *URLMaps, wg *sync.WaitGroup) {
	for url := range jobs {
		urls := crawler.GetLinks(url)
		for _, url := range urls.URLList {
			CheckAndWriteMap(urlMaps, url)
		}
		wg.Done()
	}
}

//WorkersNotDone sees whether there are any urls left to check
func WorkersNotDone(urlMaps *URLMaps, wg *sync.WaitGroup) bool {
	_, unchecked := ReadMaps(urlMaps)
	if len(unchecked) == 0 {
		wg.Wait()
		_, unchecked := ReadMaps(urlMaps)
		if len(unchecked) == 0 {
			return false
		}
	}
	return true
}

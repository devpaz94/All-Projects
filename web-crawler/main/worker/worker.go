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
	for seed := range jobs {
		urls := crawler.GetLinks(seed)
		for _, url := range urls {
			CheckAndWriteMaps(urlMaps, url)
		}
		WriteMap(seed, urls, urlMaps)
		wg.Done()
	}
}

//WorkersNotDone sees whether there are any urls left to check
func WorkersNotDone(urlMaps *URLMaps, wg *sync.WaitGroup) bool {
	if len(ReadMap(urlMaps)) == 0 {
		wg.Wait()
		if len(ReadMap(urlMaps)) == 0 {
			return false
		}
	}
	return true
}

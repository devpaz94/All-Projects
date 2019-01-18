package worker

import (
	"sync"

	"github.com/All-Projects/web-crawler/main/crawler"
)

// InitialiseWorkers starts a number of workers (tc) in goroutines
func InitialiseWorkers(jobs <-chan string, tc int, urlMaps *URLMaps, wg *sync.WaitGroup) {
	for i := 0; i < tc; i++ {
		go worker(jobs, urlMaps, wg)
	}
}

func worker(jobs <-chan string, urlMaps *URLMaps, wg *sync.WaitGroup) {
	for seed := range jobs {
		urls := crawler.GetURLs(seed)
		for _, url := range urls {
			CheckAndWriteMaps(urlMaps, url)
		}
		WriteMap(seed, urls, urlMaps)
		wg.Done()
	}
}

//WorkersNotDone waits for all workers to finish and checks whether there are any urls left to crawl
func WorkersNotDone(urlMaps *URLMaps, wg *sync.WaitGroup) bool {
	if len(ReadMap(urlMaps)) == 0 {
		wg.Wait()
		if len(ReadMap(urlMaps)) == 0 {
			return false
		}
	}
	return true
}

package worker

import (
	"sync"

	"github.com/web-crawler/main/crawler"
)

func InitialiseWorkers(jobs <-chan string, numberOfThreads int, wg *sync.WaitGroup, m *sync.Mutex) {
	for i := 0; i < numberOfThreads; i++ {
		go worker(jobs, wg, m)
	}
}

func worker(jobs <-chan string, wg *sync.WaitGroup, m *sync.Mutex) {
	for url := range jobs {
		urls := crawler.GetLinks(url)
		checkLists(urls, wg, m)
	}
}

func checkLists(urls crawler.Page, wg *sync.WaitGroup, m *sync.Mutex) {
	for _, url := range urls.URLList {
		m.Lock()
		_, inUnchecked := UncheckedURLs[url]
		_, inChecked := CheckedURLs[url]
		if inUnchecked || inChecked {
			m.Unlock()
			continue
		}
		UncheckedURLs[url] = struct{}{}
		m.Unlock()
	}
	wg.Done()
}

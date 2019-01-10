package main

import (
	w "github.com/All-Projects/web-crawler/main/worker"
)

func main() {

	seed := "https://monzo.com/"
	numberOfThreads := 100

	wg := w.WaitGroup{}

	urlMaps := w.URLMaps{
		UnCheckedURLs: map[string]struct{}{seed:{}},
		CheckedURLs: map[string][]string{},
	}

	jobs := make(chan string)

	w.InitialiseWorkers(jobs, numberOfThreads, &urlMaps, &wg)

	for w.WorkersNotDone(&wg, &urlMaps) {
		uncheckedUrls := w.UnCheckedRead(&urlMaps.UnCheckedURLs, &wg)
		for url := range uncheckedUrls {
			wg.Worker.Add(1)
			wg.Writer.Add(1)
			urlMaps.CheckedURLs[url] = []string{}
			delete(urlMaps.UnCheckedURLs, url)
			wg.Writer.Done()
			jobs <- url
		}
	}
}

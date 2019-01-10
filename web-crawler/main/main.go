package main

import (
	"fmt"

	w "github.com/All-Projects/web-crawler/main/worker"
)

func main() {

	seed := "https://monzo.com/"
	numberOfThreads := 100

	wg := w.WaitGroup{}

	urlMaps := w.URLMaps{UnCheckedURLs: map[string]struct{}{seed}}

	jobs := make(chan string)

	w.InitialiseWorkers(jobs, numberOfThreads, &urlMaps, &wg)

	for w.WorkersNotDone(wg, urlMaps) {
		for url := range w.UncheckedURLs {
			workerWait.Add(1)
			_, inCurrentlyChecking := currentlyChecking[url]
			if inCurrentlyChecking {
				continue
			}
			currentlyChecking[url] = struct{}{}
			jobs <- url
		}
	}
	fmt.Println(w.CheckedURLs)
	fmt.Println(len(w.CheckedURLs))
}

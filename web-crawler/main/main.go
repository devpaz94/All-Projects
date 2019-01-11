package main

import (
	"fmt"

	w "github.com/All-Projects/web-crawler/main/worker"
)

func main() {

	seed := "https://monzo.com/"
	numberOfThreads := 100

	wg := w.WaitGroup{}

	urlMaps := w.URLMaps{
		UnCheckedURLs: map[string]interface{}{seed: struct{}{}},
		CheckedURLs:   map[string]interface{}{},
	}

	jobs := make(chan string)

	w.InitialiseWorkers(jobs, numberOfThreads, &urlMaps, &wg)

	for w.WorkersNotDone(&wg, &urlMaps) {
		uncheckedUrls := w.ReadMap(&urlMaps.UnCheckedURLs, &wg)
		for url := range uncheckedUrls {
			wg.Worker.Add(1)
			w.DeleteAndWriteMaps(&urlMaps, &wg, url)
			jobs <- url
		}
	}
	fmt.Println(len(urlMaps.CheckedURLs))
}

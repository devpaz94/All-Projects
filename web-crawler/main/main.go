package main

import (
	"fmt"
	"sync"
	"time"

	w "github.com/All-Projects/web-crawler/main/worker"
)

func main() {
	seed := "https://monzo.com/"
	numberOfThreads := 100

	start := time.Now()

	urlMaps := w.URLMaps{
		UnCheckedURLs: map[string]struct{}{seed: struct{}{}},
		CheckedURLs:   map[string][]string{},
	}

	jobs := make(chan string)

	var wg sync.WaitGroup
	w.InitialiseWorkers(jobs, numberOfThreads, &urlMaps, &wg)

	for w.WorkersNotDone(&urlMaps, &wg) {
		for url := range w.ReadMap(&urlMaps) {
			wg.Add(1)
			w.DeleteAndWriteMaps(&urlMaps, url)
			jobs <- url
		}
	}
	fmt.Println(urlMaps.CheckedURLs)
	fmt.Println(len(urlMaps.CheckedURLs))
	fmt.Println(time.Since(start))
}

package main

import (
	"fmt"
	"sync"
	"time"

	w "github.com/All-Projects/web-crawler/main/worker"
)

func main() {

	start := time.Now()
	seed := "https://monzo.com/"
	threadCount := 100

	urlMaps := w.URLMaps{
		CheckedURLs:   map[string][]string{},
		UnCheckedURLs: map[string]struct{}{seed: struct{}{}},
	}

	jobs := make(chan string)
	var wg sync.WaitGroup

	w.InitialiseWorkers(jobs, threadCount, &urlMaps, &wg)
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

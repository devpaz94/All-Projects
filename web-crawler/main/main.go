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
	numberOfThreads := 100

	urlMaps := w.URLMaps{
		UnCheckedURLs: map[string]interface{}{seed: struct{}{}},
		CheckedURLs:   map[string]interface{}{},
	}

	jobs := make(chan string)

	var wg sync.WaitGroup

	w.InitialiseWorkers(jobs, numberOfThreads, &urlMaps, &wg)

	for w.WorkersNotDone(&urlMaps, &wg) {
		_, uncheckedUrls := w.ReadMaps(&urlMaps)
		for url := range uncheckedUrls {
			wg.Add(1)
			w.DeleteAndWriteMaps(&urlMaps, url)
			jobs <- url
		}
	}

	fmt.Println(len(urlMaps.CheckedURLs))
	end := time.Now()
	fmt.Println(end.Sub(start))
}

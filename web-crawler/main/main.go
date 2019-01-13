package main

import (
	"fmt"
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

	w.InitialiseWorkers(jobs, numberOfThreads, &urlMaps)

	for w.WorkersNotDone(&urlMaps) {
		_, uncheckedUrls := w.ReadMaps(&urlMaps)
		for url := range uncheckedUrls {
			w.Worker.Add(1)
			w.DeleteAndWriteMaps(&urlMaps, url)
			jobs <- url
		}
	}

	fmt.Println(len(urlMaps.CheckedURLs))
	end := time.Now()
	fmt.Println(end.Sub(start))
}

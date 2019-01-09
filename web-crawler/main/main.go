package main

import (
	"fmt"
	"sync"

	w "github.com/web-crawler/main/worker"
)

func main() {
	numberOfThreads := 100
	jobs := make(chan string)
	var wg sync.WaitGroup
	var m sync.Mutex

	w.InitialiseWorkers(jobs, numberOfThreads, &wg, &m)

	for w.WorkersNotDone(&wg) {
		for url := range w.UncheckedURLs {
			wg.Add(1)
			delete(w.UncheckedURLs, url)
			w.CheckedURLs[url] = []string{}
			jobs <- url
		}
	}
	fmt.Println(w.CheckedURLs)
	fmt.Println(len(w.CheckedURLs))
}

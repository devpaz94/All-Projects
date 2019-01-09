package worker

import "sync"

var UncheckedURLs = map[string]struct{}{"https://monzo.com/": {}}

var CheckedURLs = map[string][]string{}

func WorkersNotDone(wg *sync.WaitGroup) bool {
	if len(UncheckedURLs) == 0 {
		wg.Wait()
		if len(UncheckedURLs) == 0 {
			return false
		}
	}
	return true
}

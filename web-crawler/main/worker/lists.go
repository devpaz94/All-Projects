package worker

func WorkersNotDone(wg *WaitGroup, urlMaps *URLMaps) bool {
	if len(urlMaps.unCheckedURLs) == 0 {
		wg.worker.Wait()
		if len(urlMaps.unCheckedURLs) == 0 {
			return false
		}
	}
	return true
}

package worker

func WorkersNotDone(wg *WaitGroup, urlMaps *URLMaps) bool {
	if len(urlMaps.UnCheckedURLs) == 0 {
		wg.Worker.Wait()
		if len(urlMaps.UnCheckedURLs) == 0 {
			return false
		}
	}
	return true
}

package worker

func WorkersNotDone(wg *WaitGroup, urlMaps *URLMaps) bool {
	if len(ReadMap(&urlMaps.UnCheckedURLs, wg)) == 0 {
		wg.Worker.Wait()
		if len(ReadMap(&urlMaps.UnCheckedURLs, wg)) == 0 {
			return false
		}
	}
	return true
}

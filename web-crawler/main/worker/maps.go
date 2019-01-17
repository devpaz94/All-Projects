package worker

import "sync"

//URLMaps strcut
type URLMaps struct {
	UnCheckedURLs map[string]struct{}
	CheckedURLs   map[string][]string
	mtx           sync.RWMutex
}

//CheckAndWriteMaps writes to the UnCheckedURLs map
func CheckAndWriteMaps(urlMaps *URLMaps, url string) {
	urlMaps.mtx.Lock()
	_, inUnchecked := urlMaps.UnCheckedURLs[url]
	_, inChecked := urlMaps.CheckedURLs[url]
	if inUnchecked || inChecked {
		urlMaps.mtx.Unlock()

		return
	}
	urlMaps.UnCheckedURLs[url] = struct{}{}
	urlMaps.mtx.Unlock()
}

//DeleteAndWriteMaps moves a url from unchecked to checked
func DeleteAndWriteMaps(urlMaps *URLMaps, k string) {
	urlMaps.mtx.Lock()
	delete(urlMaps.UnCheckedURLs, k)
	urlMaps.CheckedURLs[k] = []string{}
	urlMaps.mtx.Unlock()
}

//ReadMap safely reads from UnCheckedURLs
func ReadMap(urlMaps *URLMaps) map[string]struct{} {
	urlMaps.mtx.RLock()
	unChecked := urlMaps.UnCheckedURLs
	urlMaps.mtx.RUnlock()

	return unChecked
}

//WriteMap safely writes to CheckedURLs
func WriteMap(seed string, urls []string, urlMaps *URLMaps) {
	urlMaps.mtx.Lock()
	urlMaps.CheckedURLs[seed] = urls
	urlMaps.mtx.Unlock()
}

package worker

import "sync"

//URLMaps strcut
type URLMaps struct {
	UnCheckedURLs map[string]interface{}
	CheckedURLs   map[string]interface{}
	mtx           sync.RWMutex
}

//CheckAndWriteMap writes to the UnCheckedURLs map
func CheckAndWriteMap(urlMaps *URLMaps, url string) {
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
	urlMaps.CheckedURLs[k] = struct{}{}
	urlMaps.mtx.Unlock()
}

//ReadMaps safely reads both maps
func ReadMaps(urlMaps *URLMaps) (map[string]interface{}, map[string]interface{}) {
	urlMaps.mtx.RLock()
	checked := urlMaps.CheckedURLs
	unchecked := urlMaps.UnCheckedURLs
	urlMaps.mtx.RUnlock()

	return checked, unchecked
}

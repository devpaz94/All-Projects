package worker

import "sync"

//URLMaps does a thing
type URLMaps struct {
	UnCheckedURLs map[string]interface{}
	CheckedURLs   map[string]interface{}
	mtx           sync.RWMutex
}

func ReadMaps(urlMaps *URLMaps) (map[string]interface{}, map[string]interface{}) {
	urlMaps.mtx.RLock()
	checked := urlMaps.CheckedURLs
	unchecked := urlMaps.UnCheckedURLs
	urlMaps.mtx.RUnlock()
	return checked, unchecked
}

func WriteMap(urlMaps *URLMaps, k string, v struct{}) {
	urlMaps.mtx.Lock()
	urlMaps.UnCheckedURLs[k] = struct{}{}
	urlMaps.mtx.Unlock()
}

func DeleteAndWriteMaps(urlMaps *URLMaps, k string) {
	urlMaps.mtx.Lock()
	delete(urlMaps.UnCheckedURLs, k)
	urlMaps.CheckedURLs[k] = struct{}{}
	urlMaps.mtx.Unlock()

}

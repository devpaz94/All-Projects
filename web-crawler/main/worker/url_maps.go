package worker

//URLMaps does a thing
type URLMaps struct {
	UnCheckedURLs     map[string]struct{}
	CheckedURLs       map[string][]string
}

func UnCheckedRead (m *map[string]struct{}, wg *WaitGroup) map[string]struct{} {
	wg.Reader.Add(1)
	uncheckedUrls := *m
	wg.Reader.Done()
	return uncheckedUrls
}

func CheckedRead (m *map[string][]string, wg *WaitGroup) map[string][]string {
	wg.Reader.Add(1)
	checkedUrls := *m
	wg.Reader.Done()
	return checkedUrls
}

func UnCheckedWrite (m *map[string]struct{}, wg *WaitGroup, k string,  v struct{}) {
	wg.Reader.Add(1)
	m[k] = v
	wg.Reader.Done()
}

func CheckedWrite (m *map[string][]string, wg *WaitGroup, v ) {
	wg.Reader.Add(1)
	checkedUrls := *m
	wg.Reader.Done()
}

package worker

//URLMaps does a thing
type URLMaps struct {
	UnCheckedURLs map[string]interface{}
	CheckedURLs   map[string]interface{}
}

func ReadMap(m *map[string]interface{}, wg *WaitGroup) map[string]interface{} {
	wg.Writer.Wait()
	wg.Reader.Add(1)
	readMap := *m
	wg.Reader.Done()
	return readMap
}

func WriteMap(m *map[string]interface{}, wg *WaitGroup, k string, v struct{}) {
	writeMap := ReadMap(m, wg)
	writeMap[k] = v
	wg.Writer.Wait()
	wg.Writer.Add(1)
	wg.Reader.Wait()
	*m = writeMap
	wg.Writer.Done()
}

func DeleteAndWriteMaps(urlMaps *URLMaps, wg *WaitGroup, k string) {
	m := ReadMap(&urlMaps.CheckedURLs, wg)
	wg.Writer.Wait()
	wg.Writer.Add(1)
	wg.Reader.Wait()
	delete(urlMaps.UnCheckedURLs, k)
	urlMaps.CheckedURLs = m
	wg.Writer.Done()
}

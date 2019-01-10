package worker

//URLMaps does a thing
type URLMaps struct {
	UnCheckedURLs     map[string]struct{}
	checkedURLs       map[string][]string
	currentlyChecking map[string]struct{}
}

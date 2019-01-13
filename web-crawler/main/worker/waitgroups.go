package worker

import "sync"

// type WaitGroup struct {
// 	Reader sync.WaitGroup
// 	Writer sync.WaitGroup
// 	Worker sync.WaitGroup
// }

var Reader sync.WaitGroup
var Writer sync.WaitGroup
var Worker sync.WaitGroup

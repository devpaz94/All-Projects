package worker

import "sync"

type WaitGroup struct {
	Reader   sync.WaitGroup
	Writer  sync.WaitGroup
	Worker sync.WaitGroup
}

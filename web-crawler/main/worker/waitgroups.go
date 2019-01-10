package worker

import "sync"

type WaitGroup struct {
	read   sync.WaitGroup
	write  sync.WaitGroup
	worker sync.WaitGroup
}

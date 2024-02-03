package auth

import (
	"fmt"
	"log"
	"sync"
)

type waiting struct {
	items map[string]int
	cnt   int
	wg    sync.WaitGroup
}

func NewWait() *waiting {
	return &waiting{
		items: make(map[string]int),
		cnt:   0,
		wg:    sync.WaitGroup{},
	}
}

func (w *waiting) AddWait(name string) {
	w.wg.Add(1)
	w.cnt += 1
	w.items[name] = w.cnt
}

func (w *waiting) ClearWait(name string) {
	w.wg.Done()
	w.cnt -= 1
	delete(w.items, name)
	log.Printf("Cleared Wait: %s, waiting: %s", name, w.GetWaits())
}

func (w *waiting) Wait() {
	log.Printf("Waiting for Waitgroup: %s", w.GetWaits())
	w.wg.Wait()
}

func (w *waiting) GetWaits() string {
	return fmt.Sprintf("Waits: %d, Waiting: %v", w.cnt, w.items)
}

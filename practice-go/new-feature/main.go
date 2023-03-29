package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
type printContext struct {
	l  sync.Mutex
	n  atomic.Value
	n2 int
}

func printN(workid int, c *printContext, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		// n := c.n.Load().(int)
		if c.n2 > 20 {
			break
		}
		c.l.Lock()
		c.n2 += 1
		// c.n.Store(n)
		if c.n2 <= 20 {
			fmt.Printf("I'm No.%d , print n: %d \n", workid, c.n2)
		}
		time.Sleep(time.Microsecond * 30)
		c.l.Unlock()
		//runtime.Gosched()
	}
}
func main() {
	c := printContext{l: sync.Mutex{}, n: atomic.Value{}}
	c.n.Store(0)
	wg := sync.WaitGroup{}
	gN := 2
	for i := 0; i < gN; i++ {
		wg.Add(1)
		go printN(i+1, &c, &wg)
	}
	wg.Wait()
}

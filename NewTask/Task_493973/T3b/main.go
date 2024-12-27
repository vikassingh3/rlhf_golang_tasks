package main

import (
	"runtime"
	"sync"
	"testing"
	"time"
)

const numTasks = 10000

func task(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	// Simulate work with some CPU and memory usage
	time.Sleep(time.Duration(id%100) * time.Millisecond)
	runtime.Gosched()
	_ = make([]byte, 1024) // Allocate some memory
}

func BenchmarkSingleWaitGroup(b *testing.B) {
	var wg sync.WaitGroup
	for n := 0; n < b.N; n++ {
		wg.Add(numTasks)
		for i := 0; i < numTasks; i++ {
			go task(&wg, i)
		}
		wg.Wait()
	}
}

func BenchmarkMultipleWaitGroups(b *testing.B) {
	const numGroups = 10
	var wg [numGroups]sync.WaitGroup
	for n := 0; n < b.N; n++ {
		for i := 0; i < numGroups; i++ {
			wg[i].Add(numTasks/numGroups)
			for j := 0; j < numTasks/numGroups; j++ {
				go task(&wg[i], j)
			}
		}
		for i := 0; i < numGroups; i++ {
			wg[i].Wait()
		}
	}
}

func TestMain(m *testing.M) {
	// Set the maximum number of processors to use
	runtime.GOMAXPROCS(runtime.NumCPU())
	m.Run()
}
  
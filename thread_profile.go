package main

import (
	"fmt"
	"net"
	"runtime/pprof"
	"sync"
)

var threadProfile = pprof.Lookup("threadcreate")

func queryThread() {
	fmt.Printf("threads in starting:%d\n", threadProfile.Count())

	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				net.LookupHost("www.google.com")
			}
		}()
	}

	wg.Wait()
	fmt.Printf("threads after LookupHost: %d\n", threadProfile.Count())
}

package main

import (
	"../../rocket"
	"fmt"
	"sync"
)

func main() {

	fmt.Println("Start!")
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		rocket.Add(func() {
			fmt.Println("Hello world!")
			wg.Done()
		})
	}
	wg.Wait()
}

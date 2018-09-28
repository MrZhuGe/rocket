package rocket_test

import (
	"../rocket"
	"fmt"
	"runtime"
	"testing"
	"time"
)

var nTest = 10000

func doSomething1()  {
	time.Sleep(500 * time.Millisecond)
}

func TestRocketPool(t *testing.T) {
	for i := 0; i < nTest; i++ {
		rocket.Add(doSomething1)
	}
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	fmt.Println("mem usage:", mem.TotalAlloc/1024/1024)
}

func TestNoPool(t *testing.T) {
	for i := 0; i < nTest; i++ {
		go doSomething1()
	}
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	fmt.Println("mem usage:", mem.TotalAlloc/1024/1024)
}
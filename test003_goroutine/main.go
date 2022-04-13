package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	//runtime.GOMAXPROCS(1) // 设置逻辑处理器，1个
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("runtime.NumCPU() =", runtime.NumCPU())
	wg.Add(2) // 计数器2
	go func() {
		defer wg.Done()	// 延迟语句 -1
		for i:=0; i<100; i++ {
			fmt.Println("A:%d", i)
		}
	}()

	go func() {
		defer wg.Done()
		for i:=0; i<100; i++ {
			fmt.Println("B:%d", i)
		}
	}()

	wg.Wait()
}

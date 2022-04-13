package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	count int32
	wg sync.WaitGroup
	mutex sync.Mutex
)
func main() {
	//atomic_ex()
	mutex_ex()
}

func mutex_ex() {
	wg.Add(2)
	count = 8
	go intcount_mutex()
	go intcount_mutex()
	wg.Wait()
}

func intcount_mutex() {
	// 互斥锁
	defer wg.Done()
	for i:=0; i<2; i++ {
		mutex.Lock()
		value := count
		runtime.Gosched()
		value++
		count := value
		mutex.Unlock()
		fmt.Println("cout : ", count)
	}
}

func atomic_ex() {
	// 原子操作
	count = 7
	wg.Add(2)
	go intcount_atomic()
	go intcount_atomic()
	wg.Wait()
	fmt.Println(count)
}

func intcount_atomic() {
	defer wg.Done()
	fmt.Printf("----count[%d]----\n", count)
	for i:=0; i<2; i++ {
		fmt.Printf("count[%d]\n", count)
		value := atomic.LoadInt32(&count)
		runtime.Gosched() //让出cpu时间
		value++
		atomic.StoreInt32(&count, value)
		fmt.Println("count:", count, "value:", value)
	}
}

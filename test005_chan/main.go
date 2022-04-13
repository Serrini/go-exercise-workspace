package main

import "fmt"

func main() {
	ch := make(chan int)	//无缓冲通道，需要同步操作

	go func() {
		sum := 0
		for i:=0; i<10; i++ {
			sum += i
		}
		ch <- sum
	}()

	fmt.Println(<-ch)

	//用同步通道实现管道，类似bash中的|
	one := make(chan string)
	two := make(chan string)
	go func() {
		one <- "aaa"
	}()
	go func() {
		var v string
		v =<- one
		two <- v
	}()

	fmt.Println(<-two)

	//有缓冲通道，获取最先返回的服务器信息
	mirrieedQuery()
}

func mirrieedQuery() string {
	responses := make(chan string, 3)
	// 同时发起3个并发goroutine向这三个镜像获取数据
	go func() {
		responses <- requset("aaaaa.io")
	}()

	go func() {
		responses <- requset("bbbbb.io")
	}()

	go func() {
		responses <- requset("ccccc.io")
	}()

	fmt.Println(<-responses)
	return <-responses
}

func requset(hostname string) (response string) {
	return "xxxx from " + hostname
}
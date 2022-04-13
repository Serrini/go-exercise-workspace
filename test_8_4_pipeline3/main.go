package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	numbers := make(chan int)
	squares := make(chan int)

	go number(numbers)
	go square(squares, numbers)
	printer(squares)

	fmt.Println("-----------")
	strchan := make(chan string, 10)
	strchan <- "a"
	strchan <- "b"
	strchan <- "c"
	fmt.Println(cap(strchan))
	fmt.Println(len(strchan))

	fmt.Println(<-strchan)
	fmt.Println(<-strchan)
	fmt.Println(<-strchan)
	//fmt.Println(<-strchan)	//阻塞


	months := [...]string{1:"Jan", 2:"Feb", 3:"Mar", 4:"APR", 5:"May", 6:"Jun", 7:"July", 8:"Aug", 9:"Sep", 10:"Oct",
		11:"Nov", 12:"Dec"}


	for i,v := range months {
		fmt.Printf("%d-%s\n", i, v)
	}

	fmt.Println("---------")
	for _,v := range months {
		fmt.Printf("%s\n", v)
	}

	fmt.Println("------------")

	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	//%x副词参数，它用于指定以十六进制的格式打印数组或slice全部的元素
	//%t副词参数是用于打印布尔型数据，%T副词参数是用于显示一个值对应的数据类型

	fmt.Println("------------------")
	fmt.Println(months)
	fmt.Println("------tonull------")
	tonull(months[:3])					//切片,直接修改内容
	fmt.Println(months)
	fmt.Println("-------zero-------")
	zero(&months)						//指针,直接修改内容
	fmt.Println(months)
	fmt.Println("------------------")

	nums := [...]int{0, 1, 2, 3, 4}
	fmt.Println(sumNum(nums[:]))		//切片


}

func zero(ptr *[13]string) {
	*ptr = [13]string{} //[32]byte{}就可以生成一个32字节的数组。而且每个数组的元素都是零值初始化，也就是0
	//无返回值
}

func tonull(arr []string) {
	for i:=0; i<len(arr); i++ {
		//arr[i] = "null"
		arr[i] = "null"
	}
}

func sumNum(arr []int) int {
	s := 0
	for i:=0; i<len(arr); i++ {
		s += arr[i]
	}
	return s //返回值int
}

func number(outnumber chan<- int) {
	for i:=0; i<100; i++ {
		outnumber <- i
	}

	close(outnumber)
}

func square(outsquare chan<- int, innumber <-chan int) {
	for v := range innumber {
		outsquare <- v * v
	}

	close(outsquare)
}

func printer(outprint <-chan int) {
	for u := range outprint {
		fmt.Println(u)
	}
}

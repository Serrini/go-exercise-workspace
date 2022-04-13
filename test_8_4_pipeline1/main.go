package main

import "fmt"

func main() {
	numbers := make(chan int)
	squares := make(chan int)

	go func() {
		for i:=0; i<100; i++ {
			numbers <- i
		}
		close(numbers)
	}()

	go func() {
		for x := range numbers{
			squares <- x*x
		}

/*
		for {
			x, ok := <-numbers
			if !ok {
				break
			}
			squares <- x*x
		}
*/
		close(squares)
	}()

	for x := range squares{
		fmt.Println(x)
	}
}

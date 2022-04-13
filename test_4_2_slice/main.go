package main

import (
	"fmt"
)
func main() {
	a := [...]int{3, 5, 6, 7, 9}
	reverse0(a[:])
	fmt.Println(a)
}

func reverse0(arr[] int) {
	for i, j := 0, len(arr) - 1; i<j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func reverse1(arr*[] int) {
	for i, j := 0, len(*arr) - 1; i<j; i, j = i+1, j-1 {
		//*arr + i, *arr + j = *arr + j, *arr + i
	}
}
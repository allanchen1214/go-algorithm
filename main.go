package main

import "fmt"

func main() {
	//	fmt.Println("Golang algorithm tutorials")

	arr := []int{2, 5, 7, 9, 10, 222, 2345, 12345, 22222, 23456, 99999}
	target := 99999
	index := BinarySearch(arr, target)
	fmt.Println(index)

	fmt.Println("----")

	index = BinarySearchForLoop(arr, target)
	fmt.Println(index)
}

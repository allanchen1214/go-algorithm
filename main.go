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

	fmt.Println("----")

	arr = []int{100, 201, 333, 754, 65, 21}
	sum := 955
	index1, index2 := TwoSum(arr, sum)
	fmt.Printf("index1: %d, index2: %d\n", index1, index2)

}

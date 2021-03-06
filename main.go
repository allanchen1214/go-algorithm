package main

import (
	"fmt"
	"time"
)

func main() {
	//	fmt.Println("Golang algorithm tutorials")
	var n1, n2, n3, n4, n5, n6 Node
	n6.Val = 6
	n5.Val = 5
	n3.Val = 3
	n3.Left = &n5
	n3.Right = &n6
	n1.Val = 1
	n1.Right = &n3
	n4.Val = 4
	n2.Val = 2
	n2.Right = &n4
	n1.Left = &n2

	PrintBinaryTree1(&n1)
	fmt.Println()

	fmt.Println("----")

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

	fmt.Println("----")

	strArr := []string{"flower", "flow", "flight"}
	prefix := LongestCommonPrefix(strArr)
	fmt.Println(prefix)

	fmt.Println("----")

	str1 := "abcadf"
	str2 := "acbad"
	subStr := TwoStringLongestCommonSubString(str1, str2)
	fmt.Println(subStr)

	str1 = "abcadffeax12xx"
	str2 = "acadffeat"
	subStr = TwoStringLongestCommonSubString(str1, str2)
	fmt.Println(subStr)

	fmt.Println("----")

	arr = []int{100, 201, 333, 754, 65, 21, 100, 5, 300000, 555555, 44, 34567, 2, 99}
	QuickSort(arr)
	fmt.Println()
	arr = []int{100, 201, 333, 754, 65, 21, 100, 5, 300000, 555555, 44, 34567, 2, 99, 121}
	BubbleSort(arr)
	fmt.Println()

	fmt.Println("----")

	//	now := time.Now()
	//	fn := Fibonacci(50)
	//	defer func() {
	//		fmt.Println(time.Now().Sub(now).Milliseconds())
	//	}()
	//	fmt.Println(fn)

	now := time.Now()
	fn := Fibonacci_optimize(100)
	defer func() {
		fmt.Println(time.Now().Sub(now).Milliseconds())
	}()
	fmt.Println(fn)

	fmt.Println("----")

}

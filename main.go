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

	arr = []int{100, 201, 333, 754, 65, 21, 100, 5, 300000, 555555, 44, 34567, 2, 99}
	QuickSort(arr)
	fmt.Println()
	arr = []int{100, 201, 333, 754, 65, 21, 100, 5, 300000, 555555, 44, 34567, 2, 99, 121}
	BubbleSort(arr)
	fmt.Println()

	fmt.Println("----")

	fn := Fibonacci(5)
	fmt.Println(fn)
	fn = Fibonacci_optimize(5)
	fmt.Println(fn)

}

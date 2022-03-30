package main

import (
	"fmt"
)

func BubbleSort(arr []int) {
	l := len(arr)
	var i, j int
	for i = 0; i < l; i++ {
		for j = 0; j < l-i-1; j++ {
			//fmt.Printf("%d\t%d\n", i, j+1)
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	for _, v := range arr {
		fmt.Printf("%v\t", v)
	}
}

//排序思想：
//首先从数列的右边开始往左边找，我们设这个下标为 i，也就是进行减减操作（i--），找到第 1 个比基准数小的值，让它与基准值交换；
//接着从左边开始往右边找，设这个下标为 j，然后执行加加操作（j++），找到第 1 个比基准数大的值，让它与基准值交换；
//然后继续寻找，直到 i 与 j 相遇时结束，最后基准值所在的位置即 k 的位置，也就是说 k 左边的值均比 k 上的值小，而 k 右边的值都比 k 上的值大。
func QuickSort(arr []int) {
	_QuickSort(arr, 0, len(arr)-1)

	for _, v := range arr {
		fmt.Printf("%v\t", v)
	}
}

func _QuickSort(arr []int, left, right int) {
	if right-left < 1 { // 只有1个元素时，左右值相等
		return
	}

	i := left
	j := right
	pivot := left

	for {
		if i >= j {
			break
		}

		for {
			if j <= pivot {
				break
			}
			if arr[j] < arr[pivot] {
				arr[j], arr[pivot] = arr[pivot], arr[j]
				pivot = j
				j--
				break
			} else {
				j--
			}
		}

		for {
			if i >= pivot {
				break
			}
			if arr[i] > arr[pivot] {
				arr[i], arr[pivot] = arr[pivot], arr[i]
				pivot = i
				i++
				break
			} else {
				i++
			}
		}

		for _, v := range arr {
			fmt.Printf("%v\t", v)
		}

		fmt.Println()
	}

	_QuickSort(arr, left, pivot-1)
	_QuickSort(arr, pivot+1, right)
}

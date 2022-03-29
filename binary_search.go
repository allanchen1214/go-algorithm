package main

import (
	"fmt"
)

func BinarySearchForLoop(arr []int, target int) int {
	if len(arr) == 0 {
		return -2
	}
	if len(arr) == 1 {
		if arr[0] != target {
			return -1
		}
		return 0
	}

	lindex := 0
	rindex := len(arr) - 1

	for {
		if lindex > rindex {
			break
		}
		mindex := lindex + (rindex-lindex)/2
		fmt.Printf("lindex: %d, rindex: %d, mindex: %d\n", lindex, rindex, mindex)
		if arr[mindex] == target {
			return mindex
		} else if arr[mindex] > target {
			rindex = mindex - 1
		} else if arr[mindex] < target {
			lindex = mindex + 1
		}
	}
	return -1
}

func BinarySearch(arr []int, target int) int {
	if len(arr) == 0 {
		return -2
	}
	if len(arr) == 1 {
		if arr[0] != target {
			return -1
		}
		return 0
	}
	return _BinarySearch(arr, target, 0, len(arr)-1)
}

func _BinarySearch(arr []int, target, lindex, rindex int) int {
	if lindex > rindex {
		return -1
	}
	mindex := lindex + (rindex-lindex)/2
	fmt.Printf("lindex: %d, rindex: %d, mindex: %d\n", lindex, rindex, mindex)
	if arr[mindex] == target {
		return mindex
	}
	if arr[mindex] > target {
		rindex = mindex - 1
	} else if arr[mindex] < target {
		lindex = mindex + 1
	}
	return _BinarySearch(arr, target, lindex, rindex)
}

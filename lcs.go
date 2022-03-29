package main

import (
	"fmt"
)

// Longest Common SubString

func LongestCommonSubString(arr []string) string {
	return ""
}

func TwoStringLongestCommonSubString(str1, str2 string) string {
	dp := make([][]int, len(str1))
	for i := range dp {
		dp[i] = make([]int, len(str2))
	}

	for i, v1 := range str1 {
		for j, v2 := range str2 {
			if v1 == v2 {
				dp[i][j] = 1
				if i > 0 && j > 0 && dp[i-1][j-1] > 0 {
					dp[i][j] = dp[i-1][j-1] + 1
				}
			}
		}
	}

	// 非逻辑，帮助日志
	for _, v1 := range dp {
		for _, v2 := range v1 {
			fmt.Printf("%v\t", v2)
		}
		fmt.Println()
	}

	row := 0
	col := 0
	max := dp[row][col]
	for i, v1 := range dp {
		for j, v2 := range v1 {
			if v2 > max {
				row = i
				col = j
				max = v2
			}
		}
	}
	//	fmt.Printf("row: %d, col:%d, max: %d\n", row, col, max)
	return str1[row-max+1 : row+1]

}
